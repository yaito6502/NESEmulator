package main

import (
	"github.com/yaito6502/NESEmulator/internal/cartridge"
	"github.com/yaito6502/NESEmulator/internal/cpu"
	"github.com/yaito6502/NESEmulator/internal/mem"
)

func main() {
	mem := mem.NewMemory()
	prom, _ := cartridge.AttachCartridge("../test/sample1/sample1.nes")
	for i, b := range prom {
		mem.Store(uint16(0x8000+i), b)
	}
	cpu := cpu.NewCPU(mem)
	for {
		cpu.Run()
	}
}
