package types

import "math"

type Light struct {
	Position Point
	Pixel  	 int
	Show     bool
	Color    Color
}

func MakeGammaTable(gamma float64) []byte {
	gt := make([]byte, 256)
	gmdMax := math.Pow(255, gamma)
	for i := 0; i < 256; i ++ {
		gmd := math.Pow(float64(i), gamma)
		gmdNorm := math.Round(gmd / gmdMax * 255.0)
		gt[i] = byte(math.Min(255.0, gmdNorm))
	}
	return gt
}