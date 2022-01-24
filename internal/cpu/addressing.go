package cpu

import (
	"github.com/yaito6502/NESEmulator/pkg"
)

func (cpu *CPU) getAddressingModeTable() []func() (uint16, bool) {
	return []func() (uint16, bool){
		/*0x00*/ cpu.impliedAddressing, cpu.indexedIndirectAddressing, nil, cpu.indexedIndirectAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.impliedAddressing, cpu.immediateAddressing, cpu.accumulatorAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing,
		/*0x10*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, cpu.indirectIndexedAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing,
		/*0x20*/ cpu.absoluteAddressing, cpu.indexedIndirectAddressing, nil, cpu.indexedIndirectAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.impliedAddressing, cpu.immediateAddressing, cpu.accumulatorAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing,
		/*0x30*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, cpu.indirectIndexedAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing,
		/*0x40*/ cpu.impliedAddressing, cpu.indexedIndirectAddressing, nil, cpu.indexedIndirectAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.impliedAddressing, cpu.immediateAddressing, cpu.accumulatorAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing,
		/*0x50*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, cpu.indirectIndexedAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing,
		/*0x60*/ cpu.impliedAddressing, cpu.indexedIndirectAddressing, nil, cpu.indexedIndirectAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.impliedAddressing, cpu.immediateAddressing, cpu.accumulatorAddressing, nil, cpu.absoluteIndirectAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing,
		/*0x70*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, cpu.indirectIndexedAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing,
		/*0x80*/ cpu.immediateAddressing, cpu.indexedIndirectAddressing, cpu.immediateAddressing, cpu.indexedIndirectAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.impliedAddressing, cpu.immediateAddressing, cpu.impliedAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing,
		/*0x90*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, nil, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.YindexedZeroPageAddressing, cpu.YindexedZeroPageAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, nil, nil, cpu.XindexedAbsoluteAddressing, nil, nil,
		/*0xA0*/ cpu.immediateAddressing, cpu.indexedIndirectAddressing, cpu.immediateAddressing, cpu.indexedIndirectAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.impliedAddressing, cpu.immediateAddressing, cpu.impliedAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing,
		/*0xB0*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, cpu.indirectIndexedAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.YindexedZeroPageAddressing, cpu.YindexedZeroPageAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, nil, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.YindexedAbsoluteAddressing, cpu.YindexedAbsoluteAddressing,
		/*0xC0*/ cpu.immediateAddressing, cpu.indexedIndirectAddressing, cpu.immediateAddressing, cpu.indexedIndirectAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.impliedAddressing, cpu.immediateAddressing, cpu.impliedAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing,
		/*0xD0*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, cpu.indirectIndexedAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing,
		/*0xE0*/ cpu.immediateAddressing, cpu.indexedIndirectAddressing, cpu.immediateAddressing, cpu.indexedIndirectAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.impliedAddressing, cpu.immediateAddressing, cpu.impliedAddressing, cpu.immediateAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing,
		/*0xF0*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, cpu.indirectIndexedAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing,
	}
}

func (cpu *CPU) accumulatorAddressing() (uint16, bool) {
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(cpu.bus.Read(uint16(cpu.A))))
	cpu.info.ASMCODE += " #$" + pkg.ConvUpperHexString(uint64(cpu.bus.Read(uint16(cpu.A))))
	return uint16(cpu.A), false
}

func (cpu *CPU) immediateAddressing() (uint16, bool) {
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(cpu.bus.Read(cpu.PC)))
	cpu.info.ASMCODE += " #$" + pkg.ConvUpperHexString(uint64(cpu.bus.Read(cpu.PC)))
	address := cpu.PC
	cpu.PC++
	return address, true
}

func (cpu *CPU) absoluteAddressing() (uint16, bool) {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	address := high<<8 + low
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address))
	return address, true
}

func (cpu *CPU) zeroPageAddressing() (uint16, bool) {
	low := uint16(cpu.fetch())
	address := 0x0000 + low
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address)) + " = " + pkg.ConvUpperHexString(uint64(cpu.bus.Read(address)))
	return address, true
}

func (cpu *CPU) XindexedZeroPageAddressing() (uint16, bool) {
	low := cpu.fetch() + cpu.X
	address := 0x0000 + uint16(low)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address)) + ",X @ " + "?" + " = " + "?"
	return address, true
}

func (cpu *CPU) YindexedZeroPageAddressing() (uint16, bool) {
	low := cpu.fetch() + cpu.Y
	address := 0x0000 + uint16(low)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address)) + ",Y @ " + "?" + " = " + "?"
	return address, true
}

func (cpu *CPU) XindexedAbsoluteAddressing() (uint16, bool) {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	address := high<<8 + low + uint16(cpu.X)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(high)) + pkg.ConvUpperHexString(uint64(low)) + ",X @ " + pkg.ConvUpperHexString(uint64(address)) + " = " + "?"
	return address, true
}

func (cpu *CPU) YindexedAbsoluteAddressing() (uint16, bool) {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	address := high<<8 + low + uint16(cpu.Y)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address)) + ",Y @ " + "?" + " = " + "?"
	return address, true
}

func (cpu *CPU) impliedAddressing() (uint16, bool) {
	return 0x0000, false
}

func (cpu *CPU) relativeAddressing() (uint16, bool) {
	address := cpu.PC + 1
	offset := int8(cpu.fetch())
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(offset))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(int32(address)+int32(offset)))
	return uint16(int32(address) + int32(offset)), true
}

func (cpu *CPU) indexedIndirectAddressing() (uint16, bool) {
	indirect := cpu.fetch() + cpu.X
	low := uint16(cpu.bus.Read(uint16(indirect)))
	high := uint16(cpu.bus.Read(uint16(indirect + 1)))
	address := high<<8 + low
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(indirect-cpu.X))
	cpu.info.ASMCODE += " ($" + pkg.ConvUpperHexString(uint64(indirect-cpu.X)) + ",X) @ " + pkg.ConvUpperHexString(uint64(indirect)) + " = " + pkg.ConvUpperHexString(uint64(high)) + pkg.ConvUpperHexString(uint64(low)) + " = " + pkg.ConvUpperHexString(uint64(cpu.bus.Read(address)))
	return address, true
}

func (cpu *CPU) indirectIndexedAddressing() (uint16, bool) {
	low := cpu.fetch()
	low1 := uint16(cpu.bus.Read(0x0000 + uint16(low)))
	high1 := uint16(cpu.bus.Read(0x0000 + uint16(low+1)))
	address := high1<<8 + low1 + uint16(cpu.Y)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.ASMCODE += " ($" + pkg.ConvUpperHexString(uint64(low)) + "),Y = " + pkg.ConvUpperHexString(uint64(high1)) + pkg.ConvUpperHexString(uint64(low1)) + " @ " + pkg.ConvUpperHexString(uint64(high1)) + pkg.ConvUpperHexString(uint64(low1+uint16(cpu.Y))) + " = " + pkg.ConvUpperHexString(uint64(cpu.bus.Read(address)))
	return address, true
}

func (cpu *CPU) absoluteIndirectAddressing() (uint16, bool) {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	address := uint16(cpu.bus.Read(high<<8 + low))
	if low == 0xFF {
		address += 0x0300
	} else {
		address += uint16(cpu.bus.Read(high<<8+low+1)) << 8
	}
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high))
	cpu.info.ASMCODE += " ($" + pkg.ConvUpperHexString(uint64(high)) + pkg.ConvUpperHexString(uint64(low)) + ") = " + pkg.ConvUpperHexString(uint64(address))
	return address, true
}
