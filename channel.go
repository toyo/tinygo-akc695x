package akc695x

import (
	"fmt"
	"strconv"
	"time"
)

func (r AKC695X) freqToChannel(kHz uint32) (ch uint16, isfm bool) {
	if kHz < 30000 {
		isfm = false
		if r.isMode3k() {
			ch = uint16(kHz / 3)
		} else {
			ch = uint16(kHz / 5)
		}
	} else {
		isfm = true
		ch = uint16((kHz - 30000) / 25)
	}
	return
}

func (r AKC695X) channelToFreq(ch uint16) (kHz uint32) {
	if r.isFM() {
		kHz = uint32(ch)*25 + 30000
	} else {
		if r.isMode3k() {
			kHz = uint32(ch) * 3
		} else {
			kHz = uint32(ch) * 5
		}
	}
	return
}

func (r AKC695X) GetFreq() (kHz uint32) {
	return r.channelToFreq(r.getChannel())
}

func (r AKC695X) GetFreqString() (freq string) {
	if f := r.GetFreq(); f < 30000 {
		freq = strconv.FormatUint(uint64(r.GetFreq()), 10) + `kHz`
	} else {
		freq = strconv.FormatFloat(float64(r.GetFreq())/1000, 'f', 1, 32) + `MHz`
	}
	return
}

func (r AKC695X) setFreq(kHz uint32) (err error) {
	channel, isfm := r.freqToChannel(kHz)
	if r.isFM() {
		if 150 <= kHz && kHz <= 285 {
			r.amband = 0
			r.setAMBand()
		} else if 520 <= kHz && kHz <= 1710 {
			r.amband = r.mwband
			r.setAMBand()
		} else if 3200 <= kHz && kHz <= 4100 {
			r.amband = 5
			r.setAMBand()
		} else if 4700 <= kHz && kHz <= 5600 {
			r.amband = 6
			r.setAMBand()
		} else if 5700 <= kHz && kHz <= 6400 {
			r.amband = 7
			r.setAMBand()
		} else if 6800 <= kHz && kHz <= 7600 {
			r.amband = 8
			r.setAMBand()
		} else if 9200 <= kHz && kHz <= 10000 {
			r.amband = 9
			r.setAMBand()
		} else if 11400 <= kHz && kHz <= 12200 {
			r.amband = 10
			r.setAMBand()
		} else if 13500 <= kHz && kHz <= 14300 {
			r.amband = 11
			r.setAMBand()
		} else if 15000 <= kHz && kHz <= 15900 {
			r.amband = 12
			r.setAMBand()
		} else if 17400 <= kHz && kHz <= 17900 {
			r.amband = 13
			r.setAMBand()
		} else if 18900 <= kHz && kHz <= 19700 {
			r.amband = 14
			r.setAMBand()
		} else if 21400 <= kHz && kHz <= 21900 {
			r.amband = 15
			r.setAMBand()
		} else {
			r.amband = r.mwband
			r.setAMBand()
		}
		r.setFM(isfm)
		if err = r.writeReg(0, 2); err != nil { // For check the step is 3kHz or 5kHz.
			return
		}
		channel, isfm = r.freqToChannel(kHz)
	} else {
		r.setFM(isfm)
		r.setFMBand()
		if err = r.writeReg(0, 1); err != nil { // Set FM Band.
			return
		}
		if r.fmband >= 7 {
			r.setChannelRange(r.fmlowch, r.fmhighch)
			if err = r.writeReg(4, 5); err != nil { // Set FM Band Width.
				return
			}
		}
	}

	r.setChannel(channel)

	return
}

func (r AKC695X) SetFreq(kHz uint32) (err error) {
	r.setTune(false)
	if err = r.writeReg(0, 0); err != nil {
		return
	}

	if err = r.setFreq(kHz); err != nil {
		return
	}

	r.setTune(true)
	if err = r.writeReg(0, 3); err != nil {
		return
	}

	for !r.IsSeekComplete() {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println(`Tuned`)

	return nil
}

func (r AKC695X) Seek(seekup bool) (err error) {
	r.setSeekUp(seekup)
	r.setSeek(true)
	if err = r.writeReg(0, 0); err != nil {
		return
	}

	for !r.IsSeekComplete() {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println(`SeekComplete`)

	r.setSeek(false)
	if err = r.writeReg(0, 0); err != nil {
		return
	}
	return nil
}
