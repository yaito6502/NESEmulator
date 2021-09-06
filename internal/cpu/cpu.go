package cpu

import "github.com/yaito6502/NESEmulator/internal/mem"

type Instructions interface {
	push(uint8)
	pop(uint8) uint8
}

type Status struct {
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
	Mem mem.Memory
	A   uint8
	X   uint8
	Y   uint8
	S   uint8
	P   Status
	PC  uint16
}

func NewStatus() *Status {
	status := &Status{}
	status.N = false
	status.V = false
	status.R = true
	status.B = true
	status.D = false
	status.I = true
	status.Z = false
	status.C = false
	return status
}

func NewCPU() *CPU {
	cpu := &CPU{}
	cpu.A = 0x00
	cpu.X = 0x00
	cpu.Y = 0x00
	cpu.S = 0xFF
	cpu.P = *NewStatus()
	cpu.PC = 0x0000
	return cpu
}

func (cpu CPU) push(data uint8) {
	address := 0x0100 + uint16(cpu.S)
	cpu.Mem.Store(address, data)
	cpu.S--
}

func (cpu CPU) pop() uint8 {
	cpu.S++
	address := 0x0100 + uint16(cpu.S)
	return (cpu.Mem.Fetch(address))
}

func (cpu CPU) storePC(high byte, low byte) {
	ext_high := uint16(high)
	ext_low := uint16(low)
	cpu.PC = ext_high<<8 + ext_low
}

func (cpu CPU) fetchPC() (high byte, log byte) {
	return byte(cpu.PC & 0xFF00), byte(cpu.PC & 0x00FF)
}

func (cpu CPU) reset() {
	cpu.P.I = true
	cpu.storePC(cpu.Mem.Fetch(0xFFFD), cpu.Mem.Fetch(0xFFFC))
}

func (cpu CPU) nmi() {
	cpu.P.B = false
	high, low := cpu.fetchPC()
	cpu.push(high)
	cpu.push(low)
	//cpu.push(cpu.P)
	cpu.P.I = true
	cpu.storePC(cpu.Mem.Fetch(0xFFFB), cpu.Mem.Fetch(0xFFFA))
}

func (cpu CPU) irq() {
	if cpu.P.I {
		return
	}
	cpu.P.B = true
	high, low := cpu.fetchPC()
	cpu.push(high)
	cpu.push(low)
	//cpu.push(cpu.P)
	cpu.P.I = true
	cpu.storePC(cpu.Mem.Fetch(0xFFFF), cpu.Mem.Fetch(0xFFFE))
}

func (cpu CPU) brk() {
	if cpu.P.I {
		return
	}
	cpu.P.B = true
	cpu.PC++
	high, low := cpu.fetchPC()
	cpu.push(high)
	cpu.push(low)
	//cpu.push(cpu.P)
	cpu.P.I = true
	cpu.storePC(cpu.Mem.Fetch(0xFFFF), cpu.Mem.Fetch(0xFFFE))
}
