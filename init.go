package akc695x

import (
	"machine"
	"time"
)

type AKC695XConfig struct {
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

type AKC695X struct {
	i2cinterface      machine.I2C
	i2caddr           uint8
	existresetpin     bool
	resetpin          machine.Pin
	amband, fmband    uint8
	mwband            uint8
	fmlowch, fmhighch uint16
	reg               []byte
}

const Address = 0x10

func (r *AKC695X) reset() {
	if r.existresetpin {
		r.resetpin.Set(false)
		time.Sleep(10 * time.Millisecond)
		r.resetpin.Set(true)
		time.Sleep(10 * time.Millisecond)
	}
}

func (r *AKC695X) Configure(config AKC695XConfig) (err error) {

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

	r.powerOn()
	r.setMute(false)
	r.mwband = config.AMBand
	r.amband = r.mwband
	r.setAMBand()
	r.fmband = config.FMBand
	r.fmlowch, _ = r.freqToChannel(config.FMLow)
	r.fmhighch, _ = r.freqToChannel(config.FMHigh)
	r.setFMBand()

	if err := r.SetFreq(config.InitialkHz); err != nil {
		panic(err)
	}

	if config.VolumeControlI2C {
		r.setVolume(config.Volume)
		err = r.writeReg(6, 6)
		r.setVolumeControlI2C()
		err = r.writeReg(9, 9)
	}

	return
}
