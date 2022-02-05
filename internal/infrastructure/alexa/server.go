package alexa

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	router       *gin.Engine
	srv          *http.Server
	port         int
	shutdown     bool
	shutdownLock sync.Mutex
}

func NewServer(cfg Config) (*Server, error) {

	log.Println("AlexaServer, NewServer: creating")

	if cfg.GetIsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.POST("/slate-1", buildHandler(cfg.GetApplicationId()))

	log.Println("AlexaServer, NewServer: created")

	return &Server{
		router:   router,
		port:     cfg.GetAlexaPort(),
		shutdown: true,
	}, nil
}

func (s *Server) Startup() error {

	log.Println("AlexaServer, Startup: starting")

	s.shutdownLock.Lock()
	defer s.shutdownLock.Unlock()

	if s.shutdown == false {
		return errors.New("AlexaServer, Startup: Tried to startup server twice")
	}

	addr := fmt.Sprintf("0.0.0.0:%d", s.port)
	log.Println(fmt.Sprintf("AlexaServer, Startup: listening at %s", addr))

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("AlexaServer, Startup: Failed to listen: %v", err)
		return err
	}

	s.srv = &http.Server{
		Handler: s.router,
	}

	go func() {
		err := s.srv.Serve(listener)
		if err != nil {
			log.Printf("AlexaServer: reported err %s", err)
			s.Shutdown()
		}
	}()

	s.shutdown = false

	log.Println("AlexaServer, Startup: started")

	return nil
}

func (s *Server) Shutdown() {

	log.Printf("AlexaServer, Shutdown: Shutting down")

	s.shutdownLock.Lock()
	defer s.shutdownLock.Unlock()
	if s.shutdown {
		return
	}
	s.shutdown = true

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("AlexaServer something went wrong with shutdown...")
	}

	log.Printf("FrontendServer, Shutdown: success")
}