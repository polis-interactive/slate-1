package application

import (
	"github.com/polis-interactive/slate-1/internal/domain"
	"github.com/polis-interactive/slate-1/internal/domain/button"
	"github.com/polis-interactive/slate-1/internal/domain/control"
	"github.com/polis-interactive/slate-1/internal/domain/graphics"
	"github.com/polis-interactive/slate-1/internal/domain/render"
)

type applicationBus interface {
	Startup()
	Shutdown()
	BindRenderService(renderClient domain.RenderService)
	BindGraphicsService(graphicsClient domain.GraphicsService)
	BindLightingService(stateService domain.LightingService)
	BindButtonService(buttonService domain.ButtonService)
	BindControlService(controlService domain.ControlService)
	render.Bus
	graphics.Bus
	button.Bus
	control.Bus
}
