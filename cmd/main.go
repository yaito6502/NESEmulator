package main

import (
	"github.com/yaito6502/NESEmulator/internal/cartridge"
	"github.com/yaito6502/NESEmulator/internal/nes"
)

func main() {
	//[TODO]flagを使用し、debugモードやoriginal sizeの設定を可能にする
	nes := nes.NewNES(cartridge.NewCartridge("../third_party/sample1.nes"))
	nes.Run()
}
