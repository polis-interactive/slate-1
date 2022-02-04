package main

import (
	"github.com/polis-interactive/slate-1/data"
	"github.com/polis-interactive/slate-1/internal/application"
	"github.com/polis-interactive/slate-1/internal/cloud"
	"github.com/polis-interactive/slate-1/internal/domain"
	"github.com/polis-interactive/slate-1/internal/domain/button"
	"github.com/polis-interactive/slate-1/internal/types"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	serviceIp, err := cloud.DnsLookupIp("grpc.polis.tv")
	if err != nil {
		log.Fatalln(err)
		return
	}

	conf := &application.Config{
		LightingConfig: &application.LightingConfig{
			BoardConfiguration:  data.TerminalBoardConfiguration,
			DisallowedPositions: data.TerminalDisallowedPositions,
		},
		RenderConfig: &application.RenderConfig{
			RenderType:      domain.RenderTypes.TERMINAL,
			RenderFrequency: 1 * time.Second,
		},
		GraphicsConfig: &application.GraphicsConfig{
			ShaderName:     "slate-1",
			ReloadOnUpdate: false,
			DisplayOutput:  true,
			PixelSize:      30,
			Frequency:      33 * time.Millisecond,
		},
		ButtonConfig: &application.ButtonConfig{
			ButtonIsGpio: false,
			ButtonSetup: &button.Setup{
				KeyOrGpioIn: 1,
			},
		},
		ControlConfig: &application.ControlConfig{
			ServerAddress: serviceIp,
			ServerPort:    6969,
			TLSConfig: &types.TLSConfig{
				CertFile:      "./certs/client.pem",
				KeyFile:       "./certs/client-key.pem",
				CAFile:        "./certs/ca.pem",
				ServerAddress: "0.0.0.0",
				Server:        false,
			},
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
