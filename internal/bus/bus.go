package bus

import (
	"errors"
	"log"
)

type BUS struct {
	//試験的にbyte列のみでCPUを動かす
	//今後、メモリマップに従って他のパッケージをbusが持っておく
	other[0x8000]byte
	prgRom []byte
	//prgRom2 []byte
}

func NewBUS(prgRom []byte) *BUS {
	bus := new(BUS)
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
func (bus *BUS) assignAccessMemory(address uint16) (*byte, error) {
	switch {
	/*case address <= 0x07FF:
	case address <= 0x1FFF:
	case address <= 0x2007:
	case address <= 0x3FFF:
	case address <= 0x401F:
	case address <= 0x5FFF:*/
	case address <= 0x7FFF:
		return &bus.other[address], nil
	case address <= 0xBFFF:
		return &bus.prgRom[address-0x8000], nil
	case address <= 0xFFFF:
		return &bus.prgRom[address-0x8000], nil
	default:
		return nil, errors.New("address out of range")
	}
}

func (bus *BUS) Read(address uint16) byte {
	ptr, err := bus.assignAccessMemory(address)
	if err != nil {
		log.Fatal()
	}
	return *ptr
}

func (bus *BUS) Write(address uint16, data uint8) {
	ptr, err := bus.assignAccessMemory(address)
	if err != nil {
		log.Fatal()
	}
	*ptr = data
}
