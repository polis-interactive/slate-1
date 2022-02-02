package application

import (
	"github.com/polis-interactive/slate-italian-plumber-1/internal/domain/button"
	"github.com/polis-interactive/slate-italian-plumber-1/internal/domain/graphics"
	"github.com/polis-interactive/slate-italian-plumber-1/internal/domain/lighting"
	"github.com/polis-interactive/slate-italian-plumber-1/internal/domain/render"
	"github.com/polis-interactive/slate-italian-plumber-1/internal/infrastructure/bus"
	"log"
	"sync"
)

type Application struct {
	serviceBus   applicationBus
	shutdown     bool
	shutdownLock sync.Mutex
}

func NewApplication(conf *Config) (*Application, error) {

	log.Println("Application, NewApplication: creating")

	/* create application instance */
	app := &Application{
		shutdown: true,
	}

	/* create bus */
	app.serviceBus = bus.NewBus()

	/* create services */

	lightingService := lighting.NewService(conf)
	app.serviceBus.BindLightingService(lightingService)

	renderService, err := render.NewService(conf, app.serviceBus)
	if err != nil {
		log.Fatalf("Application, NewApplication: failed to initialize render service")
		return nil, err
	}
	app.serviceBus.BindRenderService(renderService)

	graphicsService, err := graphics.NewService(conf, app.serviceBus)
	if err != nil {
		log.Fatalf("Application, NewApplication: failed to initialize graphics service")
		return nil, err
	}
	app.serviceBus.BindGraphicsService(graphicsService)

	buttonService, err := button.NewService(conf, app.serviceBus)
	if err != nil {
		log.Fatalf("Application, NewApplication: failed to initialize button service")
		return nil, err
	}
	app.serviceBus.BindButtonService(buttonService)

	return app, nil
}

func (app *Application) Startup() error {

	log.Println("Application, Startup: starting")

	app.shutdownLock.Lock()
	defer app.shutdownLock.Unlock()
	if app.shutdown == false {
		return nil
	}

	app.shutdown = false

	app.serviceBus.Startup()

	log.Println("Application, Startup: started")

	return nil
}

func (app *Application) Shutdown() error {

	log.Println("Application, Shutdown: shutting down")

	app.shutdownLock.Lock()
	defer app.shutdownLock.Unlock()
	if app.shutdown {
		return nil
	}
	app.shutdown = true

	app.serviceBus.Shutdown()

	log.Println("Application, Shutdown: finished")

	return nil
}
