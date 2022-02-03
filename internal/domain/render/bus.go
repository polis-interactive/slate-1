package render

import "github.com/polis-interactive/slate-1/internal/types"

type Bus interface {
	GetLightCount() int
	CopyLightsToColorBuffer(buff []types.Color) error
	CopyLightsToUint32Buffer(buff []uint32) (isOff bool, err error)
}
