package ppu

import (
	"log"

	"github.com/yaito6502/NESEmulator/internal/ppubus"
)

type PPURegister struct {
	ppuCtrl   byte
	ppuMask   byte
	ppuStatus byte
	oamAddr   byte
	oamData   byte
	ppuScroll byte
	ppuAddr   byte
	ppuData   byte
}

type PPU struct {
	reg    PPURegister
	cycles uint16
	lines  uint16
	bus    *ppubus.PPUBUS
}

func NewPPU(bus *ppubus.PPUBUS) *PPU {
	ppu := new(PPU)
	ppu.bus = bus
	return ppu
}

func (ppu *PPU) ReadRegister(address uint16) uint8 {
	switch {
	case address == 0x2002:
		return ppu.reg.ppuStatus
	case address == 0x2004:
		return ppu.reg.oamData
	case address == 0x2007:
		return ppu.reg.ppuData
	default:
		log.Fatalf("cannnot read register on address %v", address)
		return 0
	}
}

func (ppu *PPU) WriteRegister(address uint16, data uint8) {
	switch {
	case address == 0x2000:
		ppu.reg.ppuCtrl = data
	case address == 0x2001:
		ppu.reg.ppuMask = data
	case address == 0x2003:
		ppu.reg.oamAddr = data
	case address == 0x2004:
		ppu.reg.oamData = data
	case address == 0x2005:
		ppu.reg.ppuScroll = data
	case address == 0x2006:
		ppu.reg.ppuAddr = data
	case address == 0x2007:
		ppu.reg.ppuData = data
	default:
		log.Fatalf("cannnot write register on address %v", address)
	}
}

func (ppu *PPU) Run(cycles uint16) {
	ppu.cycles += cycles

	//ppu.cycles[0 ~ 255] -> Draw Display
	//ppu.cycles[256 ~ 340] -> Hblank
	if ppu.cycles >= 341 {
		ppu.lines++
	}

	//ppu.lines[0 ~ 239] -> Draw Display
	//ppu.lines[240 ~ 261] -> Vblank
	if ppu.lines >= 262 {
		ppu.lines = 0
	}
}
