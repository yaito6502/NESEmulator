package rom

type Rom []uint8

func NewRom(size uint16) Rom {
	return make([]uint8, size)
}

func (rom *Rom) reset() {
	*rom = Rom{0}
}

func (rom Rom) read(address uint16) uint8 {
	return rom[address]
}
