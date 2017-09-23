package rpncalc

import (
	"math"
	"testing"
)

func TestUnaryOp(t *testing.T) {

	cases := []struct {
		f   func(float64) (float64, error)
		v   float64
		exp float64
		err error
	}{
		{opNegate, 1, -1, nil},
		{opNegate, -3.14, 3.14, nil},
		{opInverse, 10.0, 0.1, nil},
		{opInverse, 0.0, 0.0, errDivisionByZero},
		// TODO: Add more cases for unary operators
	}

	for i, c := range cases {
		got, err := c.f(c.v)

		if err != c.err {
			t.Errorf("Case %d: Expected error %v, but got %v", i, c.err, err)
		}

		if got != c.exp {
			t.Errorf("Case %d: Expected result %v, but got %v", i, c.exp, got)
		}
	}

}

func TestBinaryOp(t *testing.T) {

	cases := []struct {
		f   func(float64, float64) (float64, error)
		x   float64
		y   float64
		exp float64
		err error
	}{
		{opAddition, 1, -1, 0, nil},
		{opAddition, 3.1415, 2.0, 5.1415, nil},
		{opAddition, 123456789.12345, -1, 123456788.12345, nil},
		{opSubtraction, 3.1415, 2.0, 1.1415, nil},
		{opSubtraction, 3.1415, -2.0, 5.1415, nil},
		{opSubtraction, 123456789.12345, -1, 123456790.12345, nil},
		{opMultiplication, 101.0, 10.0, 1010.0, nil},
		{opMultiplication, 1.23, -100.0, -123.0, nil},
		{opDivision, 101.0, 10.0, 10.1, nil},
		{opDivision, 101.0, 0.0, 0.0, errDivisionByZero},

		// TODO: Add more cases for unary operators
	}

	for i, c := range cases {
		got, err := c.f(c.x, c.y)

		if err != c.err {
			t.Errorf("Case %d: Expected error %v, but got %v", i, c.err, err)
		}

		if !almostEqual(got, c.exp) {
			t.Errorf("Case %d: Expected result %v, but got %v", i, c.exp, got)
		}
	}

}

func almostEqual(x, y float64) bool {
	delta := 0.000000000000001
	//fmt.Printf("diff %v, delta %v\n", math.Abs(x-y), delta)
	return math.Abs(x-y) < delta
}
