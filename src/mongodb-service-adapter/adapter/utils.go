package adapter

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net"
	"sort"
	"strings"
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

// SortAddresses sorts list of ip addresses.
// It panics if any address cannot be resolved.
// If port is missing it's considered as zero.
func SortAddresses(s []string) {
	sort.Slice(s, func(i, j int) bool {
		return addrn(s[i]) < addrn(s[j])
	})
}

func addrn(addr string) int {
	if !strings.Contains(addr, ":") {
		addr = addr + ":0"
	}

	n := 0
	a, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}

	for _, b := range a.IP {
		n += int(b)
	}
	return n
}
