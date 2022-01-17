package nes

import (
	"fmt"
	"time"

	"github.com/yaito6502/NESEmulator/internal/cpu"
	"github.com/yaito6502/NESEmulator/internal/mem"
)

type NES struct {
	CPU *cpu.CPU
	//PPU *ppu
	//APU *apu
	//WRAM *wram
	//VRAM *vram
	//DMA *dma
	//PAD *pad
}

func NewNES() *NES {
	nes := new(NES)
	nes.CPU = cpu.NewCPU(mem.NewMemory())
	return nes
}

func (nes *NES) AttachCartridge(filename string) {
	//.nes拡張子のチェック
	//ファイルからpromとcromを読み込み、スライスに保管
	//promとcromをcpuのメモリマップに配置する
}

func (nes *NES) Run() {
	//
	/*cart := cartridge.NewCartridge("../sample1.nes")
	for i, b := range cart.ProgramRom {
		mem.Store(uint16(0x8000+i), b)
	}*/

	for {
		fmt.Print(" ", nes.CPU.Run(), "\n")
		time.Sleep(time.Millisecond * 100)
	}
}
