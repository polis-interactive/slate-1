package main

import (
	"github.com/polis-interactive/slate-1/internal/cloud"
	"github.com/polis-interactive/slate-1/internal/types"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := &cloud.Config{
		IpInterface: "0.0.0.0",
		GrpcPort:    6969,
		TLSConfig: &types.TLSConfig{
			CertFile:      "./certs/server.pem",
			KeyFile:       "./certs/server-key.pem",
			CAFile:        "./certs/ca.pem",
			ServerAddress: "127.0.0.1",
			Server:        true,
		},
	}

	app, err := cloud.NewApplication(conf)
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
