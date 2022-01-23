package cpu

import "github.com/yaito6502/NESEmulator/pkg"

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

func (cpu *CPU) reset() {
	cpu.sei(0)
	cpu.PC = uint16(cpu.bus.Read(0xFFFD))<<8 + uint16(cpu.bus.Read(0xFFFC))
}

func (cpu *CPU) nmi() {
	cpu.P.B = false
	cpu.push(uint8(cpu.PC >> 8))
	cpu.push(uint8(cpu.PC & 0x00FF))
	cpu.php(0)
	cpu.sei(0)
	cpu.PC = uint16(cpu.bus.Read(0xFFFB))<<8 + uint16(cpu.bus.Read(0xFFFA))
}

func (cpu *CPU) irq() {
	if cpu.P.I {
		return
	}
	cpu.P.B = false
	cpu.push(uint8(cpu.PC >> 8))
	cpu.push(uint8(cpu.PC & 0x00FF))
	cpu.php(0)
	cpu.sei(0)
	cpu.PC = uint16(cpu.bus.Read(0xFFFF))<<8 + uint16(cpu.bus.Read(0xFFFE))
}

//hello worldに最低限必要な命令を優先して実装する

const (
	C = iota
	Z
	I
	D
	B
	R
	V
	N
)

/* flagをビット演算で処理する場合の実装
func (cpu *CPU) setFlag(bit, exp uint8) {
	cpu.P = (cpu.P & (0xFF ^ (1 << bit))) | exp
}

func (cpu *CPU) getFlag(bit uint8) uint8 {
	return cpu.P & (1 << bit)
}
*/

//転送命令
func (cpu *CPU) lda(opeland uint16) {
	cpu.A = cpu.bus.Read(opeland)
	cpu.P.N = cpu.A&0x80 == 1
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) ldx(opeland uint16) {
	cpu.X = cpu.bus.Read(opeland)
	cpu.P.N = cpu.X&0x80 == 1
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) ldy(opeland uint16) {
	cpu.Y = cpu.bus.Read(opeland)
	cpu.P.N = cpu.Y&0x80 == 1
	cpu.P.Z = cpu.Y == 0
}

func (cpu *CPU) sta(opeland uint16) {
	cpu.bus.Write(opeland, cpu.A)
}

func (cpu *CPU) stx(opeland uint16) {
	cpu.bus.Write(opeland, cpu.X)
}

func (cpu *CPU) sty(opeland uint16) {
	cpu.bus.Write(opeland, cpu.Y)
}

func (cpu *CPU) tax(opeland uint16) {
	cpu.X = cpu.A
	cpu.P.N = cpu.X&0x80 == 1
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) tay(opeland uint16) {
	cpu.Y = cpu.A
	cpu.P.N = cpu.Y&0x80 == 1
	cpu.P.Z = cpu.Y == 0
}

func (cpu *CPU) tsx(opeland uint16) {
	cpu.X = cpu.S
	cpu.P.N = cpu.X&0x80 == 1
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) txa(opeland uint16) {
	cpu.A = cpu.X
	cpu.P.N = cpu.A&0x80 == 1
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) txs(opeland uint16) {
	cpu.S = cpu.X
	cpu.P.N = cpu.S&0x80 == 1
	cpu.P.Z = cpu.S == 0
}

func (cpu *CPU) tya(opeland uint16) {
	cpu.A = cpu.Y
	cpu.P.N = cpu.A&0x80 == 1
	cpu.P.Z = cpu.A == 0
}

//算術命令
func (cpu *CPU) adc(opeland uint16) {
	adc := uint16(cpu.A) + opeland + pkg.Btouint16(cpu.P.C)
	cpu.P.N = adc&0x80 == 1
	cpu.P.V = cpu.A < 0x80 && adc >= 0x80
	cpu.P.Z = adc == 0
	cpu.P.C = adc > 0xFF
	cpu.A = uint8(adc)
}

func (cpu *CPU) and(opeland uint16) {
	cpu.A &= uint8(opeland)
	cpu.P.N = cpu.A&0x80 == 1
	cpu.P.Z = cpu.A == 0
}

