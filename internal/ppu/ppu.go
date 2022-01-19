package ppu

import "github.com/yaito6502/NESEmulator/internal/bus"

type PPU struct {
	ppuCtrl   byte
	ppuMask   byte
	ppuStatus byte
	oamAddr   byte
	oamData   byte
	ppuScroll byte
	ppuAddr   byte
	ppuData   byte
	bus       *bus.PPUBUS
}

func NewPPU(bus *bus.PPUBUS) *PPU {
	ppu := new(PPU)
	ppu.bus = bus
	return ppu
}
