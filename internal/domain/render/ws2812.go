package render

import (
	"fmt"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	"log"
	"time"
)

type ws2812Render struct {
	*baseRender
	options *ws2811.Option
	strip *ws2811.WS2811
	gamma float32
}

var _ render = (*ws2812Render)(nil)

func newWs2812Render(base *baseRender, cfg ws2812RenderConfig) *ws2812Render {

	log.Println("ws2812Render, newWs2812Render: creating")

	options := ws2811.DefaultOptions
	options.Channels[0].GpioPin = int(cfg.GetGpioPin())
	options.Channels[0].StripeType = int(cfg.GetStripType())
	r := &ws2812Render{
		baseRender: base,
		options:    &options,
		strip: nil,
		gamma: cfg.GetGamma(),
	}

	base.render = r

	log.Println("ws2812Render, newWs2812Render: created")
	return r
}

func (r *ws2812Render) runMainLoop() {

	log.Println("ws2812Render, Main Loop: running")

	for {
		err := func (r *ws2812Render) error {
			r.options.Channels[0].LedCount = r.ledCount
			dev, err := ws2811.MakeWS2811(r.options)
			if err != nil {
				return err
			}
			err = dev.Init()
			if err != nil {
				return err
			}
			defer dev.Fini()
			r.strip = dev
			err = r.runRenderLoop()
			if err != nil {
				return err
			}
			r.tryBlackoutStrip()
			return nil
		}(r)
		r.strip = nil
		if err != nil {
			log.Println(fmt.Sprintf("ws2812Render, Main Loop: received error; %s", err.Error()))
		}
		select {
		case _, ok := <- r.shutdowns:
			if !ok {
				goto CloseWs2812Loop
			}
		case <-time.After(5 * time.Second):
			log.Println("ws2812Render, Main Loop: retrying connection")
		}
	}

CloseWs2812Loop:
	log.Println("ws2812Render, Main Loop: closed")
	r.wg.Done()
}

func (r *ws2812Render) runRender() error {

	err := r.bus.CopyLightsToUint32Buffer(r.strip.Leds(0))
	if err != nil {
		return err
	}

	err = r.strip.Render()
	return err
}

func (r *ws2812Render) tryBlackoutStrip() {
	if r.strip == nil {
		log.Println("ws2812Render, tryBlackoutStrip: couldn't do it, strip is null")
		return
	}
	leds := r.strip.Leds(0)
	for i, _ := range leds {
		leds[i] = 0
	}
	err := r.strip.Render()
	if err != nil {
		log.Println(fmt.Sprintf("ws2812Render, tryBlackoutStrip: failed for some reason; %s", err.Error()))
	}
}