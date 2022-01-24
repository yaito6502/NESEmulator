package cpu

import (
	"github.com/yaito6502/NESEmulator/pkg"
)

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

func (cpu *CPU) getInstructionTable() []func(uint16, bool) {
	return []func(uint16, bool){
		/*0x00*/ cpu.brk, cpu.ora, cpu.hlt, cpu.slo, cpu.dop, cpu.ora, cpu.asl, cpu.slo, cpu.php, cpu.ora, cpu.asl, nil, cpu.top, cpu.ora, cpu.asl, cpu.slo,
		/*0x10*/ cpu.bpl, cpu.ora, cpu.hlt, cpu.slo, cpu.dop, cpu.ora, cpu.asl, cpu.slo, cpu.clc, cpu.ora, cpu.nop, cpu.slo, cpu.top, cpu.ora, cpu.asl, cpu.slo,
		/*0x20*/ cpu.jsr, cpu.and, cpu.hlt, cpu.rla, cpu.bit, cpu.and, cpu.rol, cpu.rla, cpu.plp, cpu.and, cpu.rol, nil, cpu.bit, cpu.and, cpu.rol, cpu.rla,
		/*0x30*/ cpu.bmi, cpu.and, cpu.hlt, cpu.rla, cpu.dop, cpu.and, cpu.rol, cpu.rla, cpu.sec, cpu.and, cpu.nop, cpu.rla, cpu.top, cpu.and, cpu.rol, cpu.rla,
		/*0x40*/ cpu.rti, cpu.eor, cpu.hlt, cpu.sre, cpu.dop, cpu.eor, cpu.lsr, cpu.sre, cpu.pha, cpu.eor, cpu.lsr, nil, cpu.jmp, cpu.eor, cpu.lsr, cpu.sre,
		/*0x50*/ cpu.bvc, cpu.eor, cpu.hlt, cpu.sre, cpu.dop, cpu.eor, cpu.lsr, cpu.sre, cpu.cli, cpu.eor, cpu.nop, cpu.sre, cpu.top, cpu.eor, cpu.lsr, cpu.sre,
		/*0x60*/ cpu.rts, cpu.adc, cpu.hlt, cpu.rra, cpu.dop, cpu.adc, cpu.ror, cpu.rra, cpu.pla, cpu.adc, cpu.ror, nil, cpu.jmp, cpu.adc, cpu.ror, cpu.rra,
		/*0x70*/ cpu.bvs, cpu.adc, cpu.hlt, cpu.rra, cpu.dop, cpu.adc, cpu.ror, cpu.rra, cpu.sei, cpu.adc, cpu.nop, cpu.rra, cpu.top, cpu.adc, cpu.ror, cpu.rra,
		/*0x80*/ cpu.dop, cpu.sta, cpu.dop, cpu.sax, cpu.sty, cpu.sta, cpu.stx, cpu.sax, cpu.dey, cpu.dop, cpu.txa, nil, cpu.sty, cpu.sta, cpu.stx, cpu.sax,
		/*0x90*/ cpu.bcc, cpu.sta, cpu.hlt, nil, cpu.sty, cpu.sta, cpu.stx, cpu.sax, cpu.tya, cpu.sta, cpu.txs, nil, nil, cpu.sta, nil, nil,
		/*0xA0*/ cpu.ldy, cpu.lda, cpu.ldx, cpu.lax, cpu.ldy, cpu.lda, cpu.ldx, cpu.lax, cpu.tay, cpu.lda, cpu.tax, nil, cpu.ldy, cpu.lda, cpu.ldx, cpu.lax,
		/*0xB0*/ cpu.bcs, cpu.lda, cpu.hlt, cpu.lax, cpu.ldy, cpu.lda, cpu.ldx, cpu.lax, cpu.clv, cpu.lda, cpu.tsx, nil, cpu.ldy, cpu.lda, cpu.ldx, cpu.lax,
		/*0xC0*/ cpu.cpy, cpu.cmp, cpu.dop, cpu.dcp, cpu.cpy, cpu.cmp, cpu.dec, cpu.dcp, cpu.iny, cpu.cmp, cpu.dex, nil, cpu.cpy, cpu.cmp, cpu.dec, cpu.dcp,
		/*0xD0*/ cpu.bne, cpu.cmp, cpu.hlt, cpu.dcp, cpu.dop, cpu.cmp, cpu.dec, cpu.dcp, cpu.cld, cpu.cmp, cpu.nop, cpu.dcp, cpu.top, cpu.cmp, cpu.dec, cpu.dcp,
		/*0xE0*/ cpu.cpx, cpu.sbc, cpu.dop, cpu.isb, cpu.cpx, cpu.sbc, cpu.inc, cpu.isb, cpu.inx, cpu.sbc, cpu.nop, cpu.sbc, cpu.cpx, cpu.sbc, cpu.inc, cpu.isb,
		/*0xF0*/ cpu.beq, cpu.sbc, cpu.hlt, cpu.isb, cpu.dop, cpu.sbc, cpu.inc, cpu.isb, cpu.sed, cpu.sbc, cpu.nop, cpu.isb, cpu.top, cpu.sbc, cpu.inc, cpu.isb,
	}
}