//opecodeによってはメモリを左シフトする必要があるかも
func (cpu *CPU) asl(opeland uint16) {
	asl := uint16(cpu.A) << 1
	cpu.P.N = asl&0x80 == 1
	cpu.P.Z = asl == 0
	cpu.P.C = asl > 0xFF
	cpu.A = uint8(asl)
}

func (cpu *CPU) bit(opeland uint16) {
	cpu.P.N = opeland&0x80 == 1
	cpu.P.V = opeland&0x40 == 1
	cpu.P.Z = uint16(cpu.A)&opeland == 0
}

func (cpu *CPU) cmp(opeland uint16) {
	cmp := int16(uint16(cpu.A) - opeland)
	cpu.P.N = cmp&0x80 == 1
	cpu.P.Z = cmp == 0
	cpu.P.C = cmp >= 0
}

func (cpu *CPU) cpx(opeland uint16) {
	cmp := int16(uint16(cpu.X) - opeland)
	cpu.P.N = cmp&0x80 == 1
	cpu.P.Z = cmp == 0
	cpu.P.C = cmp >= 0
}

func (cpu *CPU) cpy(opeland uint16) {
	cmp := int16(uint16(cpu.Y) - opeland)
	cpu.P.N = cmp&0x80 == 1
	cpu.P.Z = cmp == 0
	cpu.P.C = cmp >= 0
}

func (cpu *CPU) dec(opeland uint16) {
	opeland--
	cpu.P.N = opeland&0x80 == 1
	cpu.P.Z = opeland == 0
}

func (cpu *CPU) dex(opeland uint16) {
	cpu.X--
	cpu.P.N = cpu.X&0x80 == 1
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) dey(opeland uint16) {
	cpu.Y--
	cpu.P.N = cpu.Y&0x80 == 1
	cpu.P.Z = cpu.Y == 0
}

func (cpu *CPU) eor(opeland uint16) {
	cpu.A ^= uint8(opeland)
	cpu.P.N = cpu.A&0x80 == 1
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) inc(opeland uint16) {
	opeland++
	cpu.P.N = opeland&0x80 == 1
	cpu.P.Z = opeland == 0
}

func (cpu *CPU) inx(opeland uint16) {
	cpu.X++
	cpu.P.N = cpu.X&0x80 == 1
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) iny(opeland uint16) {
	cpu.Y++
	cpu.P.N = cpu.Y&0x80 == 1
	cpu.P.Z = cpu.Y == 0
}

func (cpu *CPU) lsr(opeland uint16) {
	lsr := uint16(cpu.A) >> 1
	cpu.P.N = lsr&0x80 == 1
	cpu.P.Z = lsr == 0
	cpu.P.C = lsr > 0xFF
	cpu.A = uint8(lsr)
}

func (cpu *CPU) ora(opeland uint16) {
	cpu.A |= uint8(opeland)
	cpu.P.N = cpu.A&0x80 == 1
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) rol(opeland uint16) {
	cpu.A = cpu.A<<1 + pkg.Btouint8(cpu.P.C)
	cpu.P.N = cpu.A&0x80 == 1
	cpu.P.Z = cpu.A == 0
	cpu.P.C = cpu.A&0x80 != 0
}

func (cpu *CPU) ror(opeland uint16) {
	cpu.A >>= 1
	cpu.A = cpu.A>>1 + pkg.Btouint8(cpu.P.C)*0x80
	cpu.P.N = cpu.A&0x80 == 1
	cpu.P.Z = cpu.A == 0
	cpu.P.C = cpu.A&0x01 == 1
}

func (cpu *CPU) sbc(opeland uint16) {
	sbc := int16(uint16(cpu.A) - opeland - pkg.Btouint16(cpu.P.C))
	cpu.P.N = sbc&0x80 == 1
	cpu.P.V = cpu.A >= 0x80 && sbc < 0x80
	cpu.P.Z = sbc == 0
	cpu.P.C = sbc >= 0
}

//スタック命令
func (cpu *CPU) pha(opeland uint16) {
	cpu.push(cpu.A)
}

