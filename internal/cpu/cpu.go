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
	P      Flags
	PC     uint16
}

func NewFlags() *Flags {
	flags := new(Flags)
	flags.N = false
	flags.V = false
	flags.R = true
	flags.B = true
	flags.D = false
	flags.I = true
	flags.Z = false
	flags.C = false
	return flags
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
	cpu.P = *NewFlags()
	cpu.PC = 0x8000
	return cpu
}

func getCyclesTable() []uint8 {
	return []uint8{
		/*0x00*/ 7, 6, 2, 8, 3, 3, 5, 5, 3, 2, 2, 2, 4, 4, 6, 6,
		/*0x10*/ 2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 6, 7,
		/*0x20*/ 6, 6, 2, 8, 3, 3, 5, 5, 4, 2, 2, 2, 4, 4, 6, 6,
		/*0x30*/ 2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 6, 7,
		/*0x40*/ 6, 6, 2, 8, 3, 3, 5, 5, 3, 2, 2, 2, 3, 4, 6, 6,
		/*0x50*/ 2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 6, 7,
		/*0x60*/ 6, 6, 2, 8, 3, 3, 5, 5, 4, 2, 2, 2, 5, 4, 6, 6,
		/*0x70*/ 2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 6, 7,
		/*0x80*/ 2, 6, 2, 6, 3, 3, 3, 3, 2, 2, 2, 2, 4, 4, 4, 4,
		/*0x90*/ 2, 6, 2, 6, 4, 4, 4, 4, 2, 4, 2, 5, 5, 4, 5, 5,
		/*0xA0*/ 2, 6, 2, 6, 3, 3, 3, 3, 2, 2, 2, 2, 4, 4, 4, 4,
		/*0xB0*/ 2, 5, 2, 5, 4, 4, 4, 4, 2, 4, 2, 4, 4, 4, 4, 4,
		/*0xC0*/ 2, 6, 2, 8, 3, 3, 5, 5, 2, 2, 2, 2, 4, 4, 6, 6,
		/*0xD0*/ 2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 7, 7,
		/*0xE0*/ 2, 6, 3, 8, 3, 3, 5, 5, 2, 2, 2, 2, 4, 4, 6, 6,
		/*0xF0*/ 2, 5, 2, 8, 4, 4, 6, 6, 2, 4, 2, 7, 4, 4, 7, 7,
	}
}

func (cpu *CPU) getInstructionTable() []func(uint16) {
	return []func(uint16){
		/*0x00*/ cpu.brk, cpu.ora, nil, nil, nil, cpu.ora, cpu.asl, nil, cpu.php, cpu.ora, cpu.asl, nil, nil, cpu.ora, cpu.asl, nil,
		/*0x10*/ cpu.bpl, cpu.ora, nil, nil, nil, cpu.ora, cpu.asl, nil, cpu.clc, cpu.ora, nil, nil, nil, cpu.ora, cpu.asl, nil,
		/*0x20*/ cpu.jsr, cpu.and, nil, nil, cpu.bit, cpu.and, cpu.rol, nil, cpu.plp, cpu.and, cpu.rol, nil, cpu.bit, cpu.and, cpu.rol, nil,
		/*0x30*/ cpu.bmi, cpu.and, nil, nil, nil, cpu.and, cpu.rol, nil, cpu.sec, cpu.and, nil, nil, nil, cpu.and, cpu.rol, nil,
		/*0x40*/ cpu.rti, cpu.eor, nil, nil, nil, cpu.eor, cpu.lsr, nil, cpu.pha, cpu.eor, cpu.lsr, nil, cpu.jmp, cpu.eor, cpu.lsr, nil,
		/*0x50*/ cpu.bvc, cpu.eor, nil, nil, nil, cpu.eor, cpu.lsr, nil, cpu.cli, cpu.eor, nil, nil, nil, cpu.eor, cpu.lsr, nil,
		/*0x60*/ cpu.rts, cpu.adc, nil, nil, nil, cpu.adc, cpu.ror, nil, cpu.pla, cpu.adc, cpu.ror, nil, cpu.jmp, cpu.adc, cpu.ror, nil,
		/*0x70*/ cpu.bvs, cpu.adc, nil, nil, nil, cpu.adc, cpu.ror, nil, cpu.sei, cpu.adc, nil, nil, nil, cpu.adc, cpu.ror, nil,
		/*0x80*/ nil, cpu.sta, nil, nil, cpu.sty, cpu.sta, cpu.stx, nil, cpu.dey, nil, cpu.txa, nil, cpu.sty, cpu.sta, cpu.stx, nil,
		/*0x90*/ cpu.bcc, cpu.sta, nil, nil, cpu.sty, cpu.sta, cpu.stx, nil, cpu.tya, cpu.sta, cpu.txs, nil, nil, cpu.sta, nil, nil,
		/*0xA0*/ cpu.ldy, cpu.lda, cpu.ldx, nil, cpu.ldy, cpu.lda, cpu.ldx, nil, cpu.tay, cpu.lda, cpu.tax, nil, cpu.ldy, cpu.lda, cpu.ldx, nil,
		/*0xB0*/ cpu.bcs, cpu.lda, nil, nil, cpu.ldy, cpu.lda, cpu.ldx, nil, cpu.clv, cpu.lda, cpu.tsx, nil, cpu.ldy, cpu.lda, cpu.ldx, nil,
		/*0xC0*/ cpu.cpy, cpu.cmp, nil, nil, cpu.cpy, cpu.cmp, cpu.dec, nil, cpu.iny, cpu.cmp, cpu.dex, nil, cpu.cpy, cpu.cmp, cpu.dec, nil,
		/*0xD0*/ cpu.bne, cpu.cmp, nil, nil, nil, cpu.cmp, cpu.dec, nil, cpu.cld, cpu.cmp, nil, nil, nil, cpu.cmp, cpu.dec, nil,
		/*0xE0*/ cpu.cpx, cpu.sbc, nil, nil, cpu.cpx, cpu.sbc, cpu.inc, nil, cpu.inx, cpu.sbc, cpu.nop, nil, cpu.cpx, cpu.sbc, cpu.inc, nil,
		/*0xF0*/ cpu.beq, cpu.sbc, nil, nil, nil, cpu.sbc, cpu.inc, nil, cpu.sed, cpu.sbc, nil, nil, nil, cpu.sbc, cpu.inc, nil,
	}
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
