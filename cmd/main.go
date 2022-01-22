package main

import (
	"log"
	"os"

	"github.com/yaito6502/NESEmulator/internal/cartridge"
	"github.com/yaito6502/NESEmulator/internal/nes"
	"github.com/yaito6502/NESEmulator/tools"
)

//N cpu.P = (cpu.P & 0b01111111) | exp
//V cpu.P = (cpu.P & 0b10111111) | exp
//R cpu.P = (cpu.P & 0b11011111) | exp
//B cpu.P = (cpu.P & 0b11101111) | exp
//D cpu.P = (cpu.P & 0b11110111) | exp
//I cpu.P = (cpu.P & 0b11111011) | exp
//Z cpu.P = (cpu.P & 0b11111101) | exp
//C cpu.P = (cpu.P & 0b11111110) | exp

func main() {
	//[TODO]flagを使用し、debugモードやoriginal sizeの設定を可能にする
	if os.Args[1] == "" {
		log.Fatal("Input File Not Found")
	}
	tools.SpriteDump(os.Args[1])
	nes := nes.NewNES(cartridge.NewCartridge(os.Args[1]))
	nes.Run()
}
