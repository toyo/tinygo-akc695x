package main

import (
	"fmt"
	"sync"
)

func statusMonitor(c chan [2]string, w *sync.WaitGroup) {
	defer w.Done()

	toOLEDChan := make(chan [2]string)
	go toOLED(toOLEDChan)

	toSerialChan := make(chan [2]string)
	go toSerial(toSerialChan)

	for {
		if s, more := <-c; more {
			select {
			case toSerialChan <- s:
			default:
				fmt.Println(`Serial not available`)
			}

			select {
			case toOLEDChan <- s:
			default:
				fmt.Println(`OLED not available`)
			}
		} else {
			break
		}
	}
	close(toSerialChan)
	close(toOLEDChan)
}
