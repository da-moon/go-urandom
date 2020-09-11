package urandom

import (
	"crypto/rand"
	"io"

	"github.com/palantir/stacktrace"
)

// GenerateRandomBytes is used to generate random bytes of given size.
func GenerateRandomBytes(size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		err = stacktrace.Propagate(err, "failed to read '%v' random bytes", size)
		return nil, err
	}
	return buf, nil
}
