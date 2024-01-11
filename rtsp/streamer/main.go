package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bluenviron/gortsplib/v4"
	"github.com/bluenviron/gortsplib/v4/pkg/description"
	"github.com/bluenviron/gortsplib/v4/pkg/format"
	"github.com/pion/randutil"
	"github.com/pion/rtp"
)

type Payloader struct {
}

func main() {
	file, err := os.Open("./audio/Candyland.mp3")
	if err != nil {
		panic(err)
	}

	desc := &description.Session{
		Medias: []*description.Media{{
			Type:    description.MediaTypeAudio,
			Formats: []format.Format{&format.MPEG1Audio{}},
		}},
	}

	c := gortsplib.Client{}
	err = c.StartRecording("rtsp://localhost:8554/mystream", desc)
	if err != nil {
		panic("recording: " + err.Error())
	}

	buf := make([]byte, 1024)
	// packetizer := rtp.NewPacketizer(1500, 14, 0x1234ABCD, nil, rtp.NewRandomSequencer(), 90000)
	sequencer := rtp.NewRandomSequencer()
	randomizer := randutil.NewMathRandomGenerator()
	// var pkt rtp.Packet

	ticker := time.NewTicker(time.Second / 90000)

	is_empty := false

	go func() {
		for {
			select {
			case <-ticker.C:
				if is_empty {
					return
				}

				n, err := file.Read(buf)
				if err != nil {
					panic("read: " + err.Error())
				}

				err = c.WritePacketRTP(desc.Medias[0], &rtp.Packet{
					Header: rtp.Header{
						Version:        2,
						Padding:        false,
						Extension:      false,
						Marker:         n > 1024,
						PayloadType:    14,
						SequenceNumber: sequencer.NextSequenceNumber(),
						Timestamp:      randomizer.Uint32(),
						SSRC:           0x1234ABCD,
						CSRC:           []uint32{},
					},
					Payload: buf[:],
				})
				if err != nil {
					panic("write: " + err.Error())
				}

				if n < 1024 {
					is_empty = true
					fmt.Println("done")
				}

			}
		}
	}()

	// fmt.Println("Completed recording")
	for !is_empty {

	}
}
