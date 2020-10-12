package akc695x

import (
	"strconv"
)

func freqToChannelFM(kHz uint32) (ch uint16) {
	return uint16((kHz - 30000) / 25)
}

func (r AKC695X) freqToChannelAM(kHz uint32) (ch uint16) {
	if r.isMode3k() {
		ch = uint16(kHz / 3)
	} else {
		ch = uint16(kHz / 5)
	}
	return
}

func (r AKC695X) channelToFreq(ch uint16) (kHz uint32) {
	if r.IsFM() {
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

// GetFreq returns the frequency of receiving.
func (r AKC695X) GetFreq() (kHz uint32) {
	return r.channelToFreq(r.getChannel())
}

// GetFreqString returns the string of GetFreq.
func (r AKC695X) GetFreqString() (freq string) {
	if f := r.GetFreq(); f < 30000 {
		freq = strconv.FormatUint(uint64(r.GetFreq()), 10) + `kHz`
	} else {
		freq = strconv.FormatFloat(float64(r.GetFreq())/1000, 'f', 1, 32) + `MHz`
	}
	return
}
