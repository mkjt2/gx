package algorithms

import (
	"compress/flate"
	"io"
)

type Flate struct {
	Level int
}

func (z *Flate) Package() string {
	return "compress/flate"
}

func (z *Flate) GetReaderFn() func(io.Reader) (io.ReadCloser, error) {
	return func(r io.Reader) (io.ReadCloser, error) {
		return flate.NewReader(r), nil
	}
}

func (z *Flate) GetWriterFn() func(io.Writer) (io.WriteCloser, error) {
	return func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w, z.Level)
	}
}
