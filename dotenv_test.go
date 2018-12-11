package dotenv_test

import (
	"os"
	"strings"
	"testing"

	"kkn.fi/dotenv"
)

func TestApply(t *testing.T) {
	envs := []string{
		"VAR=VAL",
		"FOO=BAR",
	}
	env := dotenv.New()
	input := strings.Join(envs, "\n")
	input = input + "\n"
	env.Load(strings.NewReader(input))
	if err := env.Apply(); err != nil {
		t.Fatalf("unexpected error: '%v'", err)
	}
	for _, v := range envs {
		s := strings.Split(v, "=")
		if os.Getenv(s[0]) != s[1] {
			t.Fatalf("expected variable %v to have value '%v', got '%v'", s[0], s[1], os.Getenv(s[0]))
		}
	}
}
