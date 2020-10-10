package main

import "fmt"

func toSerial(c chan [2]string) {
	for i := 0; ; i++ {
		if s, more := <-c; more {
			if i%5 == 0 {
				fmt.Println(s[0], s[1])
			}
		} else {
			fmt.Println(`Channel closed for Serial`)
			break
		}
	}
}
