package random_test

import (
	"testing"

	"github.com/jakofys/fluid/internal/pkg/random"
)

func TestStringRandom(t *testing.T) {
	r := random.NewStringRandom()
	s := r.StringN(6)
	if len(s) != 6 {
		t.Errorf("no consider: %d chars instead of 6", len(s))
	}
}

func BenchmarkStringN(b *testing.B) {
	r := random.NewStringRandom()
	r.StringN(b.N)
}
