package graphics

import "github.com/polis-interactive/slate-1/internal/types"

type Bus interface {
	GetLightGrid() types.Grid
}