package rpncalc

import (
	"fmt"
	"math"
	"testing"
)

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
		name  string
		input []string
		stack []float64
		err   error
	}{
		// Unary stuff
		{"negate value",
			[]string{"1 2 3 4 neg"}, []float64{-4, 3, 2, 1}, nil},
		{"inverse value",
			[]string{"1 2 3 4 inv"}, []float64{0.25, 3, 2, 1}, nil},
		{"sqare a number",
			[]string{"5 sq"},
			[]float64{25, 0, 0, 0}, nil},

		// Binary stuff
		{"enter 2 vals",
			[]string{"123", "234"}, []float64{234.0, 123.0, 0.0, 0.0}, nil},
		{"divide",
			[]string{"123", "10", "/"}, []float64{12.3, 0.0, 0.0, 0.0}, nil},
		{"add",
			[]string{"123", "10", "+"}, []float64{133.0, 0.0, 0.0, 0.0}, nil},
		{"double divide",
			[]string{"1234", "100", "10", "/", "/"}, []float64{123.4, 0.0, 0.0, 0.0}, nil},
		{"single string double divide",
			[]string{"1234 100 10 / /"}, []float64{123.4, 0.0, 0.0, 0.0}, nil},
		{"multiple ops",
			[]string{"1234 100 * 1000 / 1 -"}, []float64{122.4, 0.0, 0.0, 0.0}, nil},
		{"power of 1073",
			[]string{"1073 2 **"}, []float64{math.Pow(1073, 2), 0.0, 0.0, 0.0}, nil},
		{"swap value",
			[]string{"1 2 3 4 sw"}, []float64{3, 4, 2, 1}, nil},

		// Input sanity
		{"ignore spaces",
			[]string{"   1  2  3  4 "}, []float64{4, 3, 2, 1}, nil},

		// Regs stuff
		{"regs store don't change stack",
			[]string{"1 2 3 4 rs5"}, []float64{4, 3, 2, 1}, nil},
		{"regs clear don't change stack",
			[]string{"1 2 3 4 rc1"}, []float64{4, 3, 2, 1}, nil},
		{"regs retrive change stack",
			[]string{"1 2 3 4 rr9"}, []float64{0, 4, 3, 2}, nil},
		{"regs store invalid reg fails",
			[]string{"1 2 3 4 rs99"}, []float64{4, 3, 2, 1}, errInvalidRegister},
		{"regs clear invalid reg fails",
			[]string{"1 2 3 4 rcx"}, []float64{4, 3, 2, 1}, errInvalidRegister},
		{"regs retrive invalid reg fails",
			[]string{"1 2 3 4 rrapa"}, []float64{4, 3, 2, 1}, errInvalidRegister},

		// Unknown operations
		{"unknown op",
			[]string{"123", "foo"}, []float64{123.0, 0.0, 0.0, 0.0}, errUnknownInput},
		{"fail in middle",
			[]string{"123 foo 321"}, []float64{123.0, 0.0, 0.0, 0.0}, errUnknownInput},

		// TODO: Add testcases when new functionality comes along
	}

	for _, c := range cases {
		r := New()

		var err error

		// Enter all input data
		for i := 0; i < len(c.input)-1; i++ {
			err = r.Enter(c.input[i])
			if err != nil {
				t.Errorf("%q: Error in step %d: %v", c.name, i, err)
				break
			}
		}

		if err != nil {
			continue
		}

		// Enter final operation
		err = r.Enter(c.input[len(c.input)-1])
		if err != c.err {
			t.Errorf("%q: Expected error %v, but got %v for %v", c.name, c.err, err, c.input[len(c.input)-1])
			continue
		}

		// Check stack
		for i := range c.stack {
			if c.stack[i] != r.stack[i] {
				t.Errorf("%q: Expected stack %v, but got %v", c.name, c.stack, r.stack)
			}
		}
	}
}

func TestValAndClear(t *testing.T) {
	expVal := 4.0

	r := New()

	err := r.Enter("1 2 3 4")
	if err != nil {
		t.Fatalf("Could not enter expression, got error %v", err)
	}

	val := r.Val()
	if val != expVal {
		t.Fatalf("Expected value %v, but got %v", expVal, val)
	}

	r.ClearVal()

	val = r.Val()
	if val != 0.0 {
		t.Fatalf("Expected value to be cleared, but got %v", val)
	}
}

func TestLogContentAndClear(t *testing.T) {

	expLog := []string{"1", "2", "+", ">> 3", "3", "+", ">> 6", "0", "inv", "[division by zero]"}

	r := New()

	err := r.Enter("1 2 + 3 + 0 inv")
	if err != errDivisionByZero {
		t.Fatalf("Expected error %v, got error %v", errDivisionByZero, err)
	}

	l := r.Log()
	if fmt.Sprintf("%v", l) != fmt.Sprintf("%v", expLog) {
		t.Fatalf("Expected log %v, but got %v", expLog, l)
	}

	r.ClearLog()

	if len(r.Log()) > 0 {
		t.Fatalf("Expected empty log, but it has length %v and contains %v", len(r.Log()), r.Log())
	}
}

func TestRegsAndClear(t *testing.T) {

	r := New()

	err := r.Enter("1 rs1 2 rs2")
	if err != nil {
		t.Fatalf("Faild to set valid registers, got error %v", err)
	}

	regs := r.Regs()
	if regs[1] != 1 || regs[2] != 2 {
		t.Fatalf("Bad register values expeted r1 = 1 and r2 = 2, but got %v", regs)
	}

	// Clear register 2
	err = r.ClearReg(2)
	if regs[2] != 0.0 {
		t.Errorf("Expected register 2 to be cleared, but it contained %v", regs[2])
	}

	// Try to clear invalid register should fail
	err = r.ClearReg(11)
	if err != errInvalidRegister {
		t.Errorf("Expected %v, but got %v when clearing invalid register #11", errInvalidRegister, err)
	}

	// Try to clear another invalid register should fail
	err = r.ClearReg(-1)
	if err != errInvalidRegister {
		t.Errorf("Expected %v, but got %v when clearing invalid register #-1", errInvalidRegister, err)
	}

	// Clear all registers
	r.ClearRegs()
	for i, r := range r.Regs() {
		if r != 0.0 {
			t.Errorf("Reg %v contains %v, expected it to be cleared", i, r)
		}
	}
}

func TestInvalidRegisters(t *testing.T) {

	r := New()

	err := r.Enter("1 rs10")
	if err == nil {
		t.Fatalf("Could set register 10, that should not be allowed")
	}

	err = r.Enter("rsapa")
	if err == nil {
		t.Fatalf("Could set register apa, that should not be allowed")
	}

	err = r.Enter("rs")
	if err == nil {
		t.Fatalf("Could set register with empty number, that should not be allowed")
	}

}
