// +build !ssd1306

package main

func toOLED(c chan [2]string) {
	for {
		if _, more := <-c; !more {
			break
		}
	}
}
