package ppubus

import (
	"log"

	"github.com/yaito6502/NESEmulator/internal/mem"
)

type PPUBUS struct {
	patternTable *mem.ROM
	palette      *mem.RAM
	vram         *mem.RAM
}

func NewPPUBUS(crom *mem.ROM, palette *mem.RAM, vram *mem.RAM) *PPUBUS {
	bus := new(PPUBUS)
	bus.patternTable = crom
	bus.palette = palette
	bus.vram = vram
	return bus
}

/*
PPU Memory Map
Address       | Size   | Use
-------------------------------------
0x0000-0x0FFF | 0x1000 | PatternTable0
0x1000-0x1FFF | 0x1000 | PatternTable1
0x2000-0x23BF | 0x03C0 | NameTable0
0x23C0-0x23FF | 0x0040 | AttributeTable0
0x2400-0x27BF | 0x03C0 | NameTable1
0x27C0-0x27FF | 0x0040 | AttributeTable1
0x2800-0x2BBF | 0x03C0 | NameTable2
0x2BC0-0x2BFF | 0x0040 | AttributeTable2
0x2C00-0x2FBF | 0x03C0 | NameTable3
0x2FC0-0x2FFF | 0x0040 | AttributeTable3
0x3000-0x3EFF |        | 0x2000-0x2EFF Mirror
0x3F00-0x3F0F | 0x0010 | BackGroundPalete
0x3F10-0x3F1F | 0x0010 | SpritePalette
0x3F20-0x3FFF |        | 0x3F00-0x3F1F Mirror
--------------------------------------
*/

func (bus *PPUBUS) Read(address uint16) uint8 {
	switch {
	case address <= 0x1FFF:
		return bus.patternTable.Read(address)
	case address <= 0x3EFF:
		if address >= 0x3000 {
			address -= 0x1000
		}
		//Horizontal Mirror
		if address >= 0x2800 && address <= 0x2BFF || address >= 0x2C00 && address <= 0x2FFF {
			address -= 0x0400
		}
		return bus.vram.Read(address - 0x2000)
	case address <= 0x3FFF:
		if address == 0x3F04 || address == 0x3F08 || address == 0x0C {
			address = 0x3F00
		}
		if address == 0x3F10 || address == 0x3F14 || address == 0x3F18 || address == 0x3F1C {
			address -= 0x0010
		}
		return bus.palette.Read(address - 0x3F00)
	default:
		log.Fatalf("address out of range %x", address)
		return 0
	}
}

func (bus *PPUBUS) Write(address uint16, data uint8) {
	switch {
	case address <= 0x1FFF:
		break
	case address <= 0x3EFF:
		if address >= 0x3000 {
			address -= 0x1000
		}
		//Horizontal Mirror
		if address >= 0x2800 && address <= 0x2BFF || address >= 0x2C00 && address <= 0x2FFF {
			address -= 0x0400
		}
		bus.vram.Write(address-0x2000, data)
	case address <= 0x3FFF:
		if address == 0x3F04 || address == 0x3F08 || address == 0x0C {
			address = 0x3F00
		}
		if address == 0x3F10 || address == 0x3F14 || address == 0x3F18 || address == 0x3F1C {
			address -= 0x0010
		}
		bus.palette.Write(address - 0x3F00, data)
	default:
		log.Fatalf("address out of range %x", address)
	}
}
