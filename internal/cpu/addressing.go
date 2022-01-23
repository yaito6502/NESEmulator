package cpu

import (
	"github.com/yaito6502/NESEmulator/pkg"
)

func (cpu *CPU) getAddressingModeTable() []func() uint16 {
	return []func() uint16{
		/*0x00*/ cpu.impliedAddressing, cpu.indexedIndirectAddressing, nil, nil, nil, cpu.zeroPageAddressing, cpu.zeroPageAddressing, nil, cpu.impliedAddressing, cpu.immediateAddressing, cpu.accumulatorAddressing, nil, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, nil,
		/*0x10*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, nil, nil, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, nil, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, nil, nil, nil, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, nil,
		/*0x20*/ cpu.absoluteAddressing, cpu.indexedIndirectAddressing, nil, nil, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, nil, cpu.impliedAddressing, cpu.immediateAddressing, cpu.accumulatorAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, nil,
		/*0x30*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, nil, nil, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, nil, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, nil, nil, nil, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, nil,
		/*0x40*/ cpu.impliedAddressing, cpu.indexedIndirectAddressing, nil, nil, nil, cpu.zeroPageAddressing, cpu.zeroPageAddressing, nil, cpu.impliedAddressing, cpu.immediateAddressing, cpu.accumulatorAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, nil,
		/*0x50*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, nil, nil, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, nil, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, nil, nil, nil, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, nil,
		/*0x60*/ cpu.impliedAddressing, cpu.indexedIndirectAddressing, nil, nil, nil, cpu.zeroPageAddressing, cpu.zeroPageAddressing, nil, cpu.impliedAddressing, cpu.immediateAddressing, cpu.accumulatorAddressing, nil, cpu.absoluteIndirectAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, nil,
		/*0x70*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, nil, nil, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, nil, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, nil, nil, nil, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, nil,
		/*0x80*/ nil, cpu.XindexedAbsoluteAddressing, nil, nil, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, nil, cpu.impliedAddressing, nil, cpu.impliedAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, nil,
		/*0x90*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, nil, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.YindexedZeroPageAddressing, nil, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, nil, nil, cpu.XindexedAbsoluteAddressing, nil, nil,
		/*0xA0*/ cpu.immediateAddressing, cpu.indexedIndirectAddressing, cpu.immediateAddressing, nil, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, nil, cpu.impliedAddressing, cpu.immediateAddressing, cpu.impliedAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, nil,
		/*0xB0*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, nil, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, cpu.YindexedZeroPageAddressing, nil, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, cpu.impliedAddressing, nil, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, cpu.YindexedAbsoluteAddressing, nil,
		/*0xC0*/ cpu.immediateAddressing, cpu.indexedIndirectAddressing, nil, nil, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, nil, cpu.impliedAddressing, cpu.immediateAddressing, cpu.impliedAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, nil,
		/*0xD0*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, nil, nil, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, nil, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, nil, nil, nil, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, nil,
		/*0xE0*/ cpu.immediateAddressing, cpu.indexedIndirectAddressing, nil, nil, cpu.zeroPageAddressing, cpu.zeroPageAddressing, cpu.zeroPageAddressing, nil, cpu.impliedAddressing, cpu.immediateAddressing, cpu.impliedAddressing, nil, cpu.absoluteAddressing, cpu.absoluteAddressing, cpu.absoluteAddressing, nil,
		/*0xF0*/ cpu.relativeAddressing, cpu.indirectIndexedAddressing, nil, nil, nil, cpu.XindexedZeroPageAddressing, cpu.XindexedZeroPageAddressing, nil, cpu.impliedAddressing, cpu.YindexedAbsoluteAddressing, nil, nil, nil, cpu.XindexedAbsoluteAddressing, cpu.XindexedAbsoluteAddressing, nil,
	}
}

