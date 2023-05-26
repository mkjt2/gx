package algorithms

import (
	"compress/zlib"
	"io"
)

type ZLib struct {
	Level int
}

func (g *ZLib) Package() string {
	return "compress/zlib"
}

func (g *ZLib) GetReaderFn() func(io.Reader) (io.ReadCloser, error) {
	return func(r io.Reader) (io.ReadCloser, error) {
		return zlib.NewReader(r)
	}
}

func (g *ZLib) GetWriterFn() func(io.Writer) (io.WriteCloser, error) {
	return func(w io.Writer) (io.WriteCloser, error) {
		return zlib.NewWriterLevel(w, g.Level)
	}
}
