package rpncalc

import (
	"fmt"
	"testing"
)

func TestStackSwap(t *testing.T) {

	cases := []struct {
		i int
		j int
		s []float64
		e []float64
		x error
	}{
		{0, 1, []float64{1, 2}, []float64{2, 1}, nil},
		{1, 0, []float64{1, 2}, []float64{2, 1}, nil},
		{3, 0, []float64{1, 2}, []float64{1, 2}, errIndexOutOfRange},
		{0, -1, []float64{1, 2}, []float64{1, 2}, errIndexOutOfRange},
		{3, 5, []float64{1, 2, 3, 4, 5, 6}, []float64{1, 2, 3, 6, 5, 4}, nil},
	}

	r := New()

	for i, c := range cases {

		r.stack = c.s

		err := r.stackSwap(c.i, c.j)

		if err != c.x {
			t.Errorf("Case %v: Expected error %v, but got %v\n", i, c.x, err)
			continue
		}

		if fmt.Sprintf("%v", c.e) != fmt.Sprintf("%v", r.Stack()) {
			t.Errorf("Case %v: Expected stack %v, but got %v\n", i, c.s, r.Stack())
		}
	}
}
