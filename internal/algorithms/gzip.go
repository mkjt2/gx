package algorithms

import (
	"compress/gzip"
	"io"
)

type Gzip struct {
	Level int
}

func (g *Gzip) GetReaderFn() func(io.Reader) (io.ReadCloser, error) {
	return func(r io.Reader) (io.ReadCloser, error) {
		return gzip.NewReader(r)
	}
}

func (g *Gzip) GetWriterFn() func(io.Writer) (io.WriteCloser, error) {
	return func(w io.Writer) (io.WriteCloser, error) {
		return gzip.NewWriterLevel(w, g.Level)
	}
}

func (g *Gzip) Package() string {
	return "compress/gzip"
}
