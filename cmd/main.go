package main

import (
	"github.com/yaito6502/NESEmulator/internal/nes"
)

func main() {
	//[TODO]flagを使用し、debugモードやoriginal sizeの設定を可能にする
	nes := nes.NewNES()
	nes.Run()
}
