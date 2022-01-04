package render

import "github.com/polis-interactive/slate-1/internal/types"

type Bus interface {
	GetLightCount() int
	CopyLightsToBuffer(buff []types.Color) error
}