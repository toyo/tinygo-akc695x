package akc695x

// REG06-REG07

func (r AKC695X) setVolume(vol uint8) {
	r.reg[0x06] &= 0b00000011
	r.reg[0x06] |= vol << 2
}

// SetVolume set volume. vol from 0 to 63.
func (r AKC695X) SetVolume(vol uint8) error {
	r.setVolume(vol)
	return r.writeReg(6, 6)
}

func (r AKC695X) SetLine(b bool) {
	if b {
		r.reg[0x06] |= 1 << 1
	} else {
		r.reg[0x06] &^= 1 << 1
	}
}

func (r AKC695X) SetPhaseInv(b bool) {
	if b {
		r.reg[0x06] |= 1 << 0
	} else {
		r.reg[0x06] &^= 1 << 0
	}
}

func (r AKC695X) GetVolume() (vol uint8) {
	vol = r.reg[0x06] >> 2
	return
}

func (r AKC695X) VolumeUp() (err error) {
	vol := r.GetVolume()
	if vol < 63 {
		err = r.SetVolume(vol + 1)
	}
	return
}

func (r AKC695X) VolumeDown() (err error) {
	vol := r.GetVolume()
	if vol > 0 {
		err = r.SetVolume(vol - 1)
	}
	return
}

// SetDeEmphasis ... true means 50us. false means 75us.
func (r AKC695X) SetDeEmphasis(b bool) {
	if b {
		r.reg[0x07] |= 1 << 5
	} else {
		r.reg[0x07] &^= 1 << 5
	}
}

func (r AKC695X) SetBaseBoost(b bool) {
	if b {
		r.reg[0x07] |= 1 << 4
	} else {
		r.reg[0x07] &^= 1 << 4
	}
}

func (r AKC695X) SetAutoStereo() {
	r.reg[0x07] &= 0b11110011
}

func (r AKC695X) SetStereo() {
	r.reg[0x07] &= 0b11110011
	r.reg[0x07] |= 0b00001000
}

func (r AKC695X) SetMono() {
	r.reg[0x07] |= 0b00001100
}

// SetBandWidth . bw=0 means 150kHz, bw=1 means 200kHz, bw=2 means 50kHz, bw=3 means 100kHz.
func (r AKC695X) SetBandWidth(bw uint8) {
	r.reg[0x07] &= 0b11111100
	r.reg[0x01] |= bw & 0b00000011
}
