package main

import (
	"fmt"
	"strings"
)

func toSerial(c chan []string) {
	for i := 0; ; i++ {
		if s, more := <-c; more {
			if i%10 == 0 {
				fmt.Println(strings.Join(s, ` `))
			}
		} else {
			fmt.Println(`Channel closed for Serial`)
			break
		}
	}
}
