package akc695x

// REG18, REG1B

// GetRSSIdBuV returns Pin(dBuV)
func (r AKC695X) GetRSSIdBuV() (dBu int8) {
	rssi := r.readRegister8(0x1b) & 0b01111111 // Max: 127
	reg18 := r.readRegister8(0x18)             //
	RF := (reg18 >> 5) * 6                     // Max: 7*6=42
	IF := ((reg18 >> 2) & 0b00000111) * 6      // Max: 7*6=42

	if r.IsFM() {
		dBu = int8(103) - int8(rssi) - int8(RF) - int8(IF) // Min: 103-127-42-42=-108, Max:103
	} else {
		dBu = int8(123) - int8(rssi) - int8(RF) - int8(IF) // Min: 123-127-42-42= -88, Max:123
	}
	return
}
