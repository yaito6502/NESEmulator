package main

import (
	"log"
	"os"

	"github.com/yaito6502/NESEmulator/internal/cartridge"
	"github.com/yaito6502/NESEmulator/internal/nes"
	"github.com/yaito6502/NESEmulator/tools"
)

func main() {
	//[TODO]flagを使用し、debugモードやoriginal sizeの設定を可能にする
	if os.Args[1] == "" {
		log.Fatal("Input File Not Found")
	}
	tools.SpriteDump(os.Args[1])
	nes := nes.NewNES(cartridge.NewCartridge(os.Args[1]))
	nes.Run()
}
