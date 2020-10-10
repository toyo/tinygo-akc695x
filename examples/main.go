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

	if err := machine.I2C0.Configure(machine.I2CConfig{}); err != nil {
		fmt.Println(err)
		return
	}

	messageFromTunerToDisplay := make(chan [2]string)

	wait.Add(1)
	go tuner(messageFromTunerToDisplay, &wait)

	wait.Add(1)
	go statusMonitor(messageFromTunerToDisplay, &wait)

	wait.Wait()
}