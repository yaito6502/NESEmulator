package bus

import (
	"log"

	"github.com/yaito6502/NESEmulator/internal/mem"
)

type BUS struct {
	wram       *mem.RAM
	wramMirror *mem.RAM
	ppuReg     [0x0008]byte
	apuIOPad   [0x0020]byte
	extRom     *mem.ROM
	extRam     *mem.RAM
	prgRom     *mem.ROM
}

func NewBUS(wram *mem.RAM, prgRom *mem.ROM) *BUS {
	bus := new(BUS)
	bus.wram = wram
	bus.wramMirror = wram
	bus.prgRom = prgRom
	return bus
}

/*
CPU Memory Map
Address       | Size   | Use
-------------------------------------
0x0000-0x07FF | 0x0800 | WRAM
0x0800-0x1FFF |        | WRAM mirror
0x2000-0x2007 | 0x0008 | PPU Register
0x2008-0x3FFF |        | PPU Register mirror
0x4000-0x401F | 0x0020 | APU I/O, PAD
0x4020-0x5FFF | 0x1FE0 | Extended ROM
0x6000-0x7FFF | 0x2000 | Extended RAM
0x8000-0xBFFF | 0x4000 | PRG-ROM
0xC000-0xFFFF | 0x4000 | PRG-ROM
--------------------------------------
*/

func (bus *BUS) Read(address uint16) byte {
	switch {
	case address <= 0x07FF:
		return bus.wram.Read(address)
	case address <= 0x1FFF:
		return bus.wramMirror.Read(address - 0x0800)
	case address <= 0x2007:
		return bus.ppuReg[address-0x2000]
	case address <= 0x3FFF:
		return bus.ppuReg[(address-0x2008)%8]
	case address <= 0x401F:
		return bus.apuIOPad[address-0x4000]
	case address <= 0x5FFF:
		return bus.extRom.Read(address - 0x4020)
	case address <= 0x7FFF:
		return bus.extRam.Read(address)
	case address <= 0xBFFF:
		return bus.prgRom.Read(address - 0x8000)
	case address <= 0xFFFF:
		return bus.prgRom.Read(address - 0x8000)
	default:
		log.Fatalf("address out of range %v", address)
		return 0
	}
}

func (bus *BUS) Write(address uint16, data uint8) {
	switch {
	case address <= 0x07FF:
		bus.wram.Write(address, data)
	case address <= 0x1FFF:
		bus.wramMirror.Write(address-0x0800, data)
	case address <= 0x2007:
		bus.ppuReg[address-0x2000] = data
	case address <= 0x3FFF:
		bus.ppuReg[(address-0x2008)%8] = data
	case address <= 0x401F:
		bus.apuIOPad[address-0x4000] = data
	case address <= 0x5FFF:

	case address <= 0x7FFF:
		bus.extRam.Write(address, data)
	case address <= 0xBFFF:

	case address <= 0xFFFF:

	default:
		log.Fatalf("address out of range %v", address)
	}
}
