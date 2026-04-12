package integrity

import (
	"crypto/sha256"
	"io"
	"os"
)

func FileHash(path string) ([32]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return [32]byte{}, err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)

	var result [32]byte
	copy(result[:], hash.Sum(nil))
	return result, err
}