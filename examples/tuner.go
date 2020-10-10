package main

import (
	"machine"
	"sync"

	akc695x "github.com/toyo/tinygo-akc695x"
)

func tuner(c1 chan [2]string, wait *sync.WaitGroup) {
	defer wait.Done()

	var r akc695x.AKC695X

	if err := r.Configure(akc695x.AKC695XConfig{
		I2CInterface:     machine.I2C0,
		I2CAddr:          akc695x.Address,
		AMBand:           2,     // JP Band
		FMBand:           7,     // JP Band
		FMLow:            76000, // JP Band
		FMHigh:           95000, // JP Band
		VolumeControlI2C: true,
	}); err != nil {
		panic(err)
	}

	if err := r.SetFreq(79500); err != nil { // JODV-FM
		panic(err)
	}

	if err := r.SetVolume(50); err != nil {
		panic(err)
	}

	var wait1 sync.WaitGroup

	wait1.Add(1)
	go tunerStatus(&r, c1, &wait1)

	wait1.Add(1)
	go tunerControl("Command reference\n"+
		"Type + to Volume up\n"+
		"Type - to Volume down\n"+
		"Type s to Seek incrementaly\n"+
		"Type r to Seek decrementaly\n",
		map[byte]func() error{
			'+': r.VolumeUp,
			'-': r.VolumeDown,
			's': func() error { return r.Seek(true) },
			'r': func() error { return r.Seek(false) },
			't': func() error { return r.SetFreq(80000) }, // JOAU-FM Tokyo Tower
			'h': func() error { return r.SetFreq(86600) }, // JOAU-FM Hinohara
			'b': func() error { return r.SetFreq(90500) }, // JOKR
			'q': func() error { return r.SetFreq(91600) }, // JOQR
			'l': func() error { return r.SetFreq(93000) }, // JOLF
			'3': func() error { return r.SetFreq(3925) },  // JOZ
			'j': func() error { return r.SetFreq(81300) }, // JOAV-FM
			'y': func() error { return r.SetFreq(84700) }, // JOTU-FM
			'5': func() error { return r.SetFreq(79500) }, // JODV-FM
			'c': func() error { return r.SetFreq(78000) }, // JOGV-FM
		}, &wait1) // Please customize the frequency list. This sample is only for Tokyo.

	wait1.Wait()
}
