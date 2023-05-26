package algorithms

import (
	"compress/lzw"
	"io"
)

type LZW struct{}

func (l *LZW) Package() string {
	return "compress/lzw"
}

func (l *LZW) GetReaderFn() func(io.Reader) (io.ReadCloser, error) {
	return func(r io.Reader) (io.ReadCloser, error) {
		return lzw.NewReader(r, lzw.LSB, 8), nil
	}
}

func (l *LZW) GetWriterFn() func(io.Writer) (io.WriteCloser, error) {
	return func(w io.Writer) (io.WriteCloser, error) {
		return lzw.NewWriter(w, lzw.LSB, 8), nil
	}
}
