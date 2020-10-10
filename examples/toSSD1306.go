// +build ssd1306

package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func toOLED(c chan [2]string) {

	d := ssd1306.NewI2C(machine.I2C0)

	time.Sleep(2500 * time.Millisecond)
	d.Configure(ssd1306.Config{
		Width: 128, Height: 32,
		Address: ssd1306.Address_128_32,
	})

	for {
		if s, more := <-c; more {
			var black = color.RGBA{1, 1, 1, 255}
			d.ClearBuffer()
			tinyfont.WriteLine(&d, &freesans.Regular9pt7b, 0, 15, s[0], black)
			tinyfont.WriteLine(&d, &freesans.Regular9pt7b, 0, 31, s[1], black)
			d.Display()
		} else {
			fmt.Println(`Channel closed for OLED`)
			break
		}
	}
}
