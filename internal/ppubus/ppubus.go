package ppubus

import (
	"log"

	"github.com/yaito6502/NESEmulator/internal/mem"
)

type PPUBUS struct {
	patternTable *mem.ROM
	vram         *mem.RAM
}

func NewPPUBUS(crom *mem.ROM, vram *mem.RAM) *PPUBUS {
	bus := new(PPUBUS)
	bus.patternTable = crom
	bus.vram = vram
	return bus
}

func (bus *PPUBUS) Read(address uint16) uint8 {
	switch {
	case address <= 0x1FFF:
		return bus.patternTable.Read(address)
	case address <= 0x3FFF:
		return bus.vram.Read(address - 0x2000)
	default:
		log.Fatalf("address out of range %v", address)
		return 0
	}
}

func (bus *PPUBUS) Write(address uint16, data uint8) {
	switch {
	case address <= 0x1FFF:

	case address <= 0x3FFF:
		bus.vram.Write(address-0x2000, data)
	default:
		log.Fatalf("address out of range %v", address)
	}
}
