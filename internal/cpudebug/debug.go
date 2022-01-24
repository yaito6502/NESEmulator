package cpudebug

import (
	"fmt"
)

type DebugInfo struct {
	PC          uint16
	MACHINECODE string
	ASMCODE     string
	A           uint8
	X           uint8
	Y           uint8
	P           uint8
	SP          uint8
	PPUY        uint16
	PPUX        uint16
	CYCLE       uint64
}

func (info *DebugInfo) Print() {
	//Log Example
	//PC   MACHINECODE ASMCODE [28spaces]A X Y P SP PPUY PPUX CLOCK
	//C000  4C F5 C5  JMP $C5F5                       A:00 X:00 Y:00 P:24 SP:FD PPU:  0, 21 CYC:7
	fmt.Printf("%X  %-8s  %-32sA:%2.2X X:%2.2X Y:%2.2X P:%2.2X SP:%2.2X PPU:%3d,%3d CYC:%d\n", info.PC, info.MACHINECODE, info.ASMCODE, info.A, info.X, info.Y, info.P, info.SP, info.PPUY, info.PPUX, info.CYCLE)
	info.MACHINECODE = ""
	info.ASMCODE = ""
}
