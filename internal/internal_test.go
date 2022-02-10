package internal

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// Verify that DefaultEnvSetter implements EnvSetter.
var _ EnvSetter = DefaultEnvSetter{}

func TestByteCounter(t *testing.T) {
	s := "1234567890\n"
	in := strings.NewReader(s)
	scanner := bufio.NewScanner(in)
	counter := ByteCounter{}
	splitFunc := counter.Wrap(bufio.ScanLines)
	scanner.Split(splitFunc)
	for scanner.Scan() {
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("unexpected scanner error: %v", err)
	}
	expected := 11
	if counter.BytesRead != expected {
		t.Fatalf("unexpected count of bytes read. expected %v, got %v", expected, counter.BytesRead)
	}
}

func TestEnvSetter(t *testing.T) {
	setter := &DefaultEnvSetter{}
	if err := setter.Setenv("VAR", "VAL"); err != nil {
		t.Fatalf("unexpected error: '%v'", err)
	}
	if os.Getenv("VAR") != "VAL" {
		t.Fatal("unable to set env variable VAR=VAL")
	}
}
