package lighting

import "github.com/polis-interactive/slate-italian-plumber-1/internal/types"

type Config interface {
	GetBoardConfiguration() []types.BoardConfiguration
	GetDisallowedPositions() []types.Point
}
