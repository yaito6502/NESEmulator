package cpu

import "fmt"

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
	return uint16(cpu.A)
}

func (cpu *CPU) immediateAddressing() uint16 {
	address := cpu.PC
	cpu.PC++
	return address
}

func (cpu *CPU) absoluteAddressing() uint16 {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	return high<<8 + low
}

func (cpu *CPU) zeroPageAddressing() uint16 {
	low := uint16(cpu.fetch())
	high := uint16(0x0000)
	return high<<8 + low
}

func (cpu *CPU) XindexedZeroPageAddressing() uint16 {
	low := cpu.fetch() + cpu.X
	high := uint16(0x0000)
	return high<<8 + uint16(low)
}

func (cpu *CPU) YindexedZeroPageAddressing() uint16 {
	low := cpu.fetch() + cpu.Y
	high := uint16(0x0000)
	return high<<8 + uint16(low)
}

func (cpu *CPU) XindexedAbsoluteAddressing() uint16 {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	return high<<8 + low + uint16(cpu.X)
}

func (cpu *CPU) YindexedAbsoluteAddressing() uint16 {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	return high<<8 + low + uint16(cpu.Y)
}

func (cpu *CPU) impliedAddressing() uint16 {
	return 0x0000
}

func (cpu *CPU) relativeAddressing() uint16 {
	address := cpu.PC + 1
	offset := int8(cpu.fetch())
	return uint16(int32(address) + int32(offset))
}

func (cpu *CPU) indexedIndirectAddressing() uint16 {
	low := cpu.fetch() + cpu.X
	high := uint16(0x0000)
	low1 := uint16(cpu.mem.Fetch(uint16(high<<8) + uint16(low) + 1))
	high1 := uint16(cpu.mem.Fetch(uint16(high<<8) + uint16(low)))
	return high1<<8 + low1
}

func (cpu *CPU) indirectIndexedAddressing() uint16 {
	low := cpu.fetch()
	high := uint16(0x0000)
	low1 := uint16(cpu.mem.Fetch(uint16(high<<8) + uint16(low) + 1))
	high1 := uint16(cpu.mem.Fetch(uint16(high<<8) + uint16(low)))
	return high1<<8 + low1 + uint16(cpu.Y)
}

func (cpu *CPU) absoluteIndirectAddressing() uint16 {
	low := uint16(cpu.fetch())
	high := uint16(cpu.fetch())
	low = high<<8 + low
	high = high<<8 + low + 1
	return high<<8 + low
}

/*
func (cpu *CPU) reset() {
	cpu.P.I = true
	cpu.storePC(cpu.mem.Fetch(0xFFFD), cpu.mem.Fetch(0xFFFC))
}

func (cpu *CPU) nmi() {
	cpu.P.B = false
	high, low := cpu.fetchPC()
	cpu.push(high)
	cpu.push(low)
	//cpu.push(cpu.P)
	cpu.P.I = true
	cpu.storePC(cpu.mem.Fetch(0xFFFB), cpu.mem.Fetch(0xFFFA))
}

func (cpu *CPU) irq() {
	if cpu.P.I {
		return
	}
	cpu.P.B = true
	high, low := cpu.fetchPC()
	cpu.push(high)
	cpu.push(low)
	//cpu.push(cpu.P)
	cpu.P.I = true
	cpu.storePC(cpu.mem.Fetch(0xFFFF), cpu.mem.Fetch(0xFFFE))
}
*/

//hello worldに最低限必要な命令を優先して実装する

//転送命令
func (cpu *CPU) lda(opeland uint16) {
	fmt.Print(cpu.PC, " lda")
	cpu.A = cpu.mem.Fetch(opeland)
	cpu.P.N = (cpu.A>>7)&1 == 1
	cpu.P.Z = (cpu.A == 0)
}

func (cpu *CPU) ldx(opeland uint16) {
	fmt.Print(cpu.PC, " ldx")
	cpu.X = cpu.mem.Fetch(opeland)
	cpu.P.N = (cpu.X>>7)&1 == 1
	cpu.P.Z = (cpu.X == 0)
}

func (cpu *CPU) ldy(opeland uint16) {
	fmt.Print(cpu.PC, " ldy")
	cpu.Y = cpu.mem.Fetch(opeland)
	cpu.P.N = (cpu.Y>>7)&1 == 1
	cpu.P.Z = (cpu.Y == 0)
}

