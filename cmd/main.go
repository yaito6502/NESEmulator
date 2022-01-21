package main

import (
	"github.com/yaito6502/NESEmulator/internal/cartridge"
	"github.com/yaito6502/NESEmulator/internal/nes"
	"github.com/yaito6502/NESEmulator/tools"
)

func main() {
	//[TODO]flagを使用し、debugモードやoriginal sizeの設定を可能にする
	tools.SpriteDump("../third_party/sample1.nes")
	nes := nes.NewNES(cartridge.NewCartridge("../third_party/sample1.nes"))
	nes.Run()
}
