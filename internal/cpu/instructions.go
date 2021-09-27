package cpu

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

/*
func (cpu *CPU) reset() {
	cpu.P.I = true
	cpu.storePC(cpu.Mem.Fetch(0xFFFD), cpu.Mem.Fetch(0xFFFC))
}

func (cpu *CPU) nmi() {
	cpu.P.B = false
	high, low := cpu.fetchPC()
	cpu.push(high)
	cpu.push(low)
	//cpu.push(cpu.P)
	cpu.P.I = true
	cpu.storePC(cpu.Mem.Fetch(0xFFFB), cpu.Mem.Fetch(0xFFFA))
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
	cpu.storePC(cpu.Mem.Fetch(0xFFFF), cpu.Mem.Fetch(0xFFFE))
}
*/

//hello worldに最低限必要な命令を優先して実装する

//転送命令
func (cpu *CPU) lda(data byte) {
}

func (cpu *CPU) ldx(data byte) {
	cpu.P.N = (data>>7)&1 == 1
	cpu.P.Z = (data == 0)
	cpu.X = data
}

func (cpu *CPU) ldy(data byte) {

}

func (cpu *CPU) sta(data byte) {

}

func (cpu *CPU) stx(data byte) {

}

func (cpu *CPU) sty(data byte) {

}

func (cpu *CPU) tax(data byte) {

}

func (cpu *CPU) tay(data byte) {

}

func (cpu *CPU) tsx(data byte) {

}

func (cpu *CPU) txa(data byte) {

}

func (cpu *CPU) txs(data byte) {

}

func (cpu *CPU) tya(data byte) {

}

//算術命令
func (cpu *CPU) adc(data byte) {

}

func (cpu *CPU) and(data byte) {

}

func (cpu *CPU) asl(data byte) {

}

func (cpu *CPU) bit(data byte) {

}

func (cpu *CPU) cmp(data byte) {

}

func (cpu *CPU) cpx(data byte) {

}

func (cpu *CPU) cpy(data byte) {

}

func (cpu *CPU) dec(data byte) {

}

func (cpu *CPU) dex(data byte) {

}

func (cpu *CPU) dey(data byte) {

}

func (cpu *CPU) eor(data byte) {

}

func (cpu *CPU) inc(data byte) {

}

func (cpu *CPU) inx(data byte) {

}

func (cpu *CPU) iny(data byte) {

}

func (cpu *CPU) lsr(data byte) {

}

func (cpu *CPU) ora(data byte) {

}

func (cpu *CPU) rol(data byte) {

}

func (cpu *CPU) ror(data byte) {

}

func (cpu *CPU) sbc(data byte) {

}

//スタック命令
func (cpu *CPU) pha(data byte) {

}

func (cpu *CPU) php(data byte) {

}

func (cpu *CPU) pla(data byte) {

}

func (cpu *CPU) plp(data byte) {

}

//ジャンプ命令
func (cpu *CPU) jmp(data byte) {

}

func (cpu *CPU) jsr(data byte) {

}

func (cpu *CPU) rts(data byte) {

}

func (cpu *CPU) rti(data byte) {

}

//分岐命令
func (cpu *CPU) bcc(data byte) {

}

func (cpu *CPU) bcs(data byte) {

}

func (cpu *CPU) beq(data byte) {

}

func (cpu *CPU) bmi(data byte) {

}

func (cpu *CPU) bne(data byte) {

}

func (cpu *CPU) bpl(data byte) {

}

func (cpu *CPU) bvc(data byte) {

}

func (cpu *CPU) bvs(data byte) {

}

//フラグ変更命令
func (cpu *CPU) clc(data byte) {

}

func (cpu *CPU) cld(data byte) {

}

func (cpu *CPU) cli(data byte) {

}

func (cpu *CPU) clv(data byte) {

}

func (cpu *CPU) sec(data byte) {

}
func (cpu *CPU) sed(data byte) {

}

func (cpu *CPU) sei(data byte) {
	cpu.P.I = true
}

//その他の命令

func (cpu *CPU) brk() {
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

func (cpu *CPU) nop(data byte) {

}
