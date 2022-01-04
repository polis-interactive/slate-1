package types

type Color struct {
	R uint8
	G uint8
	B uint8
	W uint8
}

func (c *Color) ToBits() (out uint32) {
	return uint32(c.W) << 24 | uint32(c.R) << 16 | uint32(c.G) << 8 | uint32(c.B)
}