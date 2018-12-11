package dotenv

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"kkn.fi/dotenv/internal"
)

// ErrInvalidFormat describes an error condition where input line is missing '=' character.
var ErrInvalidFormat = errors.New("input line missing character '='")

type (
	// Env holds variable-value pairs.
	Env struct {
		values map[string]string
		setter internal.EnvSetter
	}
	// reader reads variable-value pairs to memory from io.Reader.
	reader struct {
		io.Reader
		r io.Reader
		e *Env
	}
)

// New creates a new Env.
func New() *Env {
	return &Env{
		values: make(map[string]string),
		setter: &internal.DefaultEnvSetter{},
	}
}

// Load reads given io.Reader to Env.
func (e *Env) Load(r io.Reader) error {
	reader := e.newReader(r)
	_, err := reader.Read(make([]byte, 0))
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

// Apply sets currently loaded variable-value pairs to shell environment.
func (e *Env) Apply() error {
	for k, v := range e.values {
		if err := e.setter.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}

// newReader reads variable-value pairs from given in Reader.
func (e *Env) newReader(r io.Reader) *reader {
	return &reader{
		r: r,
		e: e,
	}
}

func (r *reader) Read(p []byte) (int, error) {
	scanner := bufio.NewScanner(r.r)
	counter := internal.ByteCounter{}
	splitFunc := counter.Wrap(bufio.ScanLines)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}
		keyVal := strings.Split(line, "=")
		if len(keyVal) != 2 {
			return counter.BytesRead, ErrInvalidFormat
		}
		r.e.values[keyVal[0]] = keyVal[1]
	}
	if err := scanner.Err(); err != nil {
		return counter.BytesRead, err
	}
	return counter.BytesRead, io.EOF
}
