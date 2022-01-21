package tools

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/yaito6502/NESEmulator/internal/cartridge"
	"github.com/yaito6502/NESEmulator/internal/mem"
)

func SpriteDump(nespath string) {
	pngpath := nespath[:len(nespath)-len(filepath.Ext(nespath))] + ".png"
	f, err := os.Create(pngpath)
	if err != nil {
		log.Fatal(err)
	}
	cart := cartridge.NewCartridge(nespath)
	//lenからいい感じのサイズで作る
	const width = 8 * 50
	height := 8 * (len(cart.CharacterRom) / 2 / width)
	img := image.NewGray(image.Rect(0, 0, width, height))

	for spriteID := 0; spriteID < (width/8)*(height/8); spriteID++ {
		sprite := NewSprite(&cart.CharacterRom, uint16(spriteID))
		topleftX := (spriteID % (width / 8)) * 8
		topleftY := (spriteID / (width / 8)) * 8
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				img.Set(topleftX+x, topleftY+y, color.Gray{(sprite[y][x]+1)*64 - 1})
			}
		}
	}
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dump Character from .nes -> .png ")
	fmt.Println("FilePath : ", pngpath)
}

func NewSprite(crom *mem.ROM, spriteID uint16) *[8][8]byte {
	sprite := new([8][8]byte)
	for y := uint16(0); y < 8; y++ {
		low := crom.Read(0x10*spriteID + y)
		high := crom.Read(0x10*spriteID + y + 8)
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
