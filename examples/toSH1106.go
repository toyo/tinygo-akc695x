// +build sh1106

package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"

	sh1106 "github.com/toyo/tinygo-sh1106"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func toOLED(c chan []string) {

	d := sh1106.NewI2C(machine.I2C0)

	time.Sleep(500 * time.Millisecond)
	d.Configure(sh1106.Config{
		Width: 128, Height: 64,
		Address: sh1106.Address,
	})

	for {
		if s, more := <-c; more {
			var black = color.RGBA{1, 1, 1, 255}
			d.ClearBuffer()
			for i := range s {
				if i < 4 {
					tinyfont.WriteLine(&d, &freesans.Regular9pt7b, 0, int16(i)*16+15, s[i], black)
				}
			}
			d.Display()
		} else {
			fmt.Println(`Channel closed for OLED`)
			break
		}
	}
}
