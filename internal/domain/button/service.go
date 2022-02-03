package button

import (
	"github.com/polis-interactive/slate-1/internal/domain"
	"log"
	"sync"
)

type service struct {
	button *button
	mu     *sync.Mutex
}

var _ domain.ButtonService = (*service)(nil)

func NewService(cfg Config, bus Bus) (*service, error) {
	log.Println("Button, NewService: creating")

	b, err := newButton(cfg, bus)
	if err != nil {
		log.Println("Graphics, NewService: error creating graphics")
		return nil, err
	}
	return &service{
		button: b,
		mu:     &sync.Mutex{},
	}, nil
}

func (s *service) Startup() {
	log.Println("RenderService Startup: starting")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.button != nil {
		s.button.startup()
	}
}

func (s *service) Shutdown() {
	log.Println("RenderService Shutdown: shutting down")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.button != nil {
		s.button.shutdown()
	}
}
