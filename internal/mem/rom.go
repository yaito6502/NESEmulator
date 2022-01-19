package mem

type ROM []uint8

func NewROM(mem []uint8) ROM {
	rom := make([]uint8, len(mem))
	copy(rom, mem)
	return rom
}

func (rom *ROM) Reset() {
	*rom = ROM{0}
}

func (rom *ROM) Read(address uint16) uint8 {
	return (*rom)[address]
}
