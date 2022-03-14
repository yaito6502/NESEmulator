package ppu

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/yaito6502/NESEmulator/internal/cpudebug"
	"github.com/yaito6502/NESEmulator/internal/ppubus"
	"github.com/yaito6502/NESEmulator/pkg"
)

type PPUCtrl struct {
	bits                       byte
	nameTableAddr              uint16
	vramAddrInc                uint16
	spritePatternTableAddr     uint16
	backgroundPatternTableAddr uint16
	spriteSize                 uint8
	vblankNMI                  bool
}

type PPUMASK struct {
	bits           byte
	greyScale      bool
	backgroundMask bool
	spriteMask     bool
	showBackground bool
	showSprite     bool
	emphasizeRed   bool
	emphasizeGreen bool
	emphasizeBlue  bool
}

type PPURegister struct {
	ppuCtrl   PPUCtrl
	ppuMask   PPUMASK
	ppuStatus byte
	oamAddr   byte
	oamData   byte
	ppuScroll byte
	ppuAddr   uint16
	ppuData   byte
}

type PPU struct {
	reg        PPURegister
	clock      uint16
	line       uint16
	bus        *ppubus.PPUBUS
	background *ebiten.Image
	info       *cpudebug.DebugInfo
}

const (
	WIDTH  = 256
	HEIGHT = 240
)

func NewPPU(bus *ppubus.PPUBUS, info *cpudebug.DebugInfo) *PPU {
	ppu := new(PPU)
	ppu.bus = bus
	ppu.background = ebiten.NewImage(WIDTH, HEIGHT)
	ppu.info = info
	return ppu
}

func (ppu *PPU) ReadRegister(address uint16) uint8 {
	switch {
	case address == 0x2002:
		return ppu.reg.ppuStatus
	case address == 0x2004:
		return ppu.reg.oamData
	case address == 0x2007:
		return ppu.ReadPPUData()
		//default:
		//	log.Fatalf("cannot read register on address %#x", address)
	}
	return 0
}

func (ppu *PPU) ReadPPUData() byte {
	data := ppu.reg.ppuData
	ppu.reg.ppuData = ppu.bus.Read(ppu.reg.ppuAddr)
	ppu.reg.ppuAddr += ppu.reg.ppuCtrl.vramAddrInc
	if ppu.reg.ppuAddr >= 0x3F00 {
		return ppu.reg.ppuData
	}
	return data
}

func (ppu *PPU) WriteRegister(address uint16, data uint8) {
	switch {
	case address == 0x2000:
		ppu.reg.ppuCtrl.WritePPUCTRL(data)
	case address == 0x2001:
		ppu.reg.ppuMask.WritePPUMASK(data)
	case address == 0x2003:
		ppu.reg.oamAddr = data
	case address == 0x2004:
		ppu.reg.WriteOAMDATA(data)
	case address == 0x2005:
		ppu.reg.ppuScroll = data
	case address == 0x2006:
		ppu.reg.WritePPUADDR(data)
	case address == 0x2007:
		ppu.WritePPUDATA(data)
		//default:
		//	log.Fatalf("cannnot write register on address %#x", address)
	}
}

func (reg *PPUCtrl) WritePPUCTRL(data uint8) {
	reg.bits = data
	//(0 = $2000; 1 = $2400; 2 = $2800; 3 = $2C00)
	reg.nameTableAddr = 0x2000 + uint16(reg.bits&0x03)*0x0400
	if pkg.Uint8tob(reg.bits & 0x04) {
		reg.vramAddrInc += 32
	} else {
		reg.vramAddrInc += 1
	}
	if pkg.Uint8tob(reg.bits & 0x08) {
		reg.spritePatternTableAddr = 0x1000
	} else {
		reg.spritePatternTableAddr = 0x0000
	}
	if pkg.Uint8tob(reg.bits & 0x10) {
		reg.backgroundPatternTableAddr = 0x1000
	} else {
		reg.backgroundPatternTableAddr = 0x0000
	}
	reg.spriteSize = 8 + (reg.bits&0x20)*8
	reg.vblankNMI = (reg.bits & 0x80) == 1
}

func (reg *PPUMASK) WritePPUMASK(data uint8) {
	reg.bits = data
	reg.greyScale = pkg.Uint8tob(reg.bits & 0x01)
	reg.backgroundMask = pkg.Uint8tob(reg.bits & 0x02)
	reg.spriteMask = pkg.Uint8tob(reg.bits & 0x04)
	reg.showBackground = pkg.Uint8tob(reg.bits & 0x08)
	reg.showSprite = pkg.Uint8tob(reg.bits & 0x10)
	reg.emphasizeRed = pkg.Uint8tob(reg.bits & 0x20)
	reg.emphasizeGreen = pkg.Uint8tob(reg.bits & 0x40)
	reg.emphasizeBlue = pkg.Uint8tob(reg.bits & 0x80)
}

