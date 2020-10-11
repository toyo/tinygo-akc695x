package akc695x

import (
	"machine"
	"time"
)

// Config is config setting for AKC695X class.
type Config struct {
	I2CInterface     machine.I2C
	I2CAddr          uint8
	ExistResetPin    bool // if No ResetPin, set false, otherwise set true and ResetPin
	ResetPin         machine.Pin
	AMBand, FMBand   uint8
	FMLow, FMHigh    uint32
	VolumeControlI2C bool
	Volume           uint8
	InitialkHz       uint32
}

// AKC695X is instance of Chip.
type AKC695X struct {
	i2cinterface      machine.I2C
	i2caddr           uint8
	existresetpin     bool
	resetpin          machine.Pin
	mwband, fmband    uint8
	fmlowch, fmhighch uint16
	reg               []byte
}

// Address is I2C address of AKC695X
const Address = 0x10

func (r *AKC695X) reset() {
	if r.existresetpin {
		r.resetpin.Set(false)
		time.Sleep(10 * time.Millisecond)
		r.resetpin.Set(true)
		time.Sleep(10 * time.Millisecond)
	}
}

// Configure is initialize function for AKC695X.
func (r *AKC695X) Configure(config Config) (err error) {

	r.i2cinterface = config.I2CInterface
	r.i2caddr = config.I2CAddr
	r.existresetpin = config.ExistResetPin
	r.resetpin = config.ResetPin

	if r.existresetpin {
		r.resetpin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		r.reset()
	}

	if r.reg == nil {
		r.reg = make([]byte, 0x13)
	}
	if err = r.i2cinterface.ReadRegister(config.I2CAddr, 0, r.reg); err != nil {
		return
	}

	r.reg[0x00] |= 1 << 7  // PowerOn
	r.reg[0x00] &^= 1 << 2 // No Mute
	r.mwband = config.AMBand
	//r.setAMBand(config.AMBand)
	r.fmband = config.FMBand
	r.fmlowch = r.freqToChannelFM(config.FMLow)
	r.fmhighch = r.freqToChannelFM(config.FMHigh)
	//r.setFMBand(config.FMBand)

	if err := r.SetFreq(config.InitialkHz); err != nil {
		panic(err)
	}

	r.setVolumeControlI2C(config.VolumeControlI2C)
	err = r.writeReg(9, 9)
	if config.VolumeControlI2C {
		r.setVolume(config.Volume)
		err = r.writeReg(6, 6)
	}

	return
}