func (cpu *CPU) sta(opeland uint16) {
	fmt.Print(cpu.PC, " sta")
	cpu.mem.Store(opeland, cpu.A)
}

func (cpu *CPU) stx(opeland uint16) {

}

func (cpu *CPU) sty(opeland uint16) {

}

func (cpu *CPU) tax(opeland uint16) {

}

func (cpu *CPU) tay(opeland uint16) {

}

func (cpu *CPU) tsx(opeland uint16) {

}

func (cpu *CPU) txa(opeland uint16) {

}

func (cpu *CPU) txs(opeland uint16) {
	fmt.Print(cpu.PC, " txs")
	cpu.S = cpu.X
}

func (cpu *CPU) tya(opeland uint16) {

}

//算術命令
func (cpu *CPU) adc(opeland uint16) {

}

func (cpu *CPU) and(opeland uint16) {

}

func (cpu *CPU) asl(opeland uint16) {

}

func (cpu *CPU) bit(opeland uint16) {

}

func (cpu *CPU) cmp(opeland uint16) {

}

func (cpu *CPU) cpx(opeland uint16) {

}

func (cpu *CPU) cpy(opeland uint16) {

}

func (cpu *CPU) dec(opeland uint16) {

}

func (cpu *CPU) dex(opeland uint16) {

}

func (cpu *CPU) dey(opeland uint16) {
	fmt.Print(cpu.PC, " dey")
	cpu.Y--
	cpu.P.N = (cpu.Y>>7)&1 == 1
	cpu.P.Z = (cpu.Y == 0)
}

func (cpu *CPU) eor(opeland uint16) {

}

func (cpu *CPU) inc(opeland uint16) {

}

func (cpu *CPU) inx(opeland uint16) {
	fmt.Print(cpu.PC, " inx")
	cpu.X++
	cpu.P.N = (cpu.X>>7)&1 == 1
	cpu.P.Z = (cpu.X == 0)
}

func (cpu *CPU) iny(opeland uint16) {

}

func (cpu *CPU) lsr(opeland uint16) {

}

func (cpu *CPU) ora(opeland uint16) {

}

func (cpu *CPU) rol(opeland uint16) {

}

func (cpu *CPU) ror(opeland uint16) {

}

func (cpu *CPU) sbc(opeland uint16) {

}

//スタック命令
func (cpu *CPU) pha(opeland uint16) {

}

func (cpu *CPU) php(opeland uint16) {

}

func (cpu *CPU) pla(opeland uint16) {

}

func (cpu *CPU) plp(opeland uint16) {

}

//ジャンプ命令
func (cpu *CPU) jmp(opeland uint16) {
	fmt.Print(cpu.PC, " jmp")
	cpu.PC = opeland
}

func (cpu *CPU) jsr(opeland uint16) {

}

func (cpu *CPU) rts(opeland uint16) {

}

func (cpu *CPU) rti(opeland uint16) {

}

//分岐命令
func (cpu *CPU) bcc(opeland uint16) {

}

func (cpu *CPU) bcs(opeland uint16) {

}

func (cpu *CPU) beq(opeland uint16) {

}

func (cpu *CPU) bmi(opeland uint16) {

}

func (cpu *CPU) bne(opeland uint16) {
	fmt.Print(cpu.PC, " bne")
	if !cpu.P.Z {
		cpu.PC = opeland
	} else {
		cpu.PC++
	}
}

func (cpu *CPU) bpl(opeland uint16) {

}

func (cpu *CPU) bvc(opeland uint16) {

}

func (cpu *CPU) bvs(opeland uint16) {

}

//フラグ変更命令
func (cpu *CPU) clc(opeland uint16) {

}

func (cpu *CPU) cld(opeland uint16) {

}

func (cpu *CPU) cli(opeland uint16) {

}

func (cpu *CPU) clv(opeland uint16) {

}

func (cpu *CPU) sec(opeland uint16) {

}
func (cpu *CPU) sed(opeland uint16) {

}

func (cpu *CPU) sei(opeland uint16) {
	fmt.Print(cpu.PC, " sei")
	cpu.P.I = true
}

//その他の命令
func (cpu *CPU) brk(opeland uint16) {
	fmt.Print(cpu.PC, " brk")
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
	cpu.storePC(cpu.mem.Fetch(0xFFFF), cpu.mem.Fetch(0xFFFE))
}

func (cpu *CPU) nop(opeland uint16) {
}
