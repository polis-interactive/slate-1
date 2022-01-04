package render

import (
	"fmt"
	"log"
)

type terminalRender struct {
	*baseRender
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

	log.Println(fmt.Sprintf("Going to render to %d leds", r.ledCount))

	return nil
}
