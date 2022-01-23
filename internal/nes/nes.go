package nes

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yaito6502/NESEmulator/internal/cartridge"
	"github.com/yaito6502/NESEmulator/internal/cpu"
	"github.com/yaito6502/NESEmulator/internal/cpubus"
	"github.com/yaito6502/NESEmulator/internal/cpudebug"
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
	image  *ebiten.Image
	Info   cpudebug.DebugInfo
	cycles uint64
}

func NewNES(cart *cartridge.Cartridge) *NES {
	nes := new(NES)
	nes.WRAM = mem.NewRAM(0x0800)
	nes.VRAM = mem.NewRAM(0x2000)
	nes.PPUBUS = ppubus.NewPPUBUS(&cart.CharacterRom, &nes.VRAM)
	nes.PPU = ppu.NewPPU(nes.PPUBUS, &nes.Info)
	nes.CPUBUS = cpubus.NewCPUBUS(&nes.WRAM, nes.PPU, &cart.ProgramRom)
	nes.CPU = cpu.NewCPU(nes.CPUBUS, &nes.Info)
	return nes
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
		cycle := nes.CPU.Run()
		nes.image = nes.PPU.Run(uint16(cycle) * 3)
		nes.cycles += uint64(cycle)
		nes.Info.CYCLE = nes.cycles
		nes.Info.Print()
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
