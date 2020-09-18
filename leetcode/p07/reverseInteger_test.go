package p07

import "testing"

func Test(t *testing.T) {
	if 123 != reverse(321) {
		t.Error("321 reverse shall be 123, but: ", reverse(321))
	}

	if 1112332 != reverse(2332111) {
		t.Error("2332111 reverse shall be 1112332, but: ", reverse(2332111))
	}

	if 21 != reverse(120) {
		t.Error("120 reverse shall be 21, but: ", reverse(120))
	}

	if 0 != reverse(1534236469) {
		t.Error("1534236469 reverse shall be 0, but: ", reverse(1534236469))
	}

	if -321 != reverse(-123) {
		t.Error("-123 reverse shall be -321, but: ", reverse(-123))
	}

}
