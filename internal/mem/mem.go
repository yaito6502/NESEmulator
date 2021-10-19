package mem

import (
	"errors"
	"log"
)

type MemoryMap struct {
	wram           [0x0800]byte
	wram_mirror    [0x1800]byte
	ppu_reg        [0x0008]byte
	ppu_reg_mirror [0x0008]byte
	apu_ioreg      [0x0020]byte
	ext_rom        [0x1FE0]byte
	ext_ram        [0x2000]byte
	prg_rom        [0x4000]byte
	prg_rom2       [0x4000]byte
}

func NewMemory() *MemoryMap {
	memory := new(MemoryMap)
	return memory
}

func (mem *MemoryMap) assignAccessMemory(address uint16) (*byte, error) {
	switch {
	case address <= 0x07FF:
		return &mem.wram[address], nil
	case address <= 0x1FFF:
		return &mem.wram_mirror[address-0x0800], nil
	case address <= 0x2007:
		return &mem.ppu_reg[address-0x2000], nil
	case address <= 0x3FFF:
		return &mem.ppu_reg_mirror[address-0x2008], nil
	case address <= 0x401F:
		return &mem.apu_ioreg[address-0x4000], nil
	case address <= 0x5FFF:
		return &mem.ext_rom[address-0x4020], nil
	case address <= 0x7FFF:
		return &mem.ext_ram[address-0x6000], nil
	case address <= 0xBFFF:
		return &mem.prg_rom[address-0x8000], nil
	case address <= 0xFFFF:
		return &mem.prg_rom2[address-0xC000], nil
	default:
		return &mem.wram[0x0000], errors.New("address out of range")
	}
}

func (mem *MemoryMap) Fetch(address uint16) byte {
	data, err := mem.assignAccessMemory(address)
	if err != nil {
		log.Fatal(err)
	}
	return *data
}

func (mem *MemoryMap) Store(address uint16, data uint8) {
	ptr, err := mem.assignAccessMemory(address)
	if err != nil {
		log.Fatal(err)
	}
	*ptr = data
}
