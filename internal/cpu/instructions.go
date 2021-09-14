package cpu

func getCyclesTable() []int {
	return []int{
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

/*
func getInstructionTable() []func(byte) {
	return []func(byte){
		NewCPU().adc, NewCPU().sbc,
	}
}
*/

func getAddressingModeTable() func() {
	return func() {
	}
}

func (cpu *CPU) accumulatorAddressing() byte {
	return cpu.A
}

func (cpu *CPU) immediateAddressing() byte {
	return cpu.fetch()
}

func (cpu *CPU) absoluteAddressing() byte {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	return cpu.Mem.Fetch(high<<8 + low)
}

func (cpu *CPU) zeroPageAddressing() byte {
	low := uint16(cpu.fetch())
	high := uint16(0x0000)
	return cpu.Mem.Fetch(high<<8 + low)
}

func (cpu *CPU) XindexedZeroPageAddressing() byte {
	low := cpu.fetch() + cpu.X
	high := uint16(0x0000)
	return cpu.Mem.Fetch(high<<8 + uint16(low))
}

func (cpu *CPU) YindexedZeroPageAddressing() byte {
	low := cpu.fetch() + cpu.Y
	high := uint16(0x0000)
	return cpu.Mem.Fetch(high<<8 + uint16(low))
}

func (cpu *CPU) XindexedAbsoluteAddressing() byte {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	return cpu.Mem.Fetch(uint16(high<<8 + low + uint16(cpu.X)))
}

func (cpu *CPU) YindexedAbsoluteAddressing() byte {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	return cpu.Mem.Fetch(uint16(high<<8 + low + uint16(cpu.Y)))
}

func (cpu *CPU) impliedAddressing() byte {
	return 0x00
}

func (cpu *CPU) relativeAddressing() byte {
	address := int16(cpu.PC)
	offset := int16(cpu.fetch())
	return cpu.Mem.Fetch(uint16(address + offset))
}

func (cpu *CPU) indexedIndirectAddressing() byte {
	low := cpu.fetch() + cpu.X
	high := uint16(0x0000)
	low1 := uint16(cpu.Mem.Fetch(uint16(high<<8) + uint16(low) + 1))
	high1 := uint16(cpu.Mem.Fetch(uint16(high<<8) + uint16(low)))
	return cpu.Mem.Fetch(high1<<8 + low1)
}

func (cpu *CPU) indirectIndexedAddressing() byte {
	low := cpu.fetch()
	high := uint16(0x0000)
	low1 := uint16(cpu.Mem.Fetch(uint16(high<<8) + uint16(low) + 1))
	high1 := uint16(cpu.Mem.Fetch(uint16(high<<8) + uint16(low)))
	return cpu.Mem.Fetch(high1<<8 + low1 + uint16(cpu.Y))
}

func (cpu *CPU) absoluteIndirectAddressing() byte {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	low = high<<8 + low
	high = high<<8 + low + 1
	return cpu.Mem.Fetch(high<<8 + low)
}
