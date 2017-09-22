package rpncalc

import "testing"

/*
	Enter(string) error
	Val() (float64, error)
	Stack() []float64
	Regs() []float64
	Log() []string
	ClearVal()
	ClearStack()
	ClearReg(i int) error
	ClearRegs()
	ClearLog()
*/

func TestNew(t *testing.T) {
	r := New()

	if r == nil {
		t.Errorf("Expected to get a RpnCalc but it was %v", r)
	}

	if s := len(r.Stack()); s != newStackSize {
		t.Errorf("Expected stack size %d, but got %d", newStackSize, s)
	}

	if s := len(r.Regs()); s != newRegsSize {
		t.Errorf("Expected regs size %d, but got %d", newRegsSize, s)
	}

	if s := len(r.Log()); s != newLogSize {
		t.Errorf("Expected log size %d, but got %d", newLogSize, s)
	}

	for i := range r.stack {
		if r.stack[i] != 0.0 {
			t.Errorf("Expected stack to contain zero values, but found %v", r.stack[i])
		}
	}

	for i := range r.regs {
		if r.regs[i] != 0.0 {
			t.Errorf("Expected regs to contain zero values, but found %v", r.regs[i])
		}
	}

}

func TestEnterVal(t *testing.T) {
	cases := []struct {
		input []string
		stack []float64
		err   error
	}{
		{[]string{"123", "234"}, []float64{234.0, 123.0, 0.0, 0.0}, nil},
		{[]string{"123", "foo"}, []float64{123.0, 0.0, 0.0, 0.0}, errUnknownOperation},
		{[]string{"123", "10", "/"}, []float64{12.3, 0.0, 0.0, 0.0}, nil},
		{[]string{"123", "10", "+"}, []float64{133.0, 0.0, 0.0, 0.0}, nil},
		{[]string{"1234", "100", "10", "/", "/"}, []float64{123.4, 0.0, 0.0, 0.0}, nil},

		// TODO: Add testcases when new functionality comes along

	}

	for ci, c := range cases {
		r := New()

		var err error

		// Enter all input data
		for i := 0; i < len(c.input)-1; i++ {
			err = r.Enter(c.input[i])
			if err != nil {
				t.Errorf("Case %d: Error in step %d: %v", ci, i, err)
				break
			}
		}

		if err != nil {
			continue
		}

		// Enter final operation
		err = r.Enter(c.input[len(c.input)-1])
		if err != c.err {
			t.Errorf("Case %d: Expected error %v, but got %q for %v", ci, c.err, err, c.input[len(c.input)-1])
			continue
		}

		// Check stack
		for i := range c.stack {
			if c.stack[i] != r.stack[i] {
				t.Errorf("Case %d: Expected stack %v, but got %v", ci, c.stack, r.stack)
			}
		}
	}
}
