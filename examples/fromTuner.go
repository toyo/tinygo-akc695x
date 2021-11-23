package main

import (
	"machine"
	"sync"

	akc695x "github.com/toyo/tinygo-akc695x"
)

// Please set resetpin to 255 if there is no p_on connetction.
func fromTuner(c1 chan []string, givenWait *sync.WaitGroup) {
	defer givenWait.Done()

	r := akc695x.New(machine.I2C0)

	if err := r.Configure(akc695x.Config{
		MWBand: 2,     // JP Band
		FMBand: 7,     // JP Band
		FMLow:  76000, // JP Band
		FMHigh: 95000, // JP Band
	}); err != nil {
		panic(err)
	}

	var wait sync.WaitGroup

	wait.Add(1)
	go tunerStatus(&r, c1, &wait)

	wait.Add(1)
	go tunerControl("Command reference\n"+
		"Type P for Power On\n"+
		"Type p for Power Off\n"+
		"Type + for Volume up\n"+
		"Type - for Volume down\n"+
		"Type s for Seek incrementaly\n"+
		"Type r for Seek decrementaly\n",
		map[byte]func() error{
			'P': func() error { return r.PowerOn(79500, 40, true) }, // JODV-FM
			'p': r.PowerOff,
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
		}, &wait) // Please customize the frequency list. This sample is only for Tokyo.

	wait.Wait()
}
