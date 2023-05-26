package algorithms

import (
	"github.com/klauspost/compress/zstd"
	"io"
)

type ZSTD struct {
	Level zstd.EncoderLevel
}

func (z *ZSTD) Package() string {
	return "github.com/klauspost/compress/zstd"
}

func (z *ZSTD) GetReaderFn() func(io.Reader) (io.ReadCloser, error) {
	return func(r io.Reader) (io.ReadCloser, error) {
		decoder, err := zstd.NewReader(r)
		if err != nil {
			return nil, err
		}
		return decoder.IOReadCloser(), nil
	}
}

func (z *ZSTD) GetWriterFn() func(io.Writer) (io.WriteCloser, error) {
	return func(w io.Writer) (io.WriteCloser, error) {
		encoder, err := zstd.NewWriter(w, zstd.WithEncoderLevel(z.Level))
		if err != nil {
			return nil, err
		}
		return encoder, nil
	}
}
