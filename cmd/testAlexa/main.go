package main

import (
	"fmt"
	"github.com/polis-interactive/slate-1/internal/infrastructure/alexa"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type AlexaConfig struct {
}

func (a AlexaConfig) GetIpInterface() string {
	return "0.0.0.0"
}

func (a AlexaConfig) GetApplicationId() string {
	return "amzn1.ask.skill.69a5128a-d6b6-4bd2-888d-f388e8986c7b"
}

func (a AlexaConfig) GetAlexaPort() int {
	return 420
}

func (a AlexaConfig) GetIsProduction() bool {
	return false
}

type FakeProxy struct{}

func (f FakeProxy) HandleAlexaCommand(isOff bool) error {
	log.Println(fmt.Sprintf("new stat is %t", isOff))
	return nil
}

func main() {
	cfg := &AlexaConfig{}
	p := &FakeProxy{}
	srv, err := alexa.NewServer(cfg, p)
	if err != nil {
		panic(err)
	}

	err = srv.Startup()
	if err != nil {
		log.Println("Main: failed to startup, shutting down")
		srv.Shutdown()
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	srv.Shutdown()
}
