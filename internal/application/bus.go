package application

import (
	"github.com/polis-interactive/slate-1/internal/domain"
	"github.com/polis-interactive/slate-1/internal/domain/graphics"
	"github.com/polis-interactive/slate-1/internal/domain/render"
)

type applicationBus interface {
	Startup()
	Shutdown()
	BindRenderService(renderClient domain.RenderService)
	BindGraphicsService(graphicsClient domain.GraphicsService)
	BindLightingService(stateService domain.LightingService)
	render.Bus
	graphics.Bus
}
