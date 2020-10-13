package main

import (
	"fmt"
	"machine"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup

	time.Sleep(3 * time.Second)

	if err := machine.I2C0.Configure(machine.I2CConfig{Frequency: 100000}); err != nil {
		fmt.Println(err)
		return
	}

	messageFromTuner := make(chan []string)

	wait.Add(1)
	go fromTuner(messageFromTuner, &wait)

	messageToDisplay := make(chan []string)

	wait.Add(1)
	go addClock(messageFromTuner, messageToDisplay, &wait)

	wait.Add(1)
	go toDisplay(messageToDisplay, &wait)

	wait.Wait()
}