func (cpu *CPU) reset() {
	cpu.sei(0, false)
	cpu.PC = uint16(cpu.bus.Read(0xFFFD))<<8 + uint16(cpu.bus.Read(0xFFFC))
}

func (cpu *CPU) nmi() {
	cpu.P.B = false
	cpu.push(uint8(cpu.PC >> 8))
	cpu.push(uint8(cpu.PC & 0x00FF))
	cpu.php(0, false)
	cpu.sei(0, false)
	cpu.PC = uint16(cpu.bus.Read(0xFFFB))<<8 + uint16(cpu.bus.Read(0xFFFA))
}

func (cpu *CPU) irq() {
	if cpu.P.I {
		return
	}
	cpu.P.B = false
	cpu.push(uint8(cpu.PC >> 8))
	cpu.push(uint8(cpu.PC & 0x00FF))
	cpu.php(0, false)
	cpu.sei(0, false)
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

//UnOfficial Opecodes

//転送命令
func (cpu *CPU) lda(opeland uint16, isaddr bool) {
	cpu.A = cpu.bus.Read(opeland)
	cpu.P.N = cpu.A&0x80 != 0
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) ldx(opeland uint16, isaddr bool) {
	cpu.X = cpu.bus.Read(opeland)
	cpu.P.N = cpu.X&0x80 != 0
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) ldy(opeland uint16, isaddr bool) {
	cpu.Y = cpu.bus.Read(opeland)
	cpu.P.N = cpu.Y&0x80 != 0
	cpu.P.Z = cpu.Y == 0
}

func (cpu *CPU) sta(opeland uint16, isaddr bool) {
	cpu.bus.Write(opeland, cpu.A)
}

func (cpu *CPU) stx(opeland uint16, isaddr bool) {
	cpu.bus.Write(opeland, cpu.X)
}

func (cpu *CPU) sty(opeland uint16, isaddr bool) {
	cpu.bus.Write(opeland, cpu.Y)
}

func (cpu *CPU) tax(opeland uint16, isaddr bool) {
	cpu.X = cpu.A
	cpu.P.N = cpu.X&0x80 != 0
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) tay(opeland uint16, isaddr bool) {
	cpu.Y = cpu.A
	cpu.P.N = cpu.Y&0x80 != 0
	cpu.P.Z = cpu.Y == 0
}

func (cpu *CPU) tsx(opeland uint16, isaddr bool) {
	cpu.X = cpu.S
	cpu.P.N = cpu.X&0x80 != 0
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) txa(opeland uint16, isaddr bool) {
	cpu.A = cpu.X
	cpu.P.N = cpu.A&0x80 != 0
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) txs(opeland uint16, isaddr bool) {
	cpu.S = cpu.X
	cpu.P.N = true
	cpu.P.Z = false
}

func (cpu *CPU) tya(opeland uint16, isaddr bool) {
	cpu.A = cpu.Y
	cpu.P.N = cpu.A&0x80 != 0
	cpu.P.Z = cpu.A == 0
}

//算術命令

//UnderStanding OverFlow Flag
func (cpu *CPU) adc(opeland uint16, isaddr bool) {
	adc := uint16(cpu.A) + uint16(cpu.bus.Read(opeland)) + pkg.Btouint16(cpu.P.C)
	cpu.P.N = adc&0x80 != 0
	cpu.P.V = (uint16(cpu.A)^adc)&(uint16(cpu.bus.Read(opeland))^adc)&0x80 != 0
	cpu.P.Z = uint8(adc) == 0
	cpu.P.C = adc > 0xFF
	cpu.A = uint8(adc)
}

func (cpu *CPU) and(opeland uint16, isaddr bool) {
	cpu.A &= cpu.bus.Read(opeland)
	cpu.P.N = cpu.A&0x80 != 0
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) asl(opeland uint16, isaddr bool) {
	if isaddr {
		asl := uint16(cpu.bus.Read(opeland)) << 1
		cpu.P.N = asl&0x80 != 0
		cpu.P.Z = uint8(asl) == 0
		cpu.P.C = asl > 0xFF
		cpu.bus.Write(opeland, uint8(asl))
	} else {
		asl := uint16(cpu.A) << 1
		cpu.P.N = asl&0x80 != 0
		cpu.P.Z = uint8(asl) == 0
		cpu.P.C = asl > 0xFF
		cpu.A = uint8(asl)
	}
}

func (cpu *CPU) bit(opeland uint16, isaddr bool) {
	data := cpu.bus.Read(opeland)
	cpu.P.N = data&0x80 != 0
	cpu.P.V = data&0x40 != 0
	cpu.P.Z = (cpu.A & data) == 0
}

func (cpu *CPU) cmp(opeland uint16, isaddr bool) {
	cmp := int16(uint16(cpu.A) - uint16(cpu.bus.Read(opeland)))
	cpu.P.N = cmp&0x80 != 0
	cpu.P.Z = cmp == 0
	cpu.P.C = cmp >= 0
}

func (cpu *CPU) cpx(opeland uint16, isaddr bool) {
	cmp := int16(uint16(cpu.X) - uint16(cpu.bus.Read(opeland)))
	cpu.P.N = cmp&0x80 != 0
	cpu.P.Z = cmp == 0
	cpu.P.C = cmp >= 0
}

func (cpu *CPU) cpy(opeland uint16, isaddr bool) {
	cmp := int16(uint16(cpu.Y) - uint16(cpu.bus.Read(opeland)))
	cpu.P.N = cmp&0x80 != 0
	cpu.P.Z = cmp == 0
	cpu.P.C = cmp >= 0
}

func (cpu *CPU) dec(opeland uint16, isaddr bool) {
	dec := cpu.bus.Read(opeland) - 1
	cpu.bus.Write(opeland, dec)
	cpu.P.N = dec&0x80 != 0
	cpu.P.Z = dec == 0
}

func (cpu *CPU) dex(opeland uint16, isaddr bool) {
	cpu.X--
	cpu.P.N = cpu.X&0x80 != 0
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) dey(opeland uint16, isaddr bool) {
	cpu.Y--
	cpu.P.N = cpu.Y&0x80 != 0
	cpu.P.Z = cpu.Y == 0
}

func (cpu *CPU) eor(opeland uint16, isaddr bool) {
	cpu.A ^= cpu.bus.Read(opeland)
	cpu.P.N = cpu.A&0x80 != 0
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) inc(opeland uint16, isaddr bool) {
	inc := cpu.bus.Read(opeland) + 1
	cpu.bus.Write(opeland, inc)
	cpu.P.N = inc&0x80 != 0
	cpu.P.Z = inc == 0
}

func (cpu *CPU) inx(opeland uint16, isaddr bool) {
	cpu.X++
	cpu.P.N = cpu.X&0x80 != 0
	cpu.P.Z = cpu.X == 0
}

func (cpu *CPU) iny(opeland uint16, isaddr bool) {
	cpu.Y++
	cpu.P.N = cpu.Y&0x80 != 0
	cpu.P.Z = cpu.Y == 0
}

func (cpu *CPU) lsr(opeland uint16, isaddr bool) {
	if isaddr {
		lsr := cpu.bus.Read(opeland) >> 1
		cpu.P.N = lsr&0x80 != 0
		cpu.P.Z = lsr == 0
		cpu.P.C = cpu.bus.Read(opeland)&0x01 != 0
		cpu.bus.Write(opeland, lsr)
	} else {
		lsr := cpu.A >> 1
		cpu.P.N = lsr&0x80 != 0
		cpu.P.Z = lsr == 0
		cpu.P.C = cpu.A&0x01 != 0
		cpu.A = lsr
	}
}

func (cpu *CPU) ora(opeland uint16, isaddr bool) {
	cpu.A |= cpu.bus.Read(opeland)
	cpu.P.N = cpu.A&0x80 != 0
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) rol(opeland uint16, isaddr bool) {
	if isaddr {
		rol := cpu.bus.Read(opeland)<<1 + pkg.Btouint8(cpu.P.C)
		cpu.P.N = rol&0x80 != 0
		cpu.P.Z = rol == 0
		cpu.P.C = cpu.bus.Read(opeland)&0x80 != 0
		cpu.bus.Write(opeland, rol)
	} else {
		rol := cpu.A<<1 + pkg.Btouint8(cpu.P.C)
		cpu.P.N = rol&0x80 != 0
		cpu.P.Z = rol == 0
		cpu.P.C = cpu.A&0x80 != 0
		cpu.A = rol
	}
}

func (cpu *CPU) ror(opeland uint16, isaddr bool) {
	if isaddr {
		ror := cpu.bus.Read(opeland)>>1 + pkg.Btouint8(cpu.P.C)<<7
		cpu.P.N = ror&0x80 != 0
		cpu.P.Z = ror == 0
		cpu.P.C = cpu.bus.Read(opeland)&0x01 != 0
		cpu.bus.Write(opeland, ror)
	} else {
		ror := cpu.A>>1 + pkg.Btouint8(cpu.P.C)<<7
		cpu.P.N = ror&0x80 != 0
		cpu.P.Z = ror == 0
		cpu.P.C = cpu.A&0x01 != 0
		cpu.A = ror
	}
}

func (cpu *CPU) sbc(opeland uint16, isaddr bool) {
	sbc := uint16(cpu.A) - uint16(cpu.bus.Read(opeland)) - (pkg.Btouint16(cpu.P.C) ^ 0x01)
	cpu.P.N = sbc&0x80 != 0
	cpu.P.V = (uint16(cpu.A)^sbc)&(uint16(cpu.A)^uint16(cpu.bus.Read(opeland)))&0x80 != 0
	cpu.P.Z = uint8(sbc) == 0
	cpu.P.C = sbc <= 0xFF
	cpu.A = uint8(sbc)
}

//スタック命令
func (cpu *CPU) pha(opeland uint16, isaddr bool) {
	cpu.push(cpu.A)
}

func (cpu *CPU) php(opeland uint16, isaddr bool) {
	data := pkg.Btouint8(cpu.P.N) << 7
	data += pkg.Btouint8(cpu.P.V) << 6
	data += 1 << 5
	data += 1 << 4
	data += pkg.Btouint8(cpu.P.D) << 3
	data += pkg.Btouint8(cpu.P.I) << 2
	data += pkg.Btouint8(cpu.P.Z) << 1
	data += pkg.Btouint8(cpu.P.C)
	cpu.push(data)
}

func (cpu *CPU) pla(opeland uint16, isaddr bool) {
	cpu.A = cpu.pop()
	cpu.P.N = cpu.A&0x80 != 0
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) plp(opeland uint16, isaddr bool) {
	plp := cpu.pop()
	cpu.P.N = plp&0x80 != 0
	cpu.P.V = plp&0x40 != 0
	//cpu.P.R = plp&0x20 != 0
	//cpu.P.B = plp&0x10 != 0
	cpu.P.D = plp&0x08 != 0
	cpu.P.I = plp&0x04 != 0
	cpu.P.Z = plp&0x02 != 0
	cpu.P.C = plp&0x01 != 0
}

//ジャンプ命令
func (cpu *CPU) jmp(opeland uint16, isaddr bool) {
	cpu.PC = opeland
}

func (cpu *CPU) jsr(opeland uint16, isaddr bool) {
	cpu.PC--
	cpu.push(uint8(cpu.PC >> 8))
	cpu.push(uint8(cpu.PC & 0x00FF))
	cpu.PC = opeland
}

func (cpu *CPU) rts(opeland uint16, isaddr bool) {
	cpu.PC = uint16(cpu.pop()) + uint16(cpu.pop())<<8
	cpu.PC++
}

func (cpu *CPU) rti(opeland uint16, isaddr bool) {
	rti := cpu.pop()
	cpu.P.N = rti&0x80 != 0
	cpu.P.V = rti&0x40 != 0
	//cpu.P.R = rti&0x20 != 0
	cpu.P.B = rti&0x10 != 0
	cpu.P.D = rti&0x08 != 0
	cpu.P.I = rti&0x04 != 0
	cpu.P.Z = rti&0x02 != 0
	cpu.P.C = rti&0x01 != 0
	cpu.PC = uint16(cpu.pop()) + uint16(cpu.pop())<<8
}

//分岐命令
func (cpu *CPU) bcc(opeland uint16, isaddr bool) {
	if !cpu.P.C {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bcs(opeland uint16, isaddr bool) {
	if cpu.P.C {
		cpu.PC = opeland
	}
}

func (cpu *CPU) beq(opeland uint16, isaddr bool) {
	if cpu.P.Z {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bmi(opeland uint16, isaddr bool) {
	if cpu.P.N {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bne(opeland uint16, isaddr bool) {
	if !cpu.P.Z {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bpl(opeland uint16, isaddr bool) {
	if !cpu.P.N {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bvc(opeland uint16, isaddr bool) {
	if !cpu.P.V {
		cpu.PC = opeland
	}
}

func (cpu *CPU) bvs(opeland uint16, isaddr bool) {
	if cpu.P.V {
		cpu.PC = opeland
	}
}

//フラグ変更命令
func (cpu *CPU) clc(opeland uint16, isaddr bool) {
	cpu.P.C = false
}

func (cpu *CPU) cld(opeland uint16, isaddr bool) {
	cpu.P.D = false
}

func (cpu *CPU) cli(opeland uint16, isaddr bool) {
	cpu.P.I = false
}

func (cpu *CPU) clv(opeland uint16, isaddr bool) {
	cpu.P.V = false
}

func (cpu *CPU) sec(opeland uint16, isaddr bool) {
	cpu.P.C = true
}

func (cpu *CPU) sed(opeland uint16, isaddr bool) {
	cpu.P.D = true
}

func (cpu *CPU) sei(opeland uint16, isaddr bool) {
	cpu.P.I = true
}

//その他の命令
func (cpu *CPU) brk(opeland uint16, isaddr bool) {
	if cpu.P.I {
		return
	}
	cpu.P.B = true
	cpu.PC++
	cpu.push(uint8(cpu.PC >> 8))
	cpu.push(uint8(cpu.PC & 0x00FF))
	cpu.php(0, false)
	cpu.sei(0, false)
	cpu.PC = uint16(cpu.bus.Read(0xFFFF))<<8 + uint16(cpu.bus.Read(0xFFFE))
}

func (cpu *CPU) nop(opeland uint16, isaddr bool) {

}

//UnOfficial Opecodes

func (cpu *CPU) dop(opeland uint16, isaddr bool) {
	cpu.nop(opeland, isaddr)
	cpu.nop(opeland, isaddr)
}

func (cpu *CPU) top(opeland uint16, isaddr bool) {
	cpu.nop(opeland, isaddr)
	cpu.nop(opeland, isaddr)
	cpu.nop(opeland, isaddr)
}

func (cpu *CPU) lax(opeland uint16, isaddr bool) {
	cpu.lda(opeland, isaddr)
	cpu.ldx(opeland, isaddr)
}

func (cpu *CPU) sax(opeland uint16, isaddr bool) {
	cpu.bus.Write(opeland, cpu.A&cpu.X)
}

func (cpu *CPU) dcp(opeland uint16, isaddr bool) {
	cpu.dec(opeland, isaddr)
	cpu.cmp(opeland, isaddr)
}

func (cpu *CPU) isb(opeland uint16, isaddr bool) {
	cpu.inc(opeland, isaddr)
	cpu.sbc(opeland, isaddr)
}

func (cpu *CPU) slo(opeland uint16, isaddr bool) {
	cpu.asl(opeland, isaddr)
	cpu.A |= cpu.bus.Read(opeland)
	cpu.P.N = cpu.A&0x80 != 0
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) rla(opeland uint16, isaddr bool) {
	cpu.rol(opeland, isaddr)
	cpu.A &= cpu.bus.Read(opeland)
	cpu.P.N = cpu.A&0x80 != 0
	cpu.P.Z = cpu.A == 0
}

func (cpu *CPU) sre(opeland uint16, isaddr bool) {
	cpu.lsr(opeland, isaddr)
	cpu.eor(opeland, isaddr)
}

func (cpu *CPU) rra(opeland uint16, isaddr bool) {
	cpu.ror(opeland, isaddr)
	cpu.adc(opeland, isaddr)
}

func (cpu *CPU) hlt(opeland uint16, isaddr bool) {
	cpu.reset()
}
