package bus

import (
	"github.com/polis-interactive/slate-italian-plumber-1/internal/domain"
)

type bus struct {
	lightingService domain.LightingService
	renderService   domain.RenderService
	graphicsService domain.GraphicsService
	buttonService   domain.ButtonService
}

func NewBus() *bus {
	return &bus{}
}

func (b *bus) BindRenderService(r domain.RenderService) {
	b.renderService = r
}

func (b *bus) BindGraphicsService(g domain.GraphicsService) {
	b.graphicsService = g
}

func (b *bus) BindLightingService(l domain.LightingService) {
	b.lightingService = l
}

func (b *bus) BindButtonService(btn domain.ButtonService) {
	b.buttonService = btn
}

func (b *bus) Startup() {
	// maybe signal we are starting up
	b.graphicsService.Startup()
	b.renderService.Startup()
	b.buttonService.Startup()
}

func (b *bus) Shutdown() {
	// maybe signal we are shutting down
	b.buttonService.Shutdown()
	b.graphicsService.Shutdown()
	b.renderService.Shutdown()
}
