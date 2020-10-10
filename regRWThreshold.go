package akc695x

// REG08

// SetFMCNRThreshold . if th=0 -2dB, th=1 -3dB, th=2 -4dB, th=3 -5dB.
func (r AKC695X) SetFMCNRThreshold(th uint8) {
	th &= 0b00000011
	r.reg[0x08] &= 0b00111111
	r.reg[0x08] |= th << 6
}

// SetAMCNRThreshold . if th=0 -6dB, th=1 -10dB, th=2 -14dB, th=3 -18dB.
func (r AKC695X) SetAMCNRThreshold(th uint8) {
	th &= 0b00000011
	r.reg[0x08] &= 0b11001111
	r.reg[0x08] |= th << 4
}

// SetFreqencyDiffThreshold . see datasheet.
func (r AKC695X) SetFreqencyDiffThreshold(th uint8) error {
	th &= 0b00000011
	r.reg[0x08] &= 0b11110011
	r.reg[0x08] |= th << 2
	return nil
}

// SetFMStereoCNRThreshold . if th=0 -4dB, th=1 -8dB, th=2 -12dB, th=3 -16dB.
func (r AKC695X) SetFMStereoCNRThreshold(th uint8) {
	th &= 0b00000011
	r.reg[0x08] &= 0b11111100
	r.reg[0x08] |= th
}
