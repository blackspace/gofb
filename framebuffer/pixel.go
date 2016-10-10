package framebuffer

type Pixel struct {
	A byte
	R byte
	G byte
	B byte
}

func (p Pixel)ToU32() uint32 {
	return uint32(p.B)+uint32(p.G)<<8+uint32(p.R)<<16+uint32(p.A)<<24
}
