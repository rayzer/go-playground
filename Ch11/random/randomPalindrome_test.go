package random

import (
	"github.com/rayli/go-playground/Ch11/word2"
	"math/rand"
	"testing"
	"time"
)

func Test_randomPalindrome(t *testing.T) {
	type args struct {
		rng *rand.Rand
	}

	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	tests := []struct {
		args args
		want bool
	}{
		{args{rng}, true},
	}

	for _, tt := range tests {
		t.Run("randomPalindrome", func(t *testing.T) {
			generated := randomPalindrome(tt.args.rng)
			if !word.IsPalindrome(generated) {
				t.Errorf("IsPalindrome(%q) == false", generated)
			}
		})
	}
}
