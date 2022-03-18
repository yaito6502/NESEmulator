package ppu

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/yaito6502/NESEmulator/internal/cpudebug"
	"github.com/yaito6502/NESEmulator/internal/interrupts"
	"github.com/yaito6502/NESEmulator/internal/mem"
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

const (
	SPRITECOUNT = 64
	SPRITESIZE  = 4
)

type PPU struct {
	reg        PPURegister
	clock      uint16
	scanline   uint16
	bus        *ppubus.PPUBUS
	inter      *interrupts.Interrupts
	sprites    [SPRITECOUNT]Sprite
	background *ebiten.Image
	info       *cpudebug.DebugInfo
	oam        mem.RAM
	palette    *mem.RAM
}

const (
	WIDTH  = 256
	HEIGHT = 240
)

func NewPPU(bus *ppubus.PPUBUS, inter *interrupts.Interrupts, palette *mem.RAM, info *cpudebug.DebugInfo) *PPU {
	ppu := new(PPU)
	ppu.bus = bus
	ppu.inter = inter
	for i := 0; i < SPRITECOUNT; i++ {
		ppu.sprites[i].Image = ebiten.NewImage(8, 8)
	}
	ppu.background = ebiten.NewImage(WIDTH, HEIGHT)
	ppu.info = info
	ppu.oam = mem.NewRAM(0x0100)
	ppu.palette = palette
	return ppu
}

