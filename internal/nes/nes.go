package nes

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/yaito6502/NESEmulator/internal/bus"
	"github.com/yaito6502/NESEmulator/internal/cpu"
)

type NES struct {
	CPU *cpu.CPU
	BUS *bus.BUS
	//PPU *ppu
	//APU *apu
	//WRAM *wram
	//VRAM *vram
	//DMA *dma
	//PAD *pad
}

func NewNES() *NES {
	nes := new(NES)
	//[TODO]cpu does not have memory
	//busを通して、メモリマップにアクセスするように要修正
	prom, _ := nes.AttachCartridge("../sample1.nes")
	nes.BUS = bus.NewBUS(prom)
	nes.CPU = cpu.NewCPU(nes.BUS)
	return nes
}

//promとcromをwram領域にコピーする必要がある
func (nes *NES) AttachCartridge(filename string) ([]byte, []byte) {
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
	return program_rom, character_rom
}

func (nes *NES) Run() {
	for {
		fmt.Print(" ", nes.CPU.Run(), "\n")
		time.Sleep(time.Millisecond * 100)
	}
}
