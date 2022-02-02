package button

import (
	"errors"
	"fmt"
	"log"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"time"
)

type gpioButton struct {
	button               *button
	gpio                 gpio.PinIO
	pin                  int
	lastButtonWasPressed bool
	readFrequency        time.Duration
}

func newGpioButton(conf Config, b *button) (*gpioButton, error) {
	log.Println("newGpioButton: creating")
	g := &gpioButton{
		button:               b,
		pin:                  conf.GetButtonSetup().KeyOrGpioIn,
		lastButtonWasPressed: false,
		readFrequency:        conf.GetReadFrequency(),
	}
	pinString := fmt.Sprintf("GPIO%d", g.pin)
	g.gpio = gpioreg.ByName(pinString)
	if g.gpio == nil {
		err := fmt.Sprintf("newGpioButton: failed to find %s", pinString)
		log.Println(err)
		return nil, errors.New(err)
	}
	err := g.gpio.In(gpio.PullUp, gpio.NoEdge)
	if err != nil {
		log.Print(fmt.Sprintf("Toggle, initialize: failed to initialize %s", pinString))
		return nil, err
	}
	log.Println("newGpioButton: created")
	return g, nil
}

func (g *gpioButton) runMainLoop() {
	log.Println("gpioButton, RunMainLoop: running")
	for {
		timer := time.NewTimer(g.readFrequency)
		select {
		case _, ok := <-g.button.shutdowns:
			if !ok {
				goto finish
			}
		case <-timer.C:
			isPushed := g.gpio.Read() != gpio.High
			if g.lastButtonWasPressed != isPushed {
				if isPushed {
					g.button.bus.HandleButtonPush()
				}
				g.lastButtonWasPressed = isPushed
			}
		}
	}
finish:
	log.Println("gpioButton, RunMainLoop: closing")
	g.button.wg.Done()
	log.Println("gpioButton, RunMainLoop: closed")
}