func (ppu *PPU) ReadRegister(address uint16) uint8 {
	switch {
	case address == 0x2002:
		ppu.unsetVBlank()
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
		ppu.WriteOAMDATA(data)
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
		reg.vramAddrInc = 0x20
	} else {
		reg.vramAddrInc = 0x01
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
	if pkg.Uint8tob(reg.bits & 0x20) {
		reg.spriteSize = 0x10
	} else {
		reg.spriteSize = 0x08
	}
	if pkg.Uint8tob(reg.bits & 0x20) {
		reg.spriteSize = 0x10
	} else {
		reg.spriteSize = 0x08
	}
	if pkg.Uint8tob(reg.bits & 0x80) {
		reg.vblankNMI = true
	} else {
		reg.vblankNMI = false
	}
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

func (ppu *PPU) WriteOAMDATA(data uint8) {
	ppu.reg.oamData = data
	ppu.oam.Write(uint16(ppu.reg.oamAddr), ppu.reg.oamData)
	if !ppu.inter.IsNMI() {
		ppu.reg.oamAddr++
	}
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

type Sprite struct {
	Image *ebiten.Image
	X     uint16
	Y     uint16
}

type Tile struct {
	pixel *[8][8]byte
}

const (
	TILEROWSIZE = 30
	TILECOLSIZE = 32
)

var colors = getColorTable()

func tileToBlock(tile uint8) uint8 {
	return tile / 2
}

func tileToPixel(tile uint8) uint8 {
	return tile * 8
}

func pixelToTile(pixel uint16) uint8 {
	return uint8(pixel / 8)
}

//Sprite -> Palette -> Color
func (ppu *PPU) getColor(pixel *[8][8]byte, paletteID, x, y uint8) color.Color {
	colorID := ppu.bus.Read(uint16(0x3F00) + uint16(paletteID)*4 + uint16(pixel[y][x]))
	return colors[colorID]
}

//Tile -> Pixel
func (ppu *PPU) fillTileInBackground(tile *Tile, tileCol, tileRow, paletteID uint8) {
	for y := uint16(0); y < 8; y++ {
		for x := uint16(0); x < 8; x++ {
			px := uint16(tileToPixel(tileCol))
			py := uint16(tileToPixel(tileRow))
			ppu.background.Set(int(px+x), int(py+y), ppu.getColor(tile.pixel, paletteID, uint8(x), uint8(y)))
		}
	}
}

func (ppu *PPU) getCharacterPatternID(tileCol, tileRow uint8) uint8 {
	//0x2000, 0x2400, 0x2800, 0x3200 NameTable
	baseAddr := ppu.reg.ppuCtrl.nameTableAddr
	patternID := ppu.bus.Read(baseAddr + uint16(tileRow)*TILECOLSIZE + uint16(tileCol))
	return patternID
}

func (ppu *PPU) getPaletteID(tileCol, tileRow uint8) uint8 {
	//0x23C0, 0x27C0, 0x2BC0, 0x2FC0 Attribute Table
	baseAddr := ppu.reg.ppuCtrl.nameTableAddr + uint16(0x03C0)
	blockx := tileToBlock(tileCol)
	blocky := tileToBlock(tileRow)
	attribute := ppu.bus.Read(baseAddr + uint16(blocky)/2*8 + uint16(blockx)/2)
	offsetx := blockx % 2
	offsety := blocky % 2
	blockID := offsety*2 + offsetx
	paletteID := (attribute >> uint8(blockID*2)) & 0x03
	return paletteID
}

func (ppu *PPU) buildPatternTable(patternID uint8, patternTableAddr uint16) *[8][8]byte {
	pixel := new([8][8]byte)
	baseAddr := patternTableAddr + uint16(patternID)*0x10
	for y := uint16(0); y < 8; y++ {
		low := ppu.bus.Read(baseAddr + y)
		high := ppu.bus.Read(baseAddr + y + 8)
		for x := 0; x < 8; x++ {
			if pkg.Uint8tob(high & (0x80 >> x)) {
				pixel[y][x] += 2
			}
			if pkg.Uint8tob(low & (0x80 >> x)) {
				pixel[y][x] += 1
			}
		}
	}
	return pixel
}

func (ppu *PPU) placeTile(tileCol, tileRow uint8) {
	tile := new(Tile)
	patternID := ppu.getCharacterPatternID(tileCol, tileRow)
	paletteID := ppu.getPaletteID(tileCol, tileRow)
	tile.pixel = ppu.buildPatternTable(patternID, ppu.reg.ppuCtrl.backgroundPatternTableAddr)
	ppu.fillTileInBackground(tile, tileCol, tileRow, paletteID)
}

func (ppu *PPU) fillBackGround() {
	tileRow := uint8(ppu.scanline / 8)
	for tileCol := uint8(0); tileCol < TILECOLSIZE; tileCol++ {
		ppu.placeTile(tileCol, tileRow)
	}
}

func (ppu *PPU) fillSpriteInImage(sprite *Sprite, attribute, patternID uint8) {
	pixel := ppu.buildPatternTable(patternID, ppu.reg.ppuCtrl.spritePatternTableAddr)
	paletteID := 0x04 + (attribute & 0x03)
	//priority := pkg.Uint8tob(attribute & 0x20)
	flipHorizontal := pkg.Uint8tob(attribute & 0x40)
	flipVertical := pkg.Uint8tob(attribute & 0x80)
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			xx := x
			yy := y
			if flipHorizontal {
				xx = 7 - xx
			}
			if flipVertical {
				yy = 7 - yy
			}
			sprite.Image.Set(x, y, ppu.getColor(pixel, paletteID, uint8(xx), uint8(yy)))
		}
	}
}

func (ppu *PPU) fillSprites() {
	for i := uint16(0); i < SPRITECOUNT; i++ {
		offset := i * SPRITESIZE
		ppu.sprites[i].Y = uint16(ppu.oam.Read(offset)) + 1
		patternID := ppu.oam.Read(offset + 1)
		attribute := ppu.oam.Read(offset + 2)
		ppu.sprites[i].X = uint16(ppu.oam.Read(offset + 3))
		ppu.fillSpriteInImage(&ppu.sprites[i], attribute, patternID)
	}
}

func (ppu *PPU) isSpriteZeroHits() bool {
	return uint16(ppu.sprites[0].Y) == ppu.scanline
}

func (ppu *PPU) setSpriteZeroHits() {
	ppu.reg.ppuStatus |= (1 << 6)
}

func (ppu *PPU) unsetSpriteZeroHits() {
	ppu.reg.ppuStatus ^= (1 << 6)
}

func (ppu *PPU) setVBlank() {
	ppu.reg.ppuStatus |= (1 << 7)
	if ppu.reg.ppuCtrl.vblankNMI {
		ppu.inter.SetNMI()
	}
}

func (ppu *PPU) unsetVBlank() {
	ppu.reg.ppuStatus ^= (1 << 7)
	ppu.inter.UnSetNMI()
}

func (ppu *PPU) Run(cycles uint16) (*ebiten.Image, *[SPRITECOUNT]Sprite) {
	//ppu.info.PPUX = ppu.clock
	//ppu.info.PPUY = ppu.scanline
	ppu.clock += cycles

	if ppu.scanline == 0 {
		ppu.background.Clear()
		ppu.fillSprites()
	}
	if ppu.clock >= 341 {
		ppu.clock -= 341

		if ppu.isSpriteZeroHits() {
			ppu.setSpriteZeroHits()
		}
		//Visible scanscanlines (0-239)
		if ppu.scanline%8 == 0 && ppu.scanline <= 239 {
			ppu.fillBackGround()
		}
		ppu.scanline++
		//Post-render scanscanline (240)
		//Vertical blanking scanlines (241-260)
		if ppu.scanline == 241 {
			ppu.setVBlank()
		}
		//Pre-render scanscanline (-1 or 261)
		if ppu.scanline == 261 {
			ppu.unsetVBlank()
			ppu.unsetSpriteZeroHits()
			ppu.scanline = 0
			if !ppu.reg.ppuMask.showBackground && !ppu.reg.ppuMask.showSprite {
				return nil, nil
			}
			if !ppu.reg.ppuMask.showBackground && ppu.reg.ppuMask.showSprite {
				return nil, &ppu.sprites
			}
			if ppu.reg.ppuMask.showBackground && !ppu.reg.ppuMask.showSprite {
				return ppu.background, nil
			}
			return ppu.background, &ppu.sprites
		}
	}
	return nil, nil
}
