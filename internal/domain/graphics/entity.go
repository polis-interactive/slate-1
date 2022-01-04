package graphics

import (
	"fmt"
	"github.com/polis-interactive/slate-1/internal/domain"
	"github.com/polis-interactive/slate-1/internal/types/shader"
	"log"
	"sync"
	"time"
)

type graphics struct {
	fileHandle string
	displayOutput bool
	reloadOnUpdate bool
	pixelSize int
	graphicsFrequency time.Duration
	gs *shader.GraphicsShader
	bus Bus
	mu *sync.RWMutex
	wg *sync.WaitGroup
	shutdowns chan struct{}
}

func newGraphics(cfg Config, bus Bus) (*graphics, error) {
	log.Println("graphics, newGraphics: creating")
	shaderName := cfg.GetGraphicsShaderName()
	fileHandle, err := shader.GetShaderQualifiedPath(shaderName, domain.Program)
	if err != nil {
		log.Fatalln(fmt.Sprintf("graphics, newGraphics: couldn't find shader %s", shaderName))
		return nil, err
	}
	return &graphics{
		fileHandle: fileHandle,
		reloadOnUpdate: cfg.GetGraphicsReloadOnUpdate(),
		displayOutput: cfg.GetGraphicsDisplayOutput(),
		graphicsFrequency: cfg.GetGraphicsFrequency(),
		pixelSize: cfg.GetGraphicsPixelSize(),
		gs: nil,
		bus: bus,
		mu: &sync.RWMutex{},
		wg: &sync.WaitGroup{},
	}, nil
}

func (g *graphics) startup() {

	log.Println("Graphics, startup; starting")

	if g.shutdowns == nil {
		g.shutdowns = make(chan struct{})
		g.wg.Add(1)
		go g.runMainLoop()
	}

	log.Println("Graphics, startup; started")
}

func (g *graphics) shutdown() {

	log.Println("Graphics, shutdown; shutting down")

	if g.shutdowns != nil {
		close(g.shutdowns)
		g.wg.Wait()
		g.shutdowns = nil
	}
	log.Println("Graphics, shutdown; finished")
}

func (g *graphics) runMainLoop() {
	for {
		err := g.runGraphicsLoop()
		if err != nil {
			log.Println(fmt.Sprintf("Graphics, Main Loop: received error; %s", err.Error()))
		}
		select {
		case _, ok := <- g.shutdowns:
			if !ok {
				goto CloseMainLoop
			}
		case <-time.After(5 * time.Second):
			log.Println("Graphics, Main Loop: retrying window")
		}
	}

CloseMainLoop:
	log.Println("Graphics runMainLoop, Main Loop: closed")
	g.wg.Done()
}

func (g *graphics) runGraphicsLoop() error {

	grid := g.bus.GetLightGrid()
	gridWidth := grid.MaxX - grid.MinX + 1
	gridHeight := grid.MaxY - grid.MinY + 1

	if g.displayOutput {
		gridWidth = gridWidth * g.pixelSize
		gridHeight = gridHeight * g.pixelSize
	}

	gs, err := shader.NewGraphicsShader(g.fileHandle, gridWidth, gridHeight)
	if err != nil {
		return err
	}

	g.gs = gs
	defer func() {
		g.gs.Cleanup()
		g.gs = nil
		g.mu.Lock()
		// black out pixels
		g.mu.Unlock()
	}()

	for {
		select {
		case _, ok := <-g.shutdowns:
			if !ok {
				return nil
			}
		case <-time.After(g.graphicsFrequency):
			if g.reloadOnUpdate {
				err = g.gs.ReloadShader()
				if err != nil {
					return err
				}
			}

			err = g.gs.RunShader()
			if err != nil {
				return err
			}
			g.mu.Lock()
			// copy to pixels
			g.mu.Unlock()
			if g.displayOutput {
				err = g.gs.DisplayShader()
				if err != nil {
					return err
				}
			}
		}
	}
}
