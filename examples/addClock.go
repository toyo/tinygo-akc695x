package main

import (
	"fmt"
	"machine"
	"sync"

	"tinygo.org/x/drivers/ds3231"
)

// if message from tuner is nil, add timestamp using DS3231
func addClock(chin chan []string, chout chan []string, givenWait *sync.WaitGroup) {
	defer givenWait.Done()

	ppsfall := make(chan bool)
	rtc := ds3231.New(machine.I2C0)

	rtc.Configure()

	if valid := rtc.IsTimeValid(); valid {
		if running := rtc.IsRunning(); running {
			ctrl := make([]byte, 1)
			machine.I2C0.ReadRegister(0x68, 0x0e, ctrl)
			ctrl[0x00] = ctrl[0x00] & 0b11100011 // for 1Hz
			machine.I2C0.WriteRegister(0x68, 0x0e, ctrl)

			machine.D2.SetInterrupt(machine.PinFalling, // SQW of DS3231 is connected to D2
				func(p machine.Pin) {
					select {
					case ppsfall <- true:
					default:
					}
				})
		} else {
			fmt.Println("Not running RTC")
		}
	} else {
		fmt.Println(`RTC not valud`)
	}

	var laststr []string

	for {
		select {
		case s, more := <-chin:
			if !more { // channel closed.
				return
			}
			if !sameStringaSlice(s, laststr) {
				if s != nil {
					chout <- s
				}
				laststr = s
			}
		case <-ppsfall:
			if laststr == nil {
				if dt, err := rtc.ReadTime(); err == nil {
					chout <- []string{dt.Format(`2006-01-02`), dt.Format(`15:04:05`)}
				}
			}
		}
	}
}

func sameStringaSlice(s1, s2 []string) bool {
	if s1 == nil && s2 == nil {
		return true
	}
	if s1 == nil || s2 == nil {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
