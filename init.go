package akc695x

import (
	"machine"
)

type AKC695XConfig struct {
	I2CInterface     machine.I2C
	I2CAddr          uint8
	AMBand, FMBand   uint8
	FMLow, FMHigh    uint32
	VolumeControlI2C bool
}

type AKC695X struct {
	i2cinterface      machine.I2C
	i2caddr           uint8
	reg               []byte
	amband, fmband    uint8
	mwband            uint8
	fmlowch, fmhighch uint16
}

const Address = 0x10

func (r *AKC695X) Configure(config AKC695XConfig) (err error) {

	r.i2cinterface = config.I2CInterface
	r.i2caddr = config.I2CAddr

	if r.reg == nil {
		r.reg = make([]byte, 0x0e)
	}
	if err = r.i2cinterface.ReadRegister(config.I2CAddr, 0, r.reg); err != nil {
		return
	}

	r.powerOn()
	r.mwband = config.AMBand
	r.amband = r.mwband
	r.setAMBand()
	r.fmband = config.FMBand
	r.fmlowch, _ = r.freqToChannel(config.FMLow)
	r.fmhighch, _ = r.freqToChannel(config.FMHigh)
	r.setFMBand()

	if config.VolumeControlI2C {
		r.setVolumeControlI2C()
		err = r.writeReg(9, 9)
	}

	return
}
