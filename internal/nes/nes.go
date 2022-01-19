package nes

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/yaito6502/NESEmulator/internal/bus"
	"github.com/yaito6502/NESEmulator/internal/cpu"
	"github.com/yaito6502/NESEmulator/internal/mem"
)

type NES struct {
	CPU *cpu.CPU
	BUS *bus.BUS
	//PPU *ppu
	//APU *apu
	WRAM mem.RAM
	//VRAM *vram
	//DMA *dma
	//PAD *pad
}

func NewNES() *NES {
	nes := new(NES)
	prom, _ := nes.AttachCartridge("../sample1.nes")
	nes.WRAM = mem.NewRAM(0x0800)
	nes.BUS = bus.NewBUS(&nes.WRAM, &prom)
	nes.CPU = cpu.NewCPU(nes.BUS)
	return nes
}

func (nes *NES) AttachCartridge(filename string) (mem.ROM, mem.ROM) {
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