func (cpu *CPU) accumulatorAddressing() uint16 {
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(cpu.bus.Read(uint16(cpu.A))))
	cpu.info.ASMCODE += " #$" + pkg.ConvUpperHexString(uint64(cpu.bus.Read(uint16(cpu.A))))
	return uint16(cpu.A)
}

func (cpu *CPU) immediateAddressing() uint16 {
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(cpu.bus.Read(cpu.PC)))
	cpu.info.ASMCODE += " #$" + pkg.ConvUpperHexString(uint64(cpu.bus.Read(cpu.PC)))
	address := cpu.PC
	cpu.PC++
	return address
}

func (cpu *CPU) absoluteAddressing() uint16 {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	address := high<<8 + low
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address))
	return address
}

func (cpu *CPU) zeroPageAddressing() uint16 {
	low := uint16(cpu.fetch())
	address := 0x0000 + low
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address)) + " = " + pkg.ConvUpperHexString(uint64(cpu.bus.Read(address)))
	return address
}

func (cpu *CPU) XindexedZeroPageAddressing() uint16 {
	low := cpu.fetch() + cpu.X
	address := 0x0000 + uint16(low)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address)) + ",X @ " + "?" + " = " + "?"
	return address
}

func (cpu *CPU) YindexedZeroPageAddressing() uint16 {
	low := cpu.fetch() + cpu.Y
	address := 0x0000 + uint16(low)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address)) + ",Y @ " + "?" + " = " + "?"
	return address
}

func (cpu *CPU) XindexedAbsoluteAddressing() uint16 {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	address := high<<8 + low + uint16(cpu.X)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address)) + ",X @ " + "?" + " = " + "?"
	return address
}

func (cpu *CPU) YindexedAbsoluteAddressing() uint16 {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	address := high<<8 + low + uint16(cpu.Y)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(address)) + ",Y @ " + "?" + " = " + "?"
	return address
}

func (cpu *CPU) impliedAddressing() uint16 {
	return 0x0000
}

func (cpu *CPU) relativeAddressing() uint16 {
	address := cpu.PC + 1
	offset := int8(cpu.fetch())
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(offset))
	cpu.info.ASMCODE += " $" + pkg.ConvUpperHexString(uint64(int32(address)+int32(offset)))
	return uint16(int32(address) + int32(offset))
}

func (cpu *CPU) indexedIndirectAddressing() uint16 {
	low := cpu.fetch() + cpu.X
	high := uint16(0x0000)
	low1 := uint16(cpu.bus.Read(uint16(high<<8) + uint16(low) + 1))
	high1 := uint16(cpu.bus.Read(uint16(high<<8) + uint16(low)))
	address := high1<<8 + low1
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low1))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high1))
	cpu.info.ASMCODE += " ($" + pkg.ConvUpperHexString(uint64(address)) + ",X) @ " + "?" + " = " + "?" + " = " + "?"
	return address
}

func (cpu *CPU) indirectIndexedAddressing() uint16 {
	low := cpu.fetch()
	high := uint16(0x0000)
	low1 := uint16(cpu.bus.Read(uint16(high<<8) + uint16(low) + 1))
	high1 := uint16(cpu.bus.Read(uint16(high<<8) + uint16(low)))
	address := high1<<8 + low1 + uint16(cpu.Y)
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low1))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high1))
	cpu.info.ASMCODE += " ($" + pkg.ConvUpperHexString(uint64(address)) + "),Y = " + "?" + " @ " + "?" + " = " + "?"
	//($33),Y = 0400 @ 0400 = 3F
	return address
}

func (cpu *CPU) absoluteIndirectAddressing() uint16 {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	low = high<<8 + low
	high = high<<8 + low + 1
	address := high<<8 + low
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(low))
	cpu.info.MACHINECODE += " " + pkg.ConvUpperHexString(uint64(high))
	cpu.info.ASMCODE += " ($" + pkg.ConvUpperHexString(uint64(address)) + ") = " + pkg.ConvUpperHexString(uint64(cpu.bus.Read(address)))
	return address
}
