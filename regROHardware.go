package akc695x

// REG19

func (r AKC695X) GetVCCMilliVolt() uint16 {
	return uint16(r.readRegister8(0x19)&0x3f)*50 + 1800
}
