package nes

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yaito6502/NESEmulator/internal/cartridge"
	"github.com/yaito6502/NESEmulator/internal/cpu"
	"github.com/yaito6502/NESEmulator/internal/cpubus"
	"github.com/yaito6502/NESEmulator/internal/cpudebug"
	"github.com/yaito6502/NESEmulator/internal/interrupts"
	"github.com/yaito6502/NESEmulator/internal/mem"
	"github.com/yaito6502/NESEmulator/internal/pad"
	"github.com/yaito6502/NESEmulator/internal/ppu"
	"github.com/yaito6502/NESEmulator/internal/ppubus"
)

const (
	WIDTH  = 256
	HEIGHT = 240
	SCALE  = 4
)

type NES struct {
	WRAM   mem.RAM
	VRAM   mem.RAM
	CPU    *cpu.CPU
	PPU    *ppu.PPU
	CPUBUS *cpubus.CPUBUS
	PPUBUS *ppubus.PPUBUS
	//APU *apu
	PAD        *pad.PAD
	inter      *interrupts.Interrupts
	background *ebiten.Image
	sprites    *[ppu.SPRITECOUNT]ppu.Sprite
	Info       cpudebug.DebugInfo
	cycles     uint64
}

func NewNES(cart *cartridge.Cartridge) *NES {
	nes := new(NES)
	nes.WRAM = mem.NewRAM(0x0800)
	nes.VRAM = mem.NewRAM(0x0800)
	nes.PAD = pad.NewPAD()
	nes.inter = interrupts.NewInterrupts()
	palette := mem.NewRAM(0x0020)
	nes.PPUBUS = ppubus.NewPPUBUS(&cart.CharacterRom, &palette, &nes.VRAM)
	nes.PPU = ppu.NewPPU(nes.PPUBUS, nes.inter, &palette, &nes.Info)
	nes.CPUBUS = cpubus.NewCPUBUS(&nes.WRAM, nes.PPU, &cart.ProgramRom, nes.PAD)
	nes.CPU = cpu.NewCPU(nes.CPUBUS, nes.inter, &nes.Info)
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
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	nes.PAD.SetPressKeys()
	nes.background = nil
	nes.sprites = nil
	for nes.background == nil && nes.sprites == nil {
		cycle := nes.CPU.Run()
		nes.background, nes.sprites = nes.PPU.Run(uint16(cycle) * 3)
		nes.cycles += uint64(cycle)
		//nes.Info.CYCLE = nes.cycles
		//nes.Info.Print()
	}
	return nil
}

func (nes *NES) Draw(screen *ebiten.Image) {
	op := new(ebiten.DrawImageOptions)
	op.GeoM.Scale(SCALE, SCALE)
	if nes.background != nil {
		nes.DrawBackground(screen, op)
	}
	if nes.sprites != nil {
		nes.DrawSprites(screen)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%0.2f TPS:%0.2f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
}

func (nes *NES) DrawBackground(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	screen.DrawImage(nes.background, op)
}

func (nes *NES) DrawSprites(screen *ebiten.Image) {
	op := new(ebiten.DrawImageOptions)
	for i := 0; i < ppu.SPRITECOUNT; i++ {
		sprite := nes.sprites[ppu.SPRITECOUNT-(i+1)]
		op.GeoM.Translate(float64(sprite.X), float64(sprite.Y))
		op.GeoM.Scale(SCALE, SCALE)
		screen.DrawImage(sprite.Image, op)
		op.GeoM.Reset()
	}
}

func (nes *NES) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
