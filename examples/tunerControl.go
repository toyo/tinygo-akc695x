package main

import (
	"fmt"
	"machine"
	"sync"
	"time"
)

func tunerControl(message string, funcMap map[byte]func() error, wait *sync.WaitGroup) {
	defer wait.Done()

	fmt.Print(message)
	rbuf := make([]byte, 1)

	for {
		if len, err := machine.USB.Read(rbuf); err == nil {
			if len != 0 {
				if f, ok := funcMap[rbuf[0]]; ok {
					if err := f(); err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(message)
				}
			} else {
				time.Sleep(500 * time.Millisecond)
			}
		} else {
			fmt.Println(err)
			break
		}
	}
}
