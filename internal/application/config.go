package application

import (
	"github.com/polis-interactive/slate-1/internal/domain"
	"github.com/polis-interactive/slate-1/internal/domain/button"
	"github.com/polis-interactive/slate-1/internal/types"
	"time"
)

type LightingConfig struct {
	BoardConfiguration  []types.BoardConfiguration
	DisallowedPositions []types.Point
}

func (l *LightingConfig) GetBoardConfiguration() []types.BoardConfiguration {
	return l.BoardConfiguration
}

func (l *LightingConfig) GetDisallowedPositions() []types.Point {
	return l.DisallowedPositions
}

type RenderConfig struct {
	RenderType      domain.RenderType
	RenderFrequency time.Duration
}

func (r *RenderConfig) GetRenderType() domain.RenderType {
	return r.RenderType
}

func (r *RenderConfig) GetRenderFrequency() time.Duration {
	return r.RenderFrequency
}

type Ws2812Config struct {
	GpioPin   types.GpioPinType
	StripType types.StripType
	Gamma     float32
}

func (w *Ws2812Config) GetGpioPin() types.GpioPinType {
	return w.GpioPin
}

func (w *Ws2812Config) GetStripType() types.StripType {
	return w.StripType
}

func (w *Ws2812Config) GetGamma() float32 {
	return w.Gamma
}

type GraphicsConfig struct {
	ShaderName     string
	DisplayOutput  bool
	ReloadOnUpdate bool
	PixelSize      int
	Frequency      time.Duration
}

func (g *GraphicsConfig) GetGraphicsShaderName() string {
	return g.ShaderName
}

func (g *GraphicsConfig) GetGraphicsReloadOnUpdate() bool {
	return g.ReloadOnUpdate
}

func (g *GraphicsConfig) GetGraphicsDisplayOutput() bool {
	return g.DisplayOutput
}

func (g *GraphicsConfig) GetGraphicsPixelSize() int {
	return g.PixelSize
}

func (g *GraphicsConfig) GetGraphicsFrequency() time.Duration {
	return g.Frequency
}

type ButtonConfig struct {
	ButtonIsGpio  bool
	ButtonSetup   *button.Setup
	ReadFrequency time.Duration
}

func (b *ButtonConfig) GetButtonIsGpio() bool {
	return b.ButtonIsGpio
}
func (b *ButtonConfig) GetButtonSetup() *button.Setup {
	return b.ButtonSetup
}

func (b *ButtonConfig) GetReadFrequency() time.Duration {
	return b.ReadFrequency
}

type ControlConfig struct {
	ConnectionInterface string
}

func (c *ControlConfig) GetGrpcServerAddress() string {
	return c.ConnectionInterface
}

type Config struct {
	*LightingConfig
	*RenderConfig
	*Ws2812Config
	*GraphicsConfig
	*ButtonConfig
	*ControlConfig
}
