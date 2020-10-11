package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	akc695x "github.com/toyo/tinygo-akc695x"
)

func tunerStatus(r *akc695x.AKC695X, c1 chan [2]string, wait *sync.WaitGroup) {
	defer wait.Done()

	for {
		var s [2]string
		if !r.IsSeekComplete() {
			s[0] = `Seeking`
		} else {
			s[0] = r.GetFreqString() + ` ` + strconv.Itoa(int(r.GetCNR())) + `dB`
		}

		s[1] = strconv.Itoa(int(r.GetRSSIdBuV())) + `dBu ` +
			`Vol` + strconv.Itoa(int(r.GetVolume())) + ` ` +
			strconv.FormatFloat(float64(r.GetVCCMilliVolt())/1000, 'f', 2, 32) + `V`

		select {
		case c1 <- s:
		default:
			fmt.Println("DataSink busy.")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
