package akc695x

// REG16, REG17

// isMode3k returns 3k steps or 5k steps.
func (r AKC695X) isMode3k() bool {
	return r.readRegister8(0x16)&(1<<7) != 0
}

// GetStereoDemodulate returns that the stereo signal is more than 30% percent.
func (r AKC695X) GetStereoDemodulate() bool {
	return r.readRegister8(0x17)&(1<<7) != 0
}

// GetCNR returns C/N Ratio (dB).
func (r AKC695X) GetCNR() (CNR uint8) {
	if r.IsFM() {
		CNR = r.readRegister8(0x17) & 0b01111111
	} else {
		CNR = r.readRegister8(0x16) & 0b01111111
	}
	return
}
