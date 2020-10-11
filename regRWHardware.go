package akc695x

// REG09, REG0C

func (r AKC695X) setVolumeControlI2C(b bool) {
	if b {
		r.reg[0x09] |= 1 << 3 // I2C
	} else {
		r.reg[0x09] &^= 1 << 3 // External Resistor
	}
	return
}

// SetOscillatorSourceCrystal set Oscillator source selection.
// false External XO
// true: Crystal
func (r AKC695X) SetOscillatorSourceCrystal(b bool) {
	if b {
		r.reg[0x09] |= 1 << 2
	} else {
		r.reg[0x09] &^= 1 << 2
	}
}

// SetChannelADC set Setting of channel ADC
// true; ADC enable
// false: ADC disable
func (r AKC695X) SetChannelADC(b bool) {
	if b {
		r.reg[0x0c] &= 0b01111111
	} else {
		r.reg[0x0c] |= 0b10000000
	}
}

// SetRXEnable set Setting of RX
// true: analog & RF enable
// false: anlog & RF disable
func (r AKC695X) SetRXEnable(b bool) {
	if b {
		r.reg[0x0c] &= 0b11011111
	} else {
		r.reg[0x0c] |= 0b00100000
	}
}