func (reg *PPURegister) WriteOAMDATA(data uint8) {
	reg.oamData = data
	//oam.Write(reg.oamAddr, reg.oamData)
	reg.oamAddr++
}

func (reg *PPURegister) WritePPUADDR(data uint8) {
	reg.ppuAddr = reg.ppuAddr<<8 + uint16(data)
}

func (ppu *PPU) WritePPUDATA(data uint8) {
	ppu.reg.ppuData = data
	ppu.bus.Write(ppu.reg.ppuAddr, ppu.reg.ppuData)
	ppu.reg.ppuAddr += ppu.reg.ppuCtrl.vramAddrInc
}

func getColorTable() []color.RGBA {
	return []color.RGBA{
		{0x80, 0x80, 0x80, 0xFF}, {0x00, 0x3D, 0xA6, 0xFF}, {0x00, 0x12, 0xB0, 0xFF}, {0x44, 0x00, 0x96, 0xFF},
		{0xA1, 0x00, 0x5E, 0xFF}, {0xC7, 0x00, 0x28, 0xFF}, {0xBA, 0x06, 0x00, 0xFF}, {0x8C, 0x17, 0x00, 0xFF},
		{0x5C, 0x2F, 0x00, 0xFF}, {0x10, 0x45, 0x00, 0xFF}, {0x05, 0x4A, 0x00, 0xFF}, {0x00, 0x47, 0x2E, 0xFF},
		{0x00, 0x41, 0x66, 0xFF}, {0x00, 0x00, 0x00, 0xFF}, {0x05, 0x05, 0x05, 0xFF}, {0x05, 0x05, 0x05, 0xFF},
		{0xC7, 0xC7, 0xC7, 0xFF}, {0x00, 0x77, 0xFF, 0xFF}, {0x21, 0x55, 0xFF, 0xFF}, {0x82, 0x37, 0xFA, 0xFF},
		{0xEB, 0x2F, 0xB5, 0xFF}, {0xFF, 0x29, 0x50, 0xFF}, {0xFF, 0x22, 0x00, 0xFF}, {0xD6, 0x32, 0x00, 0xFF},
		{0xC4, 0x62, 0x00, 0xFF}, {0x35, 0x80, 0x00, 0xFF}, {0x05, 0x8F, 0x00, 0xFF}, {0x00, 0x8A, 0x55, 0xFF},
		{0x00, 0x99, 0xCC, 0xFF}, {0x21, 0x21, 0x21, 0xFF}, {0x09, 0x09, 0x09, 0xFF}, {0x09, 0x09, 0x09, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF}, {0x0F, 0xD7, 0xFF, 0xFF}, {0x69, 0xA2, 0xFF, 0xFF}, {0xD4, 0x80, 0xFF, 0xFF},
		{0xFF, 0x45, 0xF3, 0xFF}, {0xFF, 0x61, 0x8B, 0xFF}, {0xFF, 0x88, 0x33, 0xFF}, {0xFF, 0x9C, 0x12, 0xFF},
		{0xFA, 0xBC, 0x20, 0xFF}, {0x9F, 0xE3, 0x0E, 0xFF}, {0x2B, 0xF0, 0x35, 0xFF}, {0x0C, 0xF0, 0xA4, 0xFF},
		{0x05, 0xFB, 0xFF, 0xFF}, {0x5E, 0x5E, 0x5E, 0xFF}, {0x0D, 0x0D, 0x0D, 0xFF}, {0x0D, 0x0D, 0x0D, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF}, {0xA6, 0xFC, 0xFF, 0xFF}, {0xB3, 0xEC, 0xFF, 0xFF}, {0xDA, 0xAB, 0xEB, 0xFF},
		{0xFF, 0xA8, 0xF9, 0xFF}, {0xFF, 0xAB, 0xB3, 0xFF}, {0xFF, 0xD2, 0xB0, 0xFF}, {0xFF, 0xEF, 0xA6, 0xFF},
		{0xFF, 0xF7, 0x9C, 0xFF}, {0xD7, 0xE8, 0x95, 0xFF}, {0xA6, 0xED, 0xAF, 0xFF}, {0xA2, 0xF2, 0xDA, 0xFF},
		{0x99, 0xFF, 0xFC, 0xFF}, {0xDD, 0xDD, 0xDD, 0xFF}, {0x11, 0x11, 0x11, 0xFF}, {0x11, 0x11, 0x11, 0xFF},
	}
}

type Palette [4]byte
type Sprite [8][8]byte

type Tile struct {
	sprite  *Sprite
	palette *Palette
}

const (
	TILE_ROW = 30
	TILE_COL = 32
)

var colors = getColorTable()

