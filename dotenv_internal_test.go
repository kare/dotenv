package dotenv

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	testData := []struct {
		input  string
		result map[string]string
	}{
		{
			"VAR=VAL\n",
			map[string]string{
				"VAR": "VAL",
			},
		},
		{
			"\n\nVAR=VAL\n\nFOO=BAR\nINT=123\nEMPTY=\n\n",
			map[string]string{
				"VAR":   "VAL",
				"FOO":   "BAR",
				"INT":   "123",
				"EMPTY": "",
			},
		},
		{
			"\n# comment\nVAR=VAL\n # comment2\nFOO=BAR\nINT=123\nEMPTY=\n\n",
			map[string]string{
				"VAR":   "VAL",
				"FOO":   "BAR",
				"INT":   "123",
				"EMPTY": "",
			},
		},
	}
	for _, tc := range testData {
		env := New()
		r := env.newReader(strings.NewReader(tc.input))
		n, err := r.Read(make([]byte, len(tc.input)))
		if err != nil {
			if err != io.EOF {
				t.Fatalf("error reading test input: %v", err)
			}
		}
		for k, v := range tc.result {
			if env.values[k] != v {
				t.Fatalf("expected key '%v' to have value '%v', got value '%v'", k, v, env.values[k])
			}
		}
		if len(tc.input) != n {
			t.Fatalf("expected to read %v bytes, but read %v", len(tc.input), n)
		}
	}
}

func TestReaderInvalidFormatError(t *testing.T) {
	testData := []struct {
		input string
		err   error
	}{
		{
			"VAR:VAL\n",
			errors.New("input line missing character '='"),
		},
	}
	for _, tc := range testData {
		env := New()
		r := env.newReader(strings.NewReader(tc.input))
		_, err := r.Read(make([]byte, len(tc.input)))
		if err.Error() != tc.err.Error() {
			t.Fatalf("expected error: '%v', got '%v'", tc.err, err)
		}
		if len(env.values) != 0 {
			t.Fatalf("expected map len '%v' to have be '%v'", 0, len(env.values))
		}
	}
}

type errorSetter struct{}

func (errorSetter) Setenv(k, v string) error {
	return errors.New("expected syscall error in test")
}

func TestSetenvError(t *testing.T) {
	env := New()
	env.setter = &errorSetter{}
	in := "VAR=VAL\n"
	r := env.newReader(strings.NewReader(in))
	_, _ = r.Read(make([]byte, len(in)))
	if err := env.Apply(); err.Error() != "expected syscall error in test" {
		t.Fatalf("unexpected error: '%v'", err)
	}
}

type errReader struct{}

func (e *errReader) Read(p []byte) (n int, err error) {
	return n, errors.New("expected i/o error")
}

func TestLoadError(t *testing.T) {
	env := New()
	if err := env.Load(&errReader{}); err.Error() != "expected i/o error" {
		t.Fatalf("expected i/o error, got '%v'", err)
	}
}

func TestReaderError(t *testing.T) {
	env := New()
	r := env.newReader(&errReader{})
	_, err := r.Read(make([]byte, 0))
	expected := "expected i/o error"
	if err.Error() != expected {
		t.Fatalf("expected error '%v', got '%v'", expected, err)
	}
}
