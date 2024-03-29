package rpncalc

import "testing"

func TestUnaryOp(t *testing.T) {

	nice := func(float64, string) (float64, error) {
		return 1.0, nil
	}
	evil := func(float64, string) (float64, error) {
		return 0.0, errOverflow
	}

	cases := []struct {
		f   func(float64, string) (float64, error)
		val float64
		err error
	}{
		{nice, 1.0, nil},
		{evil, 0.0, errOverflow},
	}

	for _, c := range cases {
		r := New()

		err := r.unaryOp(c.f)
		if err != c.err {
			t.Errorf("Expected result %v, but got %v", c.err, err)
		}
		if r.Stack()[0] != c.val {
			t.Errorf("Expected value %v, but got %v", c.val, r.Stack()[0])
		}
	}
}

func TestUnaryOpBasicOperations(t *testing.T) {

	cases := []struct {
		name string
		f    func(*RpnCalc, string) error
		v    float64
		exp  float64
		err  error
	}{
		{"negate 1", opNegate, 1, -1, nil},
		{"negate -pi", opNegate, -3.14, 3.14, nil},
		{"invers of 10", opInverse, 10.0, 0.1, nil},
		{"inverse of 0 should fail", opInverse, 0.0, 0.0, errDivisionByZero},
		{"square 0", opSquare, 0.0, 0.0, nil},
		{"square 3", opSquare, 3.0, 9.0, nil},
		{"square -3", opSquare, -3.0, 9.0, nil},
		{"square overflow", opSquare, 1e+155, 1e+155, errOverflow},
		{"square root of 9", opSquareRoot, 9, 3, nil},
		{"square root of -9", opSquareRoot, -9, -9, errNaN},
		{"dec 1 as bin", opDecToBin, 1, 1, nil},
		{"dec 2 as bin", opDecToBin, 2, 10, nil},
		{"dec 3 as bin", opDecToBin, 3, 11, nil},
		{"dec 255 as bin", opDecToBin, 255, 11111111, nil},
		{"dec 256 as bin", opDecToBin, 256, 100000000, nil},
		{"dec 678901234 as bin", opDecToBin, 67890, 10000100100110010, nil},
		{"dec to bin overflow", opDecToBin, 6789012345, 0, errOverflow},
		{"bin 0 to dec", opBinToDec, 0, 0, nil},
		{"bin 1 to dec", opBinToDec, 1, 1, nil},
		{"bin 10 to dec", opBinToDec, 10, 2, nil},
		{"bin 1001001 to dec", opBinToDec, 1001001, 73, nil},
		{"bin to dec not a binary number", opBinToDec, 123, 0, errNoBinaryNumber},

		// TODO: Add more cases for unary operators
	}

	for _, c := range cases {
		r := New()
		r.stack[0] = c.v

		err := c.f(r, "")

		if err != nil && err == c.err {
			continue // Expected error
		}

		if err != c.err {
			t.Errorf("%q: Expected error %v, but got %v", c.name, c.err, err)
			continue
		}

		got := r.stack[0]

		if got != c.exp {
			t.Errorf("%q: Expected result %v, but got %v", c.name, c.exp, got)
		}
	}

}
