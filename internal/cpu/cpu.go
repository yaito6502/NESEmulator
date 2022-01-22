package cpu

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/yaito6502/NESEmulator/internal/cpubus"
	"github.com/yaito6502/NESEmulator/internal/ppu"
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
	cycle  uint
	bus    *cpubus.CPUBUS
	ppu    *ppu.PPU
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

func NewCPU(bus *cpubus.CPUBUS, ppu *ppu.PPU) *CPU {
	cpu := new(CPU)
	cpu.iTable = cpu.getInstructionTable()
	cpu.aTable = cpu.getAddressingModeTable()
	cpu.cycles = getCyclesTable()
	cpu.cycle = 7
	cpu.ppu = ppu
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

func GetFuncName(i interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func (cpu *CPU) debug(inst func(uint16), addressing func() uint16) uint16 {
	istr := GetFuncName(inst)
	istr = strings.ToUpper(istr[:strings.Index(istr, "-")])

	data := uint8(0x00)
	if cpu.P.N {
		data |= 0x80
	}
	if cpu.P.V {
		data |= 0x40
	}
	if cpu.P.R {
		data |= 0x20
	}
	if cpu.P.B {
		data |= 0x10
	}
	if cpu.P.D {
		data |= 0x08
	}
	if cpu.P.I {
		data |= 0x04
	}
	if cpu.P.Z {
		data |= 0x02
	}
	if cpu.P.C {
		data |= 0x01
	}
	fmt.Printf("%X  %2X %2X %2X  %s                    A:%.2X X:%.2X Y:%.2X P:%.2X SP:%X PPU %3d,%3d CYC:%d\n", cpu.PC-1, 0, 0, 0, istr, cpu.A, cpu.X, cpu.Y, data, cpu.S, cpu.ppu.GetLine(), cpu.ppu.GetClock(), cpu.cycle)
	return addressing()
}

func (cpu *CPU) Run() uint8 {
	//割込割り込み(nmi, irq, brk)
	opecode := cpu.fetch()
	inst := cpu.iTable[opecode]
	addressing := cpu.aTable[opecode]
	if inst == nil || addressing == nil {
		log.Fatalf("opecode[%#x] not implement\n", opecode)
	}
	inst(cpu.debug(inst, addressing))
	cpu.cycle += uint(cpu.cycles[opecode])
	return cpu.cycles[opecode]
}
