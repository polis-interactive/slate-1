package lighting

import "github.com/polis-interactive/slate-1/internal/types"

type Config interface {
	GetBoardConfiguration() []types.BoardConfiguration
	GetDisallowedPositions() []types.Point
}
