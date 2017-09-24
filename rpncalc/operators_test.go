package rpncalc

import (
	"math"
	"testing"
)

func TestUnaryOp(t *testing.T) {

	nice := func(float64) (float64, error) {
		return 1.0, nil
	}
	evil := func(float64) (float64, error) {
		return 0.0, errOverflow
	}

	cases := []struct {
		f   func(float64) (float64, error)
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

func TestUnaryOpBasicOperations(t *testing.T) {

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
		{opSquare, 0.0, 0.0, nil},
		{opSquare, 3.0, 9.0, nil},
		{opSquare, -3.0, 9.0, nil},
		{opSquare, 1e+155, 0.0, errOverflow},
		// TODO: Add more cases for unary operators
	}

	for i, c := range cases {
		got, err := c.f(c.v)

		if err != c.err {
			t.Errorf("Case %d: Expected error %v, but got %v", i, c.err, err)
			continue
		}

		if got != c.exp {
			t.Errorf("Case %d: Expected result %v, but got %v", i, c.exp, got)
		}
	}

}

func TestBinaryOpBasicOps(t *testing.T) {

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
		{opAddition, math.MaxFloat64 - 1.0, 2.0, 0.0, errOverflow},
		{opSubtraction, 3.1415, 2.0, 1.1415, nil},
		{opSubtraction, 3.1415, -2.0, 5.1415, nil},
		{opSubtraction, -1.0, math.MaxFloat64, 0.0, errOverflow},
		{opSubtraction, -math.MaxFloat64, -2.0, 2 - math.MaxFloat64, nil},
		{opSubtraction, -math.MaxFloat64, 2.0, 0.0, errOverflow},
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
			continue
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
