package akc695x

// REG00-REG05, REG0B

func (r AKC695X) powerOn() {
	r.reg[0] |= 1 << 7
}

func (r AKC695X) PowerOff() {
	r.reg[0] &^= 1 << 7

}

func (r AKC695X) setFM(b bool) {
	if b {
		r.reg[0] |= 1 << 6
	} else {
		r.reg[0] &^= 1 << 6
	}
}

func (r AKC695X) isFM() bool {
	return r.reg[0]&(1<<6) != 0
}

func (r AKC695X) setTune(b bool) {
	if b {
		r.reg[0] |= 1 << 5

	} else {
		r.reg[0] &^= 1 << 5
	}
}

func (r AKC695X) setSeek(b bool) {
	if b {
		r.reg[0] |= 1 << 4
	} else {
		r.reg[0] &^= 1 << 4
	}
}

func (r AKC695X) setSeekUp(b bool) {
	if b {
		r.reg[0] |= 1 << 3
	} else {
		r.reg[0] &^= 1 << 3
	}
}

func (r AKC695X) setMute(b bool) {
	if b {
		r.reg[0] |= 1 << 2
	} else {
		r.reg[0] &^= 1 << 2
	}
}

func (r AKC695X) setAMBand() {
	band := r.amband
	r.reg[0x01] &= 0b00000111
	r.reg[0x01] |= (band & 0b00011111) << 3
	if band == 0 || band == 2 {
		r.setMode3k(true)
	}
}

func (r AKC695X) setFMBand() {
	band := r.fmband
	r.reg[0x01] &= 0b11111000
	r.reg[0x01] |= (band & 0b00000111)
	if r.fmband >= 7 {
		r.setChannelRange(r.fmlowch, r.fmhighch)
	}
}

func (r AKC695X) SetRef32kMode(b bool) {
	if b {
		r.reg[0x02] |= 1 << 6
	} else {
		r.reg[0x02] &^= 1 << 6
	}
}

func (r AKC695X) setMode3k(b bool) {
	if b {
		r.reg[0x02] |= 1 << 5
	} else {
		r.reg[0x02] &^= 1 << 5
	}
}

func (r AKC695X) setChannel(ch uint16) {
	r.reg[0x02] &= 0b11100000
	r.reg[0x03] = byte(ch & 0xff)
	r.reg[0x02] |= byte((ch >> 8) & 0x1f)
}

func (r AKC695X) setChannelRange(from, to uint16) (err error) {
	r.reg[0x04] = byte(from >> 5)
	r.reg[0x05] = byte(to>>5) + 1
	return r.writeReg(4, 5)
}

func (r AKC695X) SetFMSeekStep(step uint8) {
	r.reg[0x0b] &= 0b11001111
	r.reg[0x0b] |= (step & 0x03) << 4
}
