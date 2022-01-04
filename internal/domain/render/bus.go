package render

import "github.com/polis-interactive/slate-italian-plumber-1/internal/types"

type Bus interface {
	GetLightCount() int
	CopyLightsToColorBuffer(buff []types.Color) error
	CopyLightsToUint32Buffer(buff []uint32) error
}