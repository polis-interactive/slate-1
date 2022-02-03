package control

import (
	"github.com/polis-interactive/slate-1/internal/domain"
	"log"
	"sync"
)

type service struct {
	connection *connection
	mu         *sync.Mutex
}

var _ domain.ControlService = (*service)(nil)

func NewService(cfg Config, bus Bus) *service {
	log.Println("ControlService, NewService: creating")

	c := newConnection(cfg, bus)
	return &service{
		connection: c,
		mu:         &sync.Mutex{},
	}
}

func (s *service) Startup() {
	log.Println("ControlService Startup: starting")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.connection != nil {
		s.connection.startup()
	}
}

func (s *service) Shutdown() {
	log.Println("RenderService Shutdown: shutting down")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.connection != nil {
		s.connection.shutdown()
		s.connection = nil
	}
}
