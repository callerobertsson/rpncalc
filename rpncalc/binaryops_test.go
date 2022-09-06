package rpncalc

import (
	"math"
	"testing"
)

func TestBinaryOp(t *testing.T) {

	nice := func(float64, float64) (float64, error) {
		return 1.0, nil
	}
	evil := func(float64, float64) (float64, error) {
		return 0.0, errOverflow
	}

	cases := []struct {
		f   func(float64, float64) (float64, error)
		val float64
		err error
	}{
		{nice, 1.0, nil},
		{evil, 0.0, errOverflow},
	}

	for _, c := range cases {
		r := New()

		err := r.binaryOp(c.f)
		if err != c.err {
			t.Errorf("Expected result %v, but got %v", c.err, err)
		}
		if r.Stack()[0] != c.val {
			t.Errorf("Expected value %v, but got %v", c.val, r.Stack()[0])
		}
	}
}

func TestBinaryOpBasicOps(t *testing.T) {

	cases := []struct {
		name string
		f    func(*RpnCalc, string) error
		x    float64
		y    float64
		exp  float64
		err  error
	}{
		{"add 1 and -1", opAddition, 1, -1, 0, nil},
		{"add pi and 2", opAddition, 3.1415, 2.0, 5.1415, nil},
		{"add negative number", opAddition, 123456789.12345, -1, 123456788.12345, nil},
		{"add two large number will fail", opAddition, math.MaxFloat64, math.MaxFloat64 / 10, math.MaxFloat64 / 10, errOverflow},
		{"pi minus 2", opSubtraction, 3.1415, 2.0, 1.1415, nil},
		{"pi minus -2", opSubtraction, 3.1415, -2.0, 5.1415, nil},
		{"negative big number minus max float will fail", opSubtraction, -math.MaxFloat64 / 2, math.MaxFloat64, math.MaxFloat64, errOverflow},
		{"-max minus big negative number", opSubtraction, -math.MaxFloat64, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2, nil},
		{"-max minus big number will fail", opSubtraction, -math.MaxFloat64, math.MaxFloat64 / 2, math.MaxFloat64 / 2, errOverflow},
		{"big number minus -1", opSubtraction, 123456789.12345, -1, 123456790.12345, nil},
		{"101 times 10", opMultiplication, 101.0, 10.0, 1010.0, nil},
		{"multiply with negative number", opMultiplication, 1.23, -100.0, -123.0, nil},
		{"simple division", opDivision, 101.0, 10.0, 10.1, nil},
		{"divide by zero will fail", opDivision, 101.0, 0.0, 0.0, errDivisionByZero},
		{"power 2**0", opPower, 2, 0, 1, nil},
		{"power 2**1", opPower, 2, 1, 2, nil},
		{"power 2**-1", opPower, 2, -1, 0.5, nil},
		{"power 4**-2", opPower, 4, -2, 0.0625, nil},
		{"power 10**100", opPower, 10, 100, math.Pow(10, 100), nil},
		{"power 3**-4", opPower, 3, -4, math.Pow(3, -4), nil},
		{"1 modulus 2", opModulus, 1, 2, 1, nil},
		{"123 modulus 7", opModulus, 123, 7, 4, nil},
		{"10 modulus 0", opModulus, 10, 0, 4, errValueNotAllowed},
		{"10 modulus -5", opModulus, 10, -1, 4, errValueNotAllowed},

		// TODO: Add more cases for unary operators
	}

	for _, c := range cases {
		r := New()
		r.stack[1] = c.x
		r.stack[0] = c.y

		err := c.f(r, "")

		if err != nil && err == c.err {
			continue // Expected error
		}
		
		if err != c.err {
			t.Errorf("%q: Expected error %v for op(%v,%v), but got %v, val %v", c.name, c.err, c.x, c.y, err, r.stack[0])
			continue
		}

		got := r.stack[0]

		if !almostEqual(got, c.exp) {
			t.Errorf("%q: Expected result %v, but got %v", c.name, c.exp, got)
		}
	}

}
