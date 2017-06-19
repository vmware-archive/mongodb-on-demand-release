package adapter

import (
	"crypto/rand"
	"errors"
	"fmt"
)

// GenerateString generates a random string or panics
// if something goes wrong.
func GenerateString(l int) (string, error) {
	b := make([]byte, l)
	for i := l; i != 0; {
		n, err := rand.Read(b)
		if err != nil {
			return "", err
		}
		if n == 0 {
			return "", errors.New("couldn't read from crypto/rand")
		}

		i -= n
	}

	return fmt.Sprintf("%x", b)[:l], nil
}
