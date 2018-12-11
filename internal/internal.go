package internal

import (
	"bufio"
	"os"
)

type (
	// ByteCounter can Wrap a bufio.SplitFunc and count bytes that go through bufio.Scanner.
	ByteCounter struct {
		BytesRead int
	}
	// EnvSetter applies variable-value pair to OS shell environment.
	EnvSetter interface {
		Setenv(string, string) error
	}
	// DefaultEnvSetter delegates calls os.Setenv(string, string).
	DefaultEnvSetter struct{}
)

// Wrap wraps given split function and counts bytes streamed through.
func (b *ByteCounter) Wrap(split bufio.SplitFunc) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (int, []byte, error) {
		adv, tok, err := split(data, atEOF)
		b.BytesRead += adv
		return adv, tok, err
	}
}

// Setenv sets OS shell environment variable.
func (DefaultEnvSetter) Setenv(k, v string) error {
	return os.Setenv(k, v)
}
