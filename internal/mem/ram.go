package mem

type RAM []uint8

func NewRAM(size uint16) RAM {
	return make([]uint8, size)
}

func (ram *RAM) Reset() {
	*ram = RAM{0}
}

func (ram *RAM) Read(address uint16) uint8 {
	return (*ram)[address]
}

func (ram *RAM) Write(address uint16, data uint8) {
	(*ram)[address] = data
}
