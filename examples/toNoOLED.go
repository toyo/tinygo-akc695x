// +build !ssd1306,!sh1106

package main

func toOLED(c chan []string) {
	for {
		if _, more := <-c; !more {
			break
		}
	}
}