//Sprite -> Palette -> Color
func (ppu *PPU) getColor(tile *Tile, x uint16, y uint16) color.Color {
	paletteID := tile.sprite[y][x]
	if paletteID == 0 {
		return colors[ppu.bus.Read(0x3F10)]
	}
	colorID := tile.palette[paletteID]
	return colors[colorID]
}

//Tile -> Pixel
func (ppu *PPU) fillTileInImage(tile *Tile, tilex, tiley uint16) {
	for y := uint16(0); y < 8; y++ {
		for x := uint16(0); x < 8; x++ {
			ppu.background.Set(int(tilex*8+x), int(tiley*8+y), ppu.getColor(tile, x, y))
		}
	}
}

/*
func (ppu *PPU) getBlockID(tileX, tileY uint16) byte {
	return uint8((tileX%4)/2 + (tileY%4)/2*2)
}

func (ppu *PPU) getAttribute(tileX, tileY, baseAddr uint16) byte {
	return ppu.bus.Read(baseAddr + 0x03C0 + tileX/4 + tileY/4*8)
}*/

func (ppu *PPU) getSpriteID(tileX, tileY uint16) uint16 {
	//0x2000, 0x2400, 0x2800, 0x3200 NameTable
	baseAddr := ppu.reg.ppuCtrl.nameTableAddr
	spriteID := ppu.bus.Read(baseAddr + tileY*TILE_COL + tileX)
	return uint16(spriteID)
}

func (ppu *PPU) getPaletteID(tileX, tileY uint16) uint16 {
	//0x23C0, 0x27C0, 0x2BC0, 0x2FC0 Attribute Table
	baseAddr := ppu.reg.ppuCtrl.nameTableAddr + uint16(0x03C0)
	blockx := tileX / 2
	blocky := tileY / 2
	attribute := ppu.bus.Read(baseAddr + blocky/2*8 + blockx/2)
	offsetx := blockx % 2
	offsety := blocky % 2
	blockID := offsety*2 + offsetx
	paletteID := (attribute >> uint8(blockID*2)) & 0x03
	return uint16(paletteID)
}

//background
func (ppu *PPU) NewSprite(spriteID, baseAddr uint16) *Sprite {
	sprite := new(Sprite)
	for y := uint16(0); y < 8; y++ {
		low := ppu.bus.Read(baseAddr + 0x0010*spriteID + y)
		high := ppu.bus.Read(baseAddr + 0x0010*spriteID + y + 8)
		for x := 0; x < 8; x++ {
			if (high & (1 << (8 - x))) != 0 {
				sprite[y][x] += 2
			}
			if (low & (1 << (8 - x))) != 0 {
				sprite[y][x] += 1
			}
		}
	}
	return sprite
}

//background
func (ppu *PPU) NewPalette(paletteID, baseAddr uint16) *Palette {
	palette := new(Palette)
	for i := uint16(0); i < 4; i++ {
		palette[i] = ppu.bus.Read(baseAddr + paletteID*4 + i)
	}
	return palette
}

func (ppu *PPU) PlaceTile(x, y uint16) {
	tile := new(Tile)
	spriteID := ppu.getSpriteID(x, y)
	paletteID := ppu.getPaletteID(x, y)
	tile.sprite = ppu.NewSprite(spriteID, ppu.reg.ppuCtrl.backgroundPatternTableAddr)
	tile.palette = ppu.NewPalette(paletteID, 0x3F00)
	ppu.fillTileInImage(tile, x, y)
}

func (ppu *PPU) fillBackGround() {
	tiley := ppu.line / 8
	for tilex := uint16(0); tilex < TILE_COL; tilex++ {
		ppu.PlaceTile(tilex, tiley)
	}
}

func (ppu *PPU) setVBlank() {
	ppu.reg.ppuStatus |= (1 << 7)
	/*if ppu.reg.ppuCtrl.vblankNMI {
		cpu.NMI()
	}*/
}

func (ppu *PPU) unsetVBlank() {
	ppu.reg.ppuStatus ^= (1 << 7)
}

func (ppu *PPU) Run(cycles uint16) *ebiten.Image {
	ppu.info.PPUX = ppu.clock
	ppu.info.PPUY = ppu.line
	ppu.clock += cycles

	if ppu.line == 0 {
		ppu.background.Clear()
	}
	//ppu.clock[0 ~ 255] -> Draw Display
	//ppu.clock[256 ~ 340] -> Hblank
	if ppu.clock >= 341 {
		ppu.line++
		ppu.clock -= 341
	}
	//ppu.lines[0 ~ 239] -> Draw Display
	//ppu.lines[240 ~ 261] -> Vblank
	if ppu.line%8 == 0 && ppu.line <= 240 {
		ppu.fillBackGround()
	}
	if ppu.line >= 241 && ppu.line <= 261 {
		ppu.setVBlank()
	}
	if ppu.line == 262 {
		ppu.line = 0
		ppu.unsetVBlank()
		return ppu.background
	}
	return nil
}
