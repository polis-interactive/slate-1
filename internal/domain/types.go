package domain

import (
	"github.com/polis-interactive/slate-italian-plumber-1/internal/types"
	"sync"
)

const Program = "slate-italian-plumber-1"

type RenderState bool

type RenderType string

const (
	ws2812Render   RenderType = "WS2812_RENDER"
	terminalRender            = "TERMINAL_RENDER"
)

var RenderTypes = struct {
	WS2812   RenderType
	TERMINAL RenderType
}{
	WS2812:   ws2812Render,
	TERMINAL: terminalRender,
}

type RenderService interface {
	Startup()
	Reset()
	Shutdown()
}

type GraphicsService interface {
	Startup()
	Reset()
	Shutdown()
	GetPb() (pb *types.PixelBuffer, preLockedMutex *sync.RWMutex)
	HandleButtonPush() (isOff bool)
}

type ButtonService interface {
	Startup()
	Shutdown()
}

type LightingService interface {
	GetLightCount() int
	GetGrid() types.Grid
	GetLights() (lights []types.Light, preLockedMutex *sync.RWMutex)
}
