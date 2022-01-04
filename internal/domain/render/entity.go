package render

import (
	"errors"
	"fmt"
	"github.com/polis-interactive/slate-1/internal/domain"
	"log"
	"reflect"
	"sync"
	"time"
)

type render interface {
	startup()
	shutdown()
	runMainLoop()
	runRenderLoop() error
	runRender() error
}


func newRender(cfg renderConfig, bus Bus) (render, error) {

	log.Println("render, newSign: newRender")

	base := &baseRender{
		bus: bus,
		renderFrequency: cfg.GetRenderFrequency(),
		wg: &sync.WaitGroup{},
	}
	switch cfg.GetRenderType() {
	case domain.RenderTypes.WS2812:
		return newWs2812Render(base, cfg), nil
	case domain.RenderTypes.TERMINAL:
		return newTerminalRender(base), nil
	default:
		return nil, errors.New("sign, newRender: Invalid render type")
	}
}

type baseRender struct {
	bus             Bus
	renderFrequency time.Duration
	render          render
	ledCount 		int
	wg *sync.WaitGroup
	shutdowns chan struct{}
}

func (r *baseRender) startup() {

	log.Println(fmt.Sprintf("%s, startup: starting", reflect.TypeOf(r.render)))

	if r.shutdowns == nil {
		r.ledCount = r.bus.GetLightCount()
		r.shutdowns = make(chan struct{})
		r.wg.Add(1)
		go r.render.runMainLoop()
		log.Println(fmt.Sprintf("%s, startup: running", reflect.TypeOf(r.render)))
	}

}

func (r *baseRender) shutdown() {
	log.Println(fmt.Sprintf("%s, shutdown: shutting down", reflect.TypeOf(r.render)))
	if r.shutdowns != nil {
		close(r.shutdowns)
		r.wg.Wait()
		r.shutdowns = nil
	}
	log.Println(fmt.Sprintf("%s, shutdown: done", reflect.TypeOf(r.render)))
}


func (r *baseRender) runRenderLoop() error {
	/* I don't think this is necessary anymore?
	r.bus.SendRenderStateUpdate(true)
	defer func() {
		r.bus.SendRenderStateUpdate(false)
	}()
	 */
	for {
		select {
		case _, ok := <-r.shutdowns:
			if !ok {
				return nil
			}
		case <-time.After(r.renderFrequency):
			err := r.render.runRender()
			if err != nil {
				return err
			}
		}
	}
}


