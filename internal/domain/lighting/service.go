package lighting

import (
	"fmt"
	"github.com/polis-interactive/slate-1/internal/types"
	"log"
)

type Service struct {
	lights []types.Light
	lastLight int
	grid types.Grid
}

func NewService(cfg Config) *Service {

	log.Println("Lighting, NewService: creating")

	lights, lastLight := generateLights(cfg.GetBoardConfiguration(), cfg.GetDisallowedPositions())

	grid := types.Grid{}
	for _, l := range lights {
		grid.TrySetMinMax(l.Position)
	}

	log.Println(fmt.Sprintf(
		"MinP (%d, %d); MaxP (%d, %d); last led %d",
		grid.MinX, grid.MinY, grid.MaxX, grid.MaxY,lastLight,
	))


	log.Println("Lighting, NewService: created")

	return &Service{
		lights:    lights,
		lastLight: lastLight,
		grid: grid,
	}
}

func (s *Service) GetLightCount() int {
	return s.lastLight
}

func (s *Service) GetGrid() types.Grid {
	return s.grid
}

func (s *Service) GetLights() []types.Light {
	return s.lights
}