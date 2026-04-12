package fileio

import (
	"io"
	"os"
)

func ReadChunks(path string, chunkSize int) (chan []byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	out := make(chan []byte)

	go func() {
		defer file.Close()
		defer close(out)

		buf := make([]byte, chunkSize)

		for {
			n, err := file.Read(buf)
			if n > 0 {
				chunk := make([]byte, n)
				copy(chunk, buf[:n])
				out <- chunk
			}
			if err == io.EOF {
				break
			}
		}
	}()

	return out, nil
}

func WriteChunk(file *os.File, data []byte) error {
	_, err := file.Write(data)
	return err
}