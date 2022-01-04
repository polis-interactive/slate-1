package bus

import "github.com/polis-interactive/slate-italian-plumber-1/internal/types"

func (b *bus) GetLightGrid() types.Grid {
	return b.lightingService.GetGrid()
}


