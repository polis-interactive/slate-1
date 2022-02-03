package cloud

import (
	"github.com/polis-interactive/slate-1/internal/infrastructure/grpc"
	"github.com/polis-interactive/slate-1/internal/infrastructure/proxy"
	"log"
	"sync"
)

type Application struct {
	proxy        *proxy.Proxy
	grpcServer   *grpc.Server
	shutdown     bool
	shutdownLock sync.Mutex
}

func NewApplication(conf *Config) (*Application, error) {

	log.Println("Application, NewApplication: creating")

	/* create application instance */
	app := &Application{
		shutdown: true,
	}

	/* create proxy */
	app.proxy = proxy.NewProxy()

	/* create servers */
	grpcServer, err := grpc.NewServer(conf, app.proxy)
	if err != nil {
		log.Fatalf("Application, NewApplication: failed to initialize GRPC server")
		return nil, err
	}
	app.grpcServer = grpcServer

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

	err := app.grpcServer.Startup()
	if err != nil {
		return err
	}

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

	app.proxy.Shutdown()
	app.grpcServer.Shutdown()

	log.Println("Application, Shutdown: finished")

	return nil
}
