//go:build !librespot_decoders

package mp3

import (
	"io"

	librespot "github.com/devgianlu/go-librespot"
)

// Decoder is a silent placeholder: the real mpg123-backed implementation in
// decoder.go is compiled only with the librespot_decoders build tag. Jota
// streams audio through YouTube, so the librespot decoder is not linked.
type Decoder struct {
	SampleRate int32
	Channels   int32
}

func New(log librespot.Logger, r librespot.SizedReadAtSeeker, gain float32) (*Decoder, error) {
	return &Decoder{SampleRate: 44100, Channels: 2}, nil
}

func (d *Decoder) Read(p []float32) (int, error) { return 0, io.EOF }

func (d *Decoder) SetPositionMs(pos int64) error { return nil }

func (d *Decoder) PositionMs() int64 { return 0 }

func (d *Decoder) Close() error { return nil }
