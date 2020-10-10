package akc695x

// REG18, REG1B

func (r AKC695X) GetPGALevel() (RF, IF uint8) {
	reg18 := r.readRegister8(0x18)
	return reg18 >> 5, (reg18 >> 2) & 0b00000111
}

func (r AKC695X) GetPGARSSI() uint8 {
	return r.readRegister8(0x1b) & 0b01111111
}
