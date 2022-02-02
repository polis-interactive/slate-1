package button

import (
	"log"
	"sync"
)

type buttonImpl interface {
	runMainLoop()
}

type button struct {
	impl      buttonImpl
	bus       Bus
	mu        *sync.RWMutex
	wg        *sync.WaitGroup
	shutdowns chan struct{}
}

func newButton(conf Config, bus Bus) (*button, error) {
	b := &button{
		bus:       bus,
		mu:        &sync.RWMutex{},
		wg:        &sync.WaitGroup{},
		shutdowns: nil,
	}
	var impl buttonImpl
	var err error
	if conf.GetButtonIsGpio() {
		impl, err = newGpioButton(conf.GetButtonSetup(), b)
	} else {
		impl, err = newKeyboardButton(conf.GetButtonSetup(), b)
	}
	b.impl = impl
	return b, err
}

func (b *button) startup() {

	log.Println("Button, startup; starting")

	if b.shutdowns == nil {
		b.shutdowns = make(chan struct{})
		b.wg.Add(1)
		go b.impl.runMainLoop()
	}

	log.Println("Button, startup; started")
}

func (b *button) shutdown() {

	log.Println("Button, shutdown; shutting down")

	if b.shutdowns != nil {
		close(b.shutdowns)
		b.wg.Wait()
		b.shutdowns = nil
	}
	log.Println("Button, shutdown; finished")
}
