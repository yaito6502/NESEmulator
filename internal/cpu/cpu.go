package cpu

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/yaito6502/NESEmulator/internal/cpubus"
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
	A      uint8
	X      uint8
	Y      uint8
	S      uint8
	P      uint8
	PC     uint16
}

func NewCPU(bus *cpubus.CPUBUS) *CPU {
	cpu := new(CPU)
	cpu.iTable = cpu.getInstructionTable()
	cpu.aTable = cpu.getAddressingModeTable()
	cpu.cycles = getCyclesTable()
	cpu.bus = bus
	cpu.A = 0x00
	cpu.X = 0x00
	cpu.Y = 0x00
	cpu.S = 0xFF
	cpu.P = 0b00110100
	cpu.PC = 0x8000
	return cpu
}

func (cpu *CPU) push(data uint8) {
	address := 0x0100 + uint16(cpu.S)
	cpu.bus.Write(address, data)
	cpu.S--
}

/*
func (cpu *CPU) pop() uint8 {
	cpu.S++
	address := 0x0100 + uint16(cpu.S)
	return (cpu.bus.Read(address))
}
*/

func (cpu *CPU) fetch() byte {
	data := cpu.bus.Read(cpu.PC)
	cpu.PC++
	return data
}

func (cpu *CPU) storePC(high byte, low byte) {
	ext_high := uint16(high)
	ext_low := uint16(low)
	cpu.PC = ext_high<<8 + ext_low
}

func (cpu *CPU) fetchPC() (high byte, low byte) {
	return byte(cpu.PC & 0xFF00), byte(cpu.PC & 0x00FF)
}

func GetFuncName(i interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func (cpu *CPU) Run() uint8 {
	opecode := cpu.fetch()
	inst := cpu.iTable[opecode]
	addressing := cpu.aTable[opecode]
	if inst == nil || addressing == nil {
		log.Fatalf("opecode[%#x] not implement\n", opecode)
	}
	inst(addressing())
	istr := GetFuncName(inst)
	astr := GetFuncName(addressing)
	fmt.Printf("inst[%s] mode[%s]\n", istr[:strings.Index(istr, "-")], astr[:strings.Index(astr, "Addressing")])
	return cpu.cycles[opecode]
}
