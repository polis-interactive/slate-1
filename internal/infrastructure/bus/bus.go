package bus

import (
	"github.com/polis-interactive/slate-1/internal/domain"
)

type bus struct {
	lightingService domain.LightingService
	renderService   domain.RenderService
	graphicsService domain.GraphicsService
	buttonService   domain.ButtonService
	controlService  domain.ControlService
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

func (b *bus) BindControlService(c domain.ControlService) {
	b.controlService = c
}

func (b *bus) Startup() {
	b.graphicsService.Startup()
	b.renderService.Startup()
	b.buttonService.Startup()
	b.controlService.Startup()
}

func (b *bus) Shutdown() {
	b.controlService.Shutdown()
	b.buttonService.Shutdown()
	b.graphicsService.Shutdown()
	b.renderService.Shutdown()
}