func (cpu *CPU) php(opeland uint16) {
	data := pkg.Btouint8(cpu.P.N) << 7
	data += pkg.Btouint8(cpu.P.V) << 6
	data += pkg.Btouint8(cpu.P.V) << 5
	data += pkg.Btouint8(cpu.P.V) << 4
	data += pkg.Btouint8(cpu.P.V) << 3
	data += pkg.Btouint8(cpu.P.V) << 2
	data += pkg.Btouint8(cpu.P.V) << 1
	data += pkg.Btouint8(cpu.P.V)
	cpu.push(data)
}

func (cpu *CPU) pla(opeland uint16) {
	cpu.A = cpu.pop()
	cpu.P.N = cpu.A&0x80 == 1
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) plp(opeland uint16) {
	plp := cpu.pop()
	cpu.P.N = plp&0x80 == 1
	cpu.P.V = plp&0x40 == 1
	cpu.P.R = plp&0x20 == 1
	cpu.P.B = plp&0x10 == 1
	cpu.P.D = plp&0x08 == 1
	cpu.P.I = plp&0x04 == 1
	cpu.P.Z = plp&0x02 == 1
	cpu.P.C = plp&0x01 == 1
}

//ジャンプ命令
func (cpu *CPU) jmp(opeland uint16) {
	cpu.PC = opeland
}

func (cpu *CPU) jsr(opeland uint16) {
	cpu.PC--
	cpu.push(uint8(cpu.PC >> 8))
	cpu.push(uint8(cpu.PC & 0x00FF))
	cpu.PC = opeland
}

func (cpu *CPU) rts(opeland uint16) {
	cpu.PC = uint16(cpu.pop()) + uint16(cpu.pop())<<8
	cpu.PC++
}

func (cpu *CPU) rti(opeland uint16) {
	rti := cpu.pop()
	cpu.P.N = rti&0x80 == 1
	cpu.P.V = rti&0x40 == 1
	cpu.P.R = rti&0x20 == 1
	cpu.P.B = rti&0x10 == 1
	cpu.P.D = rti&0x08 == 1
	cpu.P.I = rti&0x04 == 1
	cpu.P.Z = rti&0x02 == 1
	cpu.P.C = rti&0x01 == 1
	cpu.PC = uint16(cpu.pop()) + uint16(cpu.pop())<<8
}

//分岐命令
func (cpu *CPU) bcc(opeland uint16) {
	if !cpu.P.C {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bcs(opeland uint16) {
	if cpu.P.C {
		cpu.PC = opeland
	}
}

func (cpu *CPU) beq(opeland uint16) {
	if cpu.P.Z {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bmi(opeland uint16) {
	if cpu.P.N {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bne(opeland uint16) {
	if !cpu.P.Z {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bpl(opeland uint16) {
	if !cpu.P.N {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bvc(opeland uint16) {
	if !cpu.P.V {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bvs(opeland uint16) {
	if cpu.P.V {
		cpu.PC = opeland
	}
}

//フラグ変更命令
func (cpu *CPU) clc(opeland uint16) {
	cpu.P.C = false
}

func (cpu *CPU) cld(opeland uint16) {
	cpu.P.D = false
}

func (cpu *CPU) cli(opeland uint16) {
	cpu.P.I = false
}

func (cpu *CPU) clv(opeland uint16) {
	cpu.P.V = false
}

func (cpu *CPU) sec(opeland uint16) {
	cpu.P.C = true
}

func (cpu *CPU) sed(opeland uint16) {
	cpu.P.D = true
}

func (cpu *CPU) sei(opeland uint16) {
	cpu.P.I = true
}

//その他の命令
func (cpu *CPU) brk(opeland uint16) {
	if cpu.P.I {
		return
	}
	cpu.P.B = true
	cpu.PC++
	cpu.push(uint8(cpu.PC >> 8))
	cpu.push(uint8(cpu.PC & 0x00FF))
	cpu.php(0)
	cpu.sei(0)
	cpu.PC = uint16(cpu.bus.Read(0xFFFF))<<8 + uint16(cpu.bus.Read(0xFFFE))
}

func (cpu *CPU) nop(opeland uint16) {
}
