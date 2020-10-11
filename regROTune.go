package akc695x

// REG14-REG15, REG1A

// IsStereo returns whether current operation mode is Stereo or not.
func (r AKC695X) IsStereo() bool {
	return r.readRegister8(0x14)&(1<<7) != 0
}

// IsSeekComplete returns whether Seek or Tune is complete or not.
func (r AKC695X) IsSeekComplete() bool {
	return r.readRegister8(0x14)&(1<<6) != 0
}

// IsTuned returns whether receiving signal or not.
func (r AKC695X) IsTuned() bool {
	return r.readRegister8(0x14)&(1<<5) != 0
}

func (r AKC695X) getChannel() uint16 {
	return r.readRegister16(0x14) & 0x1fff
}

// GetFrequencyDeviation returns Status of frequncy deviation.
// The number 1 is scaled to 1KHz for FM and scaled to 100Hz for AM.
func (r AKC695X) GetFrequencyDeviation() int8 {
	return int8(r.readRegister8(0x1a))
}
