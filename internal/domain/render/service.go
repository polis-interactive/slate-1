package render

import (
	"log"
	"sync"
)

type service struct {
	render render
	mu *sync.Mutex
}

func NewService(cfg Config, bus Bus) (*service, error) {
	log.Println("Render, NewService: created")

	r, err := newRender(cfg, bus)
	if err != nil {
		log.Println("Render, NewService: error creating render")
		return nil, err
	}
	return &service{
		render: r,
		mu: &sync.Mutex{},
	}, nil
}

func (s *service) Startup() {
	log.Println("RenderService Startup: starting")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.render != nil {
		s.render.startup()
	}
}

func (s *service) Reset() {
	log.Println("RenderService Startup: resetting")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.render != nil {
		s.render.shutdown()
		s.render.startup()
	}
}

func (s *service) Shutdown() {
	log.Println("RenderService Shutdown: shutting down")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.render != nil {
		s.render.shutdown()
	}
}
