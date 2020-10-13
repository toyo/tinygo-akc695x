package akc695x

import (
	"tinygo.org/x/drivers"
)

// Config is config setting for AKC695X class.
type Config struct {
	I2CAddr        uint8
	MWBand, FMBand uint8
	FMLow, FMHigh  uint32
}

// AKC695X is instance of Chip.
type AKC695X struct {
	i2cinterface      drivers.I2C
	i2caddr           uint8
	mwband, fmband    uint8
	fmlowch, fmhighch uint16
	reg               []byte
}

// Address is I2C address of AKC695X
const Address = 0x10

// New defined I2C Bus.
func New(bus drivers.I2C) AKC695X {
	return AKC695X{
		i2cinterface: bus,
		i2caddr:      Address,
	}
}

// Configure is initialize function for AKC695X.
func (r *AKC695X) Configure(config Config) (err error) {

	if config.I2CAddr != 0 {
		r.i2caddr = config.I2CAddr
	}

	r.mwband = config.MWBand
	r.fmband = config.FMBand
	r.fmlowch = freqToChannelFM(config.FMLow)
	r.fmhighch = freqToChannelFM(config.FMHigh)

	if r.reg == nil {
		r.reg = make([]byte, 0x13)
		if err = r.i2cinterface.ReadRegister(r.i2caddr, 0, r.reg); err != nil {
			return
		}
	}

	return
}

// PowerOn turns on tuner
func (r AKC695X) PowerOn(khz uint32, Volume uint8, VolumeControlI2C bool) (err error) {

	r.reg[0x00] |= 1 << 7  // PowerOn
	r.reg[0x00] &^= 1 << 2 // No Mute

	if err := r.SetFreq(khz); err != nil {
		panic(err)
	}

	r.setVolumeControlI2C(VolumeControlI2C)
	err = r.writeReg(9, 9)
	if VolumeControlI2C {
		r.setVolume(Volume)
		err = r.writeReg(6, 6)
	}
	return
}

// PowerOff turns off tuner
func (r AKC695X) PowerOff() (err error) {
	r.reg[0x00] &^= 1 << 7 // PowerOff
	return r.writeReg(0, 0)
}

// IsPowerOn returns wheather PowerOn or not.
func (r AKC695X) IsPowerOn() bool {
	return (r.reg[0x00] & (1 << 7)) != 0
}
