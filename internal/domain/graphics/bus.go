package graphics

import "github.com/polis-interactive/slate-italian-plumber-1/internal/types"

type Bus interface {
	GetLightGrid() types.Grid
}