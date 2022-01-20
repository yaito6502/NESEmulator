package nes

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yaito6502/NESEmulator/internal/cpu"
	"github.com/yaito6502/NESEmulator/internal/cpubus"
	"github.com/yaito6502/NESEmulator/internal/mem"
	"github.com/yaito6502/NESEmulator/internal/ppu"
	"github.com/yaito6502/NESEmulator/internal/ppubus"
)

const WIDTH = 256
const HEIGHT = 240
const SCALE = 4

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
	image *ebiten.Image
}

func NewNES() *NES {
	nes := new(NES)
	prom, crom := nes.attachCartridge("../sample1.nes")
	nes.WRAM = mem.NewRAM(0x0800)
	nes.VRAM = mem.NewRAM(0x2000)
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
	fmt.Println(len(program_rom), len(character_rom))
	return mem.NewROM(program_rom), mem.NewROM(character_rom)
}

func (nes *NES) Run() {
	ebiten.SetWindowSize(WIDTH*SCALE, HEIGHT*SCALE)
	ebiten.SetWindowTitle("GO NES")
	if err := ebiten.RunGame(nes); err != nil {
		log.Fatal(err)
	}
}

func (nes *NES) Update() error {
	nes.image = nil
	for nes.image == nil {
		cycles := nes.CPU.Run()
		nes.image = nes.PPU.Run(uint16(cycles) * 3)
	}
	return nil
}

func (nes *NES) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(SCALE, SCALE)
	screen.DrawImage(nes.image, op)
	//for debug
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%0.2f", ebiten.CurrentFPS()))
}

func (nes *NES) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
