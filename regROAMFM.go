package akc695x

// REG16, REG17

func (r AKC695X) isMode3k() bool {
	return r.readRegister8(0x16)&(1<<7) != 0
}

func (r AKC695X) GetCNRAM() uint8 {
	return r.readRegister8(0x16) & 0b01111111
}

func (r AKC695X) GetStereoDemodulate() bool {
	return r.readRegister8(0x17)&(1<<7) != 0
}

func (r AKC695X) GetCNRFM() uint8 {
	return r.readRegister8(0x17) & 0b01111111
}
