package cpubus

import (
	"log"

	"github.com/yaito6502/NESEmulator/internal/mem"
	"github.com/yaito6502/NESEmulator/internal/pad"
	"github.com/yaito6502/NESEmulator/internal/ppu"
)

type CPUBUS struct {
	wram *mem.RAM
	ppu  *ppu.PPU
	pad  *pad.PAD
	apu    [0x0020]byte //apu *apu.APU
	extRom [0x1FE0]byte
	extRam [0x2000]byte
	prgRom *mem.ROM
}

func NewCPUBUS(wram *mem.RAM, ppu *ppu.PPU, prgRom *mem.ROM, pad *pad.PAD) *CPUBUS {
	bus := new(CPUBUS)
	bus.wram = wram
	bus.ppu = ppu
	bus.pad = pad
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

func (bus *CPUBUS) Read(address uint16) byte {
	switch {
	case address <= 0x1FFF:
		return bus.wram.Read(address % 0x0800)
	case address <= 0x3FFF:
		return bus.ppu.ReadRegister(0x2000 + address%8)
	case address == 0x4016 || address == 0x4017:
		return bus.pad.Read(address % 0x0001)
	case address <= 0x401F:
		return bus.apu[address-0x4000]
	case address <= 0x5FFF:
		return bus.extRom[address-0x4020]
	case address <= 0x7FFF:
		return bus.extRam[address-0x6000]
	case address <= 0xBFFF:
		return bus.prgRom.Read(address - 0x8000)
	case address <= 0xFFFF:
		return bus.prgRom.Read(address - 0xC000)
	default:
		log.Fatalf("address out of range %x", address)
		return 0
	}
}

func (bus *CPUBUS) Write(address uint16, data uint8) {
	switch {
	case address <= 0x1FFF:
		bus.wram.Write(address%0x0800, data)
	case address <= 0x3FFF:
		bus.ppu.WriteRegister(0x2000+address%8, data)
	case address == 0x4014:
		bus.writeDMA(data)
	case address == 0x4016:
		bus.pad.Write(data)
	case address <= 0x401F:
		bus.apu[address-0x4000] = data
	case address <= 0x5FFF:
		break
	case address <= 0x7FFF:
		bus.extRam[address-0x6000] = data
	case address <= 0xBFFF:
		break
	case address <= 0xFFFF:
		break
	default:
		log.Fatalf("address out of range %x", address)
	}
}

func (bus *CPUBUS) writeDMA(data uint8) {
	address := uint16(data) << 8
	bus.ppu.WriteRegister(0x2003, 0x00)
	for i := uint16(0); i < 0x0100; i++ {
		bus.ppu.WriteOAMDATA(bus.Read(address + i))
	}
}
