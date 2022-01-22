package cpu

/*
func (cpu *CPU) reset() {
	cpu.P.I = true
	cpu.storePC(cpu.bus.Read(0xFFFD), cpu.bus.Read(0xFFFC))
}

func (cpu *CPU) nmi() {
	cpu.P.B = false
	high, low := cpu.fetchPC()
	cpu.push(high)
	cpu.push(low)
	//cpu.push(cpu.P)
	cpu.P.I = true
	cpu.storePC(cpu.bus.Read(0xFFFB), cpu.bus.Read(0xFFFA))
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
	cpu.storePC(cpu.bus.Read(0xFFFF), cpu.bus.Read(0xFFFE))
}
*/

//hello worldに最低限必要な命令を優先して実装する

//転送命令
func (cpu *CPU) lda(opeland uint16) {
	cpu.A = cpu.bus.Read(opeland)
	cpu.P.N = (cpu.A>>7)&1 == 1
	cpu.P.Z = (cpu.A == 0)
}

func (cpu *CPU) ldx(opeland uint16) {
	cpu.X = cpu.bus.Read(opeland)
	cpu.P.N = (cpu.X>>7)&1 == 1
	cpu.P.Z = (cpu.X == 0)
}

func (cpu *CPU) ldy(opeland uint16) {
	cpu.Y = cpu.bus.Read(opeland)
	cpu.P.N = (cpu.Y>>7)&1 == 1
	cpu.P.Z = (cpu.Y == 0)
}

func (cpu *CPU) sta(opeland uint16) {

	cpu.bus.Write(opeland, cpu.A)
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
	//fmt.Print(cpu.PC, " dey")
	cpu.Y--
	cpu.P.N = (cpu.Y>>7)&1 == 1
	cpu.P.Z = (cpu.Y == 0)
}

func (cpu *CPU) eor(opeland uint16) {

}

func (cpu *CPU) inc(opeland uint16) {

}

func (cpu *CPU) inx(opeland uint16) {
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
	if !cpu.P.Z {
		cpu.PC = opeland
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
	cpu.P.I = true
}

//その他の命令
func (cpu *CPU) brk(opeland uint16) {
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
	cpu.storePC(cpu.bus.Read(0xFFFF), cpu.bus.Read(0xFFFE))
}

func (cpu *CPU) nop(opeland uint16) {
}
