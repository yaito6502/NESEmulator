package nes

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/yaito6502/NESEmulator/internal/cpu"
	"github.com/yaito6502/NESEmulator/internal/cpubus"
	"github.com/yaito6502/NESEmulator/internal/mem"
	"github.com/yaito6502/NESEmulator/internal/ppu"
	"github.com/yaito6502/NESEmulator/internal/ppubus"
)

type NES struct {
	CPU    *cpu.CPU
	PPU    *ppu.PPU
	CPUBUS *cpubus.CPUBUS
	PPUBUS *ppubus.PPUBUS
	WRAM   mem.RAM
	VRAM   mem.RAM
	//APU *apu
	//DMA *dma
	//PAD *pad
}

func NewNES() *NES {
	nes := new(NES)
	prom, crom := nes.attachCartridge("../sample1.nes")
	nes.WRAM = mem.NewRAM(0x0800)
	nes.VRAM = mem.NewRAM(0x0800)
	nes.PPUBUS = ppubus.NewPPUBUS(&crom, &nes.VRAM)
	nes.PPU = ppu.NewPPU(nes.PPUBUS)
	nes.CPUBUS = cpubus.NewCPUBUS(&nes.WRAM, nes.PPU, &prom)
	nes.CPU = cpu.NewCPU(nes.CPUBUS)
	return nes
}

func (nes *NES) attachCartridge(filename string) (mem.ROM, mem.ROM) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal()
	}

	//check nes format
	if string(contents[0:3]) != "NES" {
		log.Fatal()
	}

	const NESHEADERSIZE int = 0x0010

	character_romstart := NESHEADERSIZE + 0x4000*int(contents[4])
	character_romend := character_romstart + 0x2000*int(contents[5])

	program_rom := contents[NESHEADERSIZE : character_romstart-1]
	character_rom := contents[character_romstart : character_romend-1]
	return mem.NewROM(program_rom), mem.NewROM(character_rom)
}

func (nes *NES) Run() {
	for {
		fmt.Print(" ", nes.CPU.Run(), "\n")
		time.Sleep(time.Millisecond * 100)
	}
}
