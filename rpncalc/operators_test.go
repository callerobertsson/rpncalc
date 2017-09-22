package rpncalc

import "testing"

func TestUnaryOp(t *testing.T) {

	cases := []struct {
		f   func(float64) (float64, error)
		v   float64
		exp float64
		err error
	}{
		{opNegate, 1, -1, nil},
		{opNegate, -3.14, 3.14, nil},
		// TODO: Add more cases for unary operators
	}

	for i, c := range cases {
		got, err := c.f(c.v)

		if err != c.err {
			t.Errorf("Case %d: Expected error %q, but got %q", i, c.err, err)
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
		{opAddition, 123456789.12345, -1, 123456788.12345, nil},
		{opDivision, 101.0, 10.0, 10.1, nil},
		{opDivision, 101.0, 0.0, 0.0, errDivisionByZero},

		// TODO: Add more cases for unary operators
	}

	for i, c := range cases {
		got, err := c.f(c.x, c.y)

		if err != c.err {
			t.Errorf("Case %d: Expected error %q, but got %q", i, c.err, err)
		}

		if got != c.exp {
			t.Errorf("Case %d: Expected result %v, but got %v", i, c.exp, got)
		}
	}

}
