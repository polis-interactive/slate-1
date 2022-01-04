package lighting

import "github.com/polis-interactive/slate-italian-plumber-1/internal/types"

func generateLights(boards []types.BoardConfiguration, disallowed []types.Point) ([]types.Light, int) {
	// outputs
	lights := make([]types.Light, 0)
	lastLight := 0

	// temporary variables
	ledCount := 49
	maxNominalY := 6
	for _, config := range boards {
		if config.Type == types.Board7x7 {
			ledCount = 49
			maxNominalY = 6
		} else {
			ledCount = 7
			maxNominalY = 0
		}
		for i := 0; i < ledCount; i++ {
			yPos := i / 7
			isOddYPos := (yPos % 2) != 0
			xPos := i % 7
			if isOddYPos {
				xPos = 6 - xPos
			}

			position := config.StartingPoint
			isAllowed := true

			switch config.Orientation {
			case types.Orient0:
				position.AlterPoint(position.X + yPos, position.Y + (6 - xPos))
			case types.Orient90:
				position.AlterPoint(position.X + xPos, position.Y + yPos)
			case types.Orient180:
				position.AlterPoint(position.X + (maxNominalY - yPos), position.Y + xPos)
			case types.Orient270:
				position.AlterPoint(position.X + (6 - xPos), position.Y + (maxNominalY - yPos))
			}

			for _, p := range disallowed {
				if position.IsEqual(p) {
					isAllowed = false
					break
				}
			}

			l := types.Light{
				Position: position,
				Pixel: 	  lastLight,
				Show:     isAllowed,
				Color:    types.Color{},
			}
			lights = append(lights, l)
			lastLight += 1
		}
	}
	return lights, lastLight
}

