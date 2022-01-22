package ppu

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yaito6502/NESEmulator/internal/ppubus"
)

type PPURegister struct {
	ppuCtrl   byte
	ppuMask   byte
	ppuStatus byte
	oamAddr   byte
	oamData   byte
	ppuScroll byte
	ppuAddr   uint16
	ppuData   byte
}

type PPU struct {
	reg   PPURegister
	clock uint16
	line  uint16
	image *ebiten.Image
	bus   *ppubus.PPUBUS
}

func NewPPU(bus *ppubus.PPUBUS) *PPU {
	ppu := new(PPU)
	ppu.bus = bus
	ppu.clock = 21
	ppu.image = ebiten.NewImage(256, 240)
	return ppu
}

func (ppu *PPU) ReadRegister(address uint16) uint8 {
	switch {
	case address == 0x2002:
		return ppu.reg.ppuStatus
	case address == 0x2004:
		return ppu.reg.oamData
	case address == 0x2007:
		return ppu.bus.Read(ppu.reg.ppuAddr)
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
		ppu.WritePPUADDR(data)
	case address == 0x2007:
		ppu.bus.Write(ppu.reg.ppuAddr, data)
		ppu.reg.ppuAddr++
	default:
		log.Fatalf("cannnot write register on address %v", address)
	}
}

func (ppu *PPU) WritePPUADDR(data uint8) {
	ppu.reg.ppuAddr = ppu.reg.ppuAddr<<8 + uint16(data)
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
	x       uint16
	y       uint16
}

var colors = getColorTable()

//Sprite -> Palette -> Color
func (ppu *PPU) getColor(tile *Tile, x uint16, y uint16) color.Color {
	paletteID := tile.sprite[y][x]
	colorID := tile.palette[paletteID]
	return colors[colorID]
}

//Tile -> Pixel
func (ppu *PPU) fillTileInImage(tile *Tile) {
	for y := uint16(0); y < 8; y++ {
		for x := uint16(0); x < 8; x++ {
			ppu.image.Set(int(tile.x*8+x), int(tile.y*8+y), ppu.getColor(tile, x, y))
		}
	}
}

func (ppu *PPU) getSpriteID(tileX, tileY uint16) uint16 {
	spriteID := ppu.bus.Read(0x2000 + tileY*32 + tileX)
	return uint16(spriteID)
}

func (ppu *PPU) getPaletteID(tileX, tileY uint16) uint16 {
	attribute := ppu.bus.Read(0x23C0 + tileY/4*8 + tileX/4)
	shift := (tileX%4)/2 + (tileY%4)/2*2
	paletteID := attribute & (0x03 << (6 - shift*2))
	return uint16(paletteID)
}

func (ppu *PPU) NewSprite(spriteID uint16) *Sprite {
	sprite := new(Sprite)
	for y := uint16(0); y < 8; y++ {
		low := ppu.bus.Read(0x0010*spriteID + y)
		high := ppu.bus.Read(0x0010*spriteID + y + 8)
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

func (ppu *PPU) NewPalette(paletteID uint16) *Palette {
	palette := new(Palette)
	for i := uint16(0); i < 4; i++ {
		palette[i] = ppu.bus.Read(0x3F00 + paletteID*4 + i)
	}
	return palette
}

func (ppu *PPU) fillLineInImage() {
	tile := new(Tile)
	tile.y = ppu.line / 8
	for tile.x = 0; tile.x < 32; tile.x++ {
		spriteID := ppu.getSpriteID(tile.x, tile.y)
		paletteID := ppu.getPaletteID(tile.x, tile.y)
		tile.sprite = ppu.NewSprite(spriteID)
		tile.palette = ppu.NewPalette(paletteID)
		ppu.fillTileInImage(tile)
	}
}

func (ppu *PPU) GetLine() uint16 {
	return ppu.line
}

func (ppu *PPU) GetClock() uint16 {
	return ppu.clock
}

func (ppu *PPU) Run(cycles uint16) *ebiten.Image {
	ppu.clock += cycles

	//ppu.clock[0 ~ 255] -> Draw Display
	//ppu.clock[256 ~ 340] -> Hblank
	for ; ppu.clock >= 341; ppu.clock -= 341 {
		ppu.line++
	}
	//ppu.lines[0 ~ 239] -> Draw Display
	//ppu.lines[240 ~ 261] -> Vblank
	if ppu.line%8 == 0 && ppu.line < 240 {
		ppu.fillLineInImage()
	}
	if ppu.line >= 262 {
		ppu.line = 0
		ppu.image.Clear()
	}
	if ppu.line >= 240 {
		return ppu.image
	}
	return nil
}
