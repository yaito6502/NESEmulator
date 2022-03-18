package pad

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yaito6502/NESEmulator/internal/mem"
	"github.com/yaito6502/NESEmulator/pkg"
)

type PAD struct {
	mem.RAM
	latch bool
}

func NewPAD() *PAD {
	pad := new(PAD)
	pad.RAM = mem.NewRAM(0x0002)
	pad.latch = false
	return pad
}

func (pad *PAD) SetPressKeys() {
	keys := uint8(0x00)
	keys |= pkg.Btouint8(ebiten.IsKeyPressed(ebiten.KeyM)) * 0x80 //0 - A
	keys |= pkg.Btouint8(ebiten.IsKeyPressed(ebiten.KeyN)) * 0x40 //1 - B
	keys |= pkg.Btouint8(ebiten.IsKeyPressed(ebiten.KeyC)) * 0x20 //2 - Select
	keys |= pkg.Btouint8(ebiten.IsKeyPressed(ebiten.KeyV)) * 0x10 //3 - Start
	keys |= pkg.Btouint8(ebiten.IsKeyPressed(ebiten.KeyW)) * 0x08 //4 - Up
	keys |= pkg.Btouint8(ebiten.IsKeyPressed(ebiten.KeyS)) * 0x04 //5 - Down
	keys |= pkg.Btouint8(ebiten.IsKeyPressed(ebiten.KeyA)) * 0x02 //6 - Left
	keys |= pkg.Btouint8(ebiten.IsKeyPressed(ebiten.KeyD)) * 0x01 //7 - Right
	pad.RAM.Write(0x0000, keys)
	pad.RAM.Write(0x0001, keys)
}

func (pad *PAD) Read(address uint16) uint8 {
	data := pad.RAM.Read(address)
	if !pad.latch {
		pad.RAM.Write(address, data<<1)
	}
	return pkg.Btouint8(data&0x80 != 0)
}

func (pad *PAD) Write(data uint8) {
	pad.latch = pkg.Uint8tob(data & 0x01)
}
