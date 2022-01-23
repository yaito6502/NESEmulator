package cpu

import (
	"log"
	"strings"

	"github.com/yaito6502/NESEmulator/internal/cpubus"
	"github.com/yaito6502/NESEmulator/internal/cpudebug"
	"github.com/yaito6502/NESEmulator/pkg"
)

type Flags struct {
	N bool
	V bool
	R bool
	B bool
	D bool
	I bool
	Z bool
	C bool
}

type CPU struct {
	iTable []func(uint16)
	aTable []func() uint16
	cycles []uint8
	bus    *cpubus.CPUBUS
	info   *cpudebug.DebugInfo
	A      uint8
	X      uint8
	Y      uint8
	S      uint8
	P      Flags
	PC     uint16
}

func NewFlags() *Flags {
	flags := new(Flags)
	flags.N = false
	flags.V = false
	flags.R = true
	flags.B = false
	flags.D = false
	flags.I = true
	flags.Z = false
	flags.C = false
	return flags
}

func NewCPU(bus *cpubus.CPUBUS, info *cpudebug.DebugInfo) *CPU {
	cpu := new(CPU)
	cpu.iTable = cpu.getInstructionTable()
	cpu.aTable = cpu.getAddressingModeTable()
	cpu.cycles = getCyclesTable()
	cpu.info = info
	cpu.bus = bus
	cpu.A = 0x00
	cpu.X = 0x00
	cpu.Y = 0x00
	cpu.S = 0xFD
	cpu.P = *NewFlags()
	cpu.PC = 0xC000
	return cpu
}

func (cpu *CPU) push(data uint8) {
	cpu.bus.Write(0x0100+uint16(cpu.S), data)
	cpu.S--
}

func (cpu *CPU) pop() uint8 {
	cpu.S++
	return cpu.bus.Read(0x0100 + uint16(cpu.S))
}

func (cpu *CPU) fetch() byte {
	data := cpu.bus.Read(cpu.PC)
	cpu.PC++
	return data
}

func (cpu *CPU) Run() uint8 {
	//割込割り込み(nmi, irq, brk)
	cpu.info.PC = cpu.PC

	opecode := cpu.fetch()
	inst := cpu.iTable[opecode]
	addressing := cpu.aTable[opecode]
	if inst == nil || addressing == nil {
		log.Fatalf("opecode[%#x] not implement\n", opecode)
	}

	cpu.info.MACHINECODE += pkg.ConvUpperHexString(uint64(opecode))
	cpu.info.ASMCODE += strings.ToUpper(strings.Split(pkg.GetFuncName(inst), "-")[0])
	cpu.info.A = cpu.A
	cpu.info.X = cpu.X
	cpu.info.Y = cpu.Y
	data := pkg.Btouint8(cpu.P.N) << 7
	data += pkg.Btouint8(cpu.P.V) << 6
	data += pkg.Btouint8(cpu.P.R) << 5
	data += pkg.Btouint8(cpu.P.B) << 4
	data += pkg.Btouint8(cpu.P.D) << 3
	data += pkg.Btouint8(cpu.P.I) << 2
	data += pkg.Btouint8(cpu.P.Z) << 1
	data += pkg.Btouint8(cpu.P.C)
	cpu.info.P = data
	cpu.info.SP = cpu.S

	inst(addressing())
	return cpu.cycles[opecode]
}
