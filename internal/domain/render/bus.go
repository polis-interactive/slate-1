package render

import "github.com/polis-interactive/slate-italian-plumber-1/internal/types"

type Bus interface {
	GetLightCount() int
	CopyLightsToBuffer(buff []types.Color) error
}