package graphics

import (
	"time"
)

type Config interface {
	GetGraphicsShaderName() string
	GetGraphicsReloadOnUpdate() bool
	GetGraphicsDisplayOutput() bool
	GetGraphicsPixelSize() int
	GetGraphicsFrequency() time.Duration
}