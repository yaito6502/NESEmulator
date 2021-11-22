package ram

type Ram []uint8

func NewRam(size uint16) Ram {
	return make([]uint8, size)
}

func (ram *Ram) reset() {
	*ram = Ram{0}
}

func (ram Ram) read(address uint16) uint8 {
	return ram[address]
}

func (ram *Ram) write(address uint16, data uint8) {
	(*ram)[address] = data
}
