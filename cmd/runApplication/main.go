package main

import (
	"github.com/polis-interactive/slate-italian-plumber-1/data"
	"github.com/polis-interactive/slate-italian-plumber-1/internal/application"
	"github.com/polis-interactive/slate-italian-plumber-1/internal/domain"
	"github.com/polis-interactive/slate-italian-plumber-1/internal/types"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)



func main() {
	conf := &application.Config{
		LightingConfig: &application.LightingConfig{
			BoardConfiguration:  data.DefaultBoardConfiguration,
			DisallowedPositions: data.DefaultDisallowedPositions,
		},
		RenderConfig: &application.RenderConfig{
			RenderType:      domain.RenderTypes.WS2812,
			RenderFrequency: 33 * time.Millisecond,
		},
		Ws2812Config: &application.Ws2812Config{
			GpioPin:   types.GpioPinTypes.GPIO19,
			StripType: types.StripTypes.WS2811GRB,
			Gamma:     1.2,
		},
		GraphicsConfig: &application.GraphicsConfig{
			ShaderName: "slate-1",
			ReloadOnUpdate: false,
			DisplayOutput: false,
			PixelSize: 1,
			Frequency: 33 * time.Millisecond,
		},
	}

	app, err := application.NewApplication(conf)
	if err != nil {
		panic(err)
	}

	err = app.Startup()
	if err != nil {
		log.Println("Main: failed to startup, shutting down")
		err2 := app.Shutdown()
		if err2 != nil {
			log.Println("Main: issue shutting down; ", err2)
		}
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	err = app.Shutdown()
	if err != nil {
		log.Println("Main: issue shutting down; ", err)
	}
}

