package data

import "github.com/polis-interactive/slate-1/internal/types"

var TerminalBoardConfiguration = []types.BoardConfiguration{
	types.NewBoardConfiguration(types.Board1x7, types.Orient90, types.CreatePoint(0, 0)),
}

var TerminalDisallowedPositions = []types.Point{
	types.CreatePoint(2, 0),
	types.CreatePoint(3, 0),
}

var PiBoardConfiguration = []types.BoardConfiguration{
	types.NewBoardConfiguration(types.Board1x7, types.Orient90, types.CreatePoint(0, 0)),
	types.NewBoardConfiguration(types.Board1x7, types.Orient90, types.CreatePoint(7, 0)),
	types.NewBoardConfiguration(types.Board1x7, types.Orient270, types.CreatePoint(7, 1)),
}

var PiDisallowedPositions = []types.Point{
	types.CreatePoint(2, 0),
	types.CreatePoint(3, 0),
	types.CreatePoint(9, 0),
	types.CreatePoint(10, 0),
	types.CreatePoint(9, 1),
	types.CreatePoint(10, 1),
}

var DefaultBoardConfiguration = []types.BoardConfiguration{
	// right singles
	types.NewBoardConfiguration(types.Board1x7, types.Orient0, types.Point{X: 0, Y: 12}),
	types.NewBoardConfiguration(types.Board1x7, types.Orient0, types.Point{X: 0, Y: 5}),
	types.NewBoardConfiguration(types.Board1x7, types.Orient180, types.Point{X: 1, Y: 4}),
	types.NewBoardConfiguration(types.Board1x7, types.Orient180, types.Point{X: 1, Y: 11}),
	types.NewBoardConfiguration(types.Board1x7, types.Orient0, types.Point{X: 2, Y: 11}),
	types.NewBoardConfiguration(types.Board1x7, types.Orient0, types.Point{X: 2, Y: 4}),

	// chonkers
	types.NewBoardConfiguration(types.Board7x7, types.Orient90, types.Point{X: 3, Y: 0}),
	types.NewBoardConfiguration(types.Board7x7, types.Orient90, types.Point{X: 3, Y: 7}),
	types.NewBoardConfiguration(types.Board7x7, types.Orient270, types.Point{X: 3, Y: 14}),

	types.NewBoardConfiguration(types.Board7x7, types.Orient90, types.Point{X: 10, Y: 14}),
	types.NewBoardConfiguration(types.Board7x7, types.Orient270, types.Point{X: 10, Y: 7}),
	types.NewBoardConfiguration(types.Board7x7, types.Orient270, types.Point{X: 10, Y: 0}),

	types.NewBoardConfiguration(types.Board7x7, types.Orient90, types.Point{X: 17, Y: 0}),
	types.NewBoardConfiguration(types.Board7x7, types.Orient90, types.Point{X: 17, Y: 7}),
	types.NewBoardConfiguration(types.Board7x7, types.Orient90, types.Point{X: 17, Y: 14}),

	// side singles
	types.NewBoardConfiguration(types.Board1x7, types.Orient0, types.Point{X: 23, Y: 14}),
	types.NewBoardConfiguration(types.Board1x7, types.Orient0, types.Point{X: 23, Y: 7}),
	types.NewBoardConfiguration(types.Board1x7, types.Orient0, types.Point{X: 23, Y: 0}),

	types.NewBoardConfiguration(types.Board1x7, types.Orient180, types.Point{X: 23, Y: 0}),
	types.NewBoardConfiguration(types.Board1x7, types.Orient180, types.Point{X: 23, Y: 7}),
	types.NewBoardConfiguration(types.Board1x7, types.Orient180, types.Point{X: 23, Y: 14}),
}

var DefaultDisallowedPositions = []types.Point{
	types.CreatePoint(0, 5), types.CreatePoint(0, 6),
	types.CreatePoint(0, 15), types.CreatePoint(0, 16), types.CreatePoint(0, 17), types.CreatePoint(0, 18),
	types.CreatePoint(1, 4), types.CreatePoint(1, 17),
	types.CreatePoint(3, 0), types.CreatePoint(3, 1), types.CreatePoint(3, 2),
	types.CreatePoint(3, 19), types.CreatePoint(3, 20),
	types.CreatePoint(4, 0), types.CreatePoint(4, 1),
	types.CreatePoint(4, 20),
	types.CreatePoint(5, 0), types.CreatePoint(5, 1),
	types.CreatePoint(5, 20),
	types.CreatePoint(6, 0),
	types.CreatePoint(7, 0),
	types.CreatePoint(8, 0),
	types.CreatePoint(9, 0),
	types.CreatePoint(10, 0),
	types.CreatePoint(11, 0),
	types.CreatePoint(12, 0),
	types.CreatePoint(13, 0),
	types.CreatePoint(14, 0),
	types.CreatePoint(15, 0),
	types.CreatePoint(16, 0),
	types.CreatePoint(17, 0),
	types.CreatePoint(18, 0),
	types.CreatePoint(19, 0),
	types.CreatePoint(20, 0),
	types.CreatePoint(21, 0),
	types.CreatePoint(22, 0),
	types.CreatePoint(23, 0),
}
