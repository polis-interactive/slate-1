package render

import (
	"fmt"
	"github.com/polis-interactive/slate-italian-plumber-1/internal/types"
	"log"
)

type terminalRender struct {
	*baseRender
	pb []types.Color
}

var _ render = (*terminalRender)(nil)

func newTerminalRender(base *baseRender) *terminalRender {

	log.Println("terminalRender, newTerminalRender: creating")

	r := &terminalRender{
		baseRender: base,
	}
	base.render = r

	log.Println("terminalRender, newTerminalRender: created")

	return r
}

func (r *terminalRender) runMainLoop() {
	r.pb = make([]types.Color, r.ledCount)
	for {
		err := r.runRenderLoop()
		if err != nil {
			log.Println(fmt.Sprintf("ws2812Render, Main Loop: received error; %s", err.Error()))
		}
		select {
		case _, ok := <- r.shutdowns:
			if !ok {
				goto CloseTerminalLoop
			}
		}
	}

CloseTerminalLoop:
	log.Println("terminalRender runMainLoop, Main Loop: closed")
	r.wg.Done()
}

func (r *terminalRender) runRender() error {

	err := r.bus.CopyLightsToBuffer(r.pb)
	if err != nil {
		return err
	}

	log.Println(r.pb)

	return nil
}
