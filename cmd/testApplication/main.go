package main

import (
	"github.com/polis-interactive/slate-1/internal/application"
	"github.com/polis-interactive/slate-1/internal/domain"
	"github.com/polis-interactive/slate-1/internal/types"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var testConfiguration = []types.BoardConfiguration{
	types.NewBoardConfiguration(types.Board1x7, types.Orient90, types.CreatePoint(0, 0)),
}

var testDisallowed = []types.Point{
	types.CreatePoint(2, 0),
	types.CreatePoint(3, 0),
}

func main() {
	conf := &application.Config{
		LightingConfig: &application.LightingConfig{
			BoardConfiguration:  testConfiguration,
			DisallowedPositions: testDisallowed,
		},
		RenderConfig: &application.RenderConfig{
			RenderType:      domain.RenderTypes.TERMINAL,
			RenderFrequency: 10 * time.Second,
		},
		GraphicsConfig: &application.GraphicsConfig{
			ShaderName: "slate-1",
			ReloadOnUpdate: true,
			DisplayOutput: true,
			PixelSize: 30,
			Frequency: 20 * time.Millisecond,
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
