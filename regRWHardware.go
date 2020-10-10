package akc695x

// REG09, REG0C

func (r AKC695X) setVolumeControlI2C() {
	r.reg[0x09] |= 1 << 3
}

func (r AKC695X) SetVolumeControlResistor() {
	r.reg[0x09] &^= 1 << 3
}

func (r AKC695X) SetOscillatorSourceCrystal(b bool) {
	if b {
		r.reg[0x09] |= 1 << 2
	} else {
		r.reg[0x09] &^= 1 << 2
	}
}

func (r AKC695X) SetChannelADC(b bool) {
	if b {
		r.reg[0x0c] |= 0b10000000

	} else {
		r.reg[0x0c] &= 0b01111111
	}
}

func (r AKC695X) SetRXEnable(b bool) {
	if b {
		r.reg[0x0c] |= 0b00100000
	} else {
		r.reg[0x0c] &= 0b11011111
	}
}
