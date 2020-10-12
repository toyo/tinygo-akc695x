package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	akc695x "github.com/toyo/tinygo-akc695x"
)

func tunerStatus(r *akc695x.AKC695X, c1 chan []string, wait *sync.WaitGroup) {
	defer wait.Done()

	for {

		var s []string
		if r.IsPowerOn() {
			s = make([]string, 4)
			s[0] = r.GetFreqString()

			switch {
			case !r.IsSeekComplete():
				s[1] = `Seeking`
			case !r.IsTuned():
				s[1] = `Not tuned`
			default:
				s[1] = strconv.Itoa(int(r.GetRSSIdBuV())) + `dBu ` + strconv.Itoa(int(r.GetCNR())) + `dB`
			}

			s[2] = `Vol` + strconv.Itoa(int(r.GetVolume())) + ` ` +
				strconv.FormatFloat(float64(r.GetVCCMilliVolt())/1000, 'f', 2, 32) + `V`
		}
		select {
		case c1 <- s:
		default:
			fmt.Println("DataSink busy.")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
