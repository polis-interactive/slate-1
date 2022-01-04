package graphics

import (
	"github.com/polis-interactive/slate-1/internal/domain"
	"github.com/polis-interactive/slate-1/internal/types"
	"log"
	"sync"
)

type service struct {
	graphics *graphics
	mu *sync.Mutex
}

var _ domain.GraphicsService = (*service)(nil)

func NewService(cfg Config, bus Bus) (*service, error) {
	log.Println("Graphics, NewService: creating")

	g, err := newGraphics(cfg, bus)
	if err != nil {
		log.Println("Graphics, NewService: error creating graphics")
		return nil, err
	}
	return &service{
		graphics: g,
		mu: &sync.Mutex{},
	}, nil
}

func (s *service) Startup() {
	log.Println("RenderService Startup: starting")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics != nil {
		s.graphics.startup()
	}
}

func (s *service) Reset() {
	log.Println("RenderService Startup: resetting")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics != nil {
		s.graphics.shutdown()
		s.graphics.startup()
	}
}

func (s *service) Shutdown() {
	log.Println("RenderService Shutdown: shutting down")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics != nil {
		s.graphics.shutdown()
	}
}

func (s *service) GetPb() (pb *types.PixelBuffer, preLockedMutex *sync.RWMutex) {
	s.graphics.mu.RLock()
	return s.graphics.pb, s.graphics.mu
}
