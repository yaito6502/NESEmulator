package interrupts

type Interrupts struct {
	nmi bool
	irq bool
}

func NewInterrupts() *Interrupts {
	inter := new(Interrupts)
	inter.nmi = false
	inter.irq = false
	return inter
}

func (inter *Interrupts) SetNMI() {
	inter.nmi = true
}

func (inter *Interrupts) UnSetNMI() {
	inter.nmi = false
}

func (inter *Interrupts) IsNMI() bool {
	return inter.nmi
}