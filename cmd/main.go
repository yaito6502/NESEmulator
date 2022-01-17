package main

import (
	"github.com/yaito6502/NESEmulator/internal/nes"
)

func main() {
	nes := nes.NewNES()
	nes.AttachCartridge("../sample1.nes")
	nes.Run()
}
