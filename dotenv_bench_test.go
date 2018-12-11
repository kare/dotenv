package dotenv

import (
	"strings"
	"testing"
)

func BenchmarkReader(b *testing.B) {
	b.ReportAllocs()
	env := New()
	in := strings.NewReader("KEY=VAL\nVAR=VAL\nFOO=BAR\n")
	r := env.newReader(in)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Read(make([]byte, in.Len()))
	}
}
