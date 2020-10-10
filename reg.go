package akc695x

import (
	"encoding/binary"
)

func (r AKC695X) writeReg(from, to uint8) error {
	return r.i2cinterface.WriteRegister(r.i2caddr, from, r.reg[from:to+1])
}

func (r AKC695X) readRegister8(reg uint8) (b byte) {
	val := make([]byte, 1)
	err := r.i2cinterface.ReadRegister(r.i2caddr, reg, val)
	if err == nil {
		b = val[0]
	} else {
		b = 0
	}
	return
}

func (r AKC695X) readRegister16(reg uint8) (b16 uint16) {
	val := make([]byte, 2)
	err := r.i2cinterface.ReadRegister(r.i2caddr, reg, val)
	if err == nil {
		b16 = binary.BigEndian.Uint16(val)
	} else {
		b16 = 0
	}
	return
}
