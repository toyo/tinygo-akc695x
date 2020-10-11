package akc695x

import (
	"errors"
	"strconv"
)

// REG00-REG05, REG0B

// IsFM return the Band is FM or not.
func (r AKC695X) IsFM() bool {
	return r.reg[0x00]&(1<<6) != 0
}

func (r AKC695X) setAMBand(band uint8) {
	r.reg[0x00] &^= 1 << 6                                      // Set AM
	r.reg[0x01] = r.reg[0x01]&0b00000111 | (band&0b00011111)<<3 // Set AM Band

	if band == 0 || band == 2 {
		r.reg[0x02] |= 1 << 5 // setMode3k
	} else {
		r.reg[0x02] &^= 1 << 5 // setMode5k
	}
}

func (r AKC695X) setFMBand(band uint8) {
	r.reg[0x00] |= 1 << 6                                      // Set FM
	r.reg[0x01] = r.reg[0x01]&0b11111000 | (band & 0b00000111) // Set FM Band

}

// Seek seek the signal.
// if seekup is true, seek up. false, seek down.
func (r AKC695X) Seek(seekup bool) (err error) {
	r.reg[0x00] &^= 1 << 4 // Reset Seek
	if err = r.writeReg(0, 0); err != nil {
		return
	}

	if seekup { // Set Seek Direction
		r.reg[0x00] |= 1 << 3
	} else {
		r.reg[0x00] &^= 1 << 3
	}

	r.reg[0x00] |= 1 << 4 // Set Seek
	if err = r.writeReg(0, 0); err != nil {
		return
	}

	return nil
}

// SetFreq sets the frequency.
func (r AKC695X) SetFreq(kHz uint32) (err error) {
	var channel uint16

	if kHz < 30000 {
		switch {
		case 150 <= kHz && kHz <= 285:
			r.setAMBand(0)
		case 520 <= kHz && kHz <= 1710:
			r.setAMBand(r.mwband) // For set the step 3kHz or 5kHz.
		case 3200 <= kHz && kHz <= 4100:
			r.setAMBand(5)
		case 4700 <= kHz && kHz <= 5600:
			r.setAMBand(6)
		case 5700 <= kHz && kHz <= 6400:
			r.setAMBand(7)
		case 6800 <= kHz && kHz <= 7600:
			r.setAMBand(8)
		case 9200 <= kHz && kHz <= 10000:
			r.setAMBand(9)
		case 11400 <= kHz && kHz <= 12200:
			r.setAMBand(10)
		case 13500 <= kHz && kHz <= 14300:
			r.setAMBand(11)
		case 15000 <= kHz && kHz <= 15900:
			r.setAMBand(12)
		case 17400 <= kHz && kHz <= 17900:
			r.setAMBand(13)
		case 18900 <= kHz && kHz <= 19700:
			r.setAMBand(14)
		case 21400 <= kHz && kHz <= 21900:
			r.setAMBand(15)
		default:
			err = errors.New(`Out of AM Band: ` + strconv.FormatUint(uint64(kHz), 10) + `kHz`)
			return
		}

		r.reg[0x00] &^= 1 << 5                  // Reset TuneBit
		if err = r.writeReg(0, 2); err != nil { // Reset TuneBit and Set Mode3k
			return
		}

		channel = r.freqToChannelAM(kHz)
	} else {
		r.reg[0x00] &^= 1 << 5                  // Reset TuneBit
		if err = r.writeReg(0, 0); err != nil { // Reset TuneBit
			return
		}

		r.setFMBand(r.fmband)
		if r.fmband >= 7 { // User-defined Band
			r.reg[0x04] = byte(r.fmlowch >> 5)        // SetChannelRange
			r.reg[0x05] = byte((r.fmhighch >> 5) + 1) // SetChannelRange
		}

		channel = r.freqToChannelFM(kHz)
	}

	r.reg[0x00] |= 1 << 5                                    // Set TuneBit
	r.reg[0x02] = r.reg[0x02]&0xe0 | byte((channel>>8)&0x1f) // Set Upper Channel.
	r.reg[0x03] = byte(channel & 0xff)                       // Set Lower Channel.

	if r.IsFM() && r.fmband >= 7 {
		err = r.writeReg(0, 5) // Tune User-defined Band
	} else {
		err = r.writeReg(0, 3) // Tune Pre-defined Band
	}

	return
}

// setRef32kMode sets the Crystal frequency setting
// true: 32.768KHz
// false: 12MHz
func (r AKC695X) setRef32kMode(b bool) {
	if b {
		r.reg[0x02] |= 1 << 6
	} else {
		r.reg[0x02] &^= 1 << 6
	}
}

// setFMSeekStep Setting of FM seek step
// 0 -25KHz
// 1 -50KHz
// 2 -100KHz
// 3 -200KHz
func (r AKC695X) setFMSeekStep(step uint8) {
	r.reg[0x0b] = r.reg[0x0b]&0b11001111 | (step&0x03)<<4
}
