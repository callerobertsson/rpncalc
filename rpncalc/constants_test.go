package rpncalc

import "testing"

func TestConstants(t *testing.T) {

	cases := []struct {
		c     string
		exp   float64
		found bool
	}{
		// Test a few constants
		{"pi", 3.1415926535897932, true},
		{"sol", 299792458, true},
		{"urk", 0.0, false},
	}

	for _, c := range cases {
		r := New()

		found := r.pushConstant(c.c)
		if found != c.found {
			t.Errorf("Expected result %v, but got %v", c.found, found)
		}
		if r.Stack()[0] != c.exp {
			t.Errorf("Expected value %v, but got %v", c.exp, r.Stack()[0])
		}
	}
}
