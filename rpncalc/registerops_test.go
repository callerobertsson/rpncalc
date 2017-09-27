package rpncalc

import "testing"

func TestRegParseAndStore(t *testing.T) {
	cases := []struct {
		name string
		cmd  string  // command to store in register
		reg  int     // register to fetch value from
		val  float64 // value to be stored
		exp  float64 // expected value in register
		err  error   // expected error
	}{
		{"register 0", "rs0", 0, 2.0, 2.0, nil},
		{"register 9", "rs9", 9, 2.0, 2.0, nil},
		{"empty", "", 0, 0, 0, errInvalidRegister},
		{"register too big", "rs10", 0, 0, 0, errInvalidRegister},
		{"register NaN", "rsapa", 0, 0, 0, errInvalidRegister},
	}

	for _, c := range cases {
		// Create new RpnCalc
		r := New()

		// Add value to be stored first in stack
		r.stack[0] = c.val

		// Execute store command
		err := dynOpRegStore(r, c.cmd)
		if err != c.err {
			t.Errorf("Case %v: Expected error %v, but got %v", c.name, c.err, err)
		}

		// Check for value in register
		got := r.regs[c.reg]
		if c.exp != got {
			t.Errorf("Case %v: Expected value %v in register %v, but got %v", c.name, c.exp, c.reg, got)
		}
	}
}

func TestRegParseAndRetrieve(t *testing.T) {
	cases := []struct {
		name string
		cmd  string  // command to store in register
		reg  int     // register to fetch value from
		exp  float64 // expected value in register
		err  error   // expected error
	}{
		{"register 0", "rs0", 0, 2.0, nil},
		{"register 9", "rs9", 9, 2.0, nil},
		{"empty", "", 0, 0, errInvalidRegister},
		{"register too big", "rs10", 0, 0, errInvalidRegister},
		{"register NaN", "rsapa", 0, 0, errInvalidRegister},
	}

	for _, c := range cases {
		// Create new RpnCalc
		r := New()

		// Store expected value in register
		r.regs[c.reg] = c.exp

		// Execute store command
		err := dynOpRegRestore(r, c.cmd)
		if err != c.err {
			t.Errorf("Case %v: Expected error %v, but got %v", c.name, c.err, err)
		}

		// Check for value in register
		got := r.stack[0]
		if c.exp != got {
			t.Errorf("Case %v: Expected value %v in register %v, but got %v", c.name, c.exp, c.reg, got)
		}
	}
}

func TestRegParseAndClear(t *testing.T) {
	cases := []struct {
		name string
		cmd  string  // command to store in register
		reg  int     // register to fetch value from
		val  float64 // value in register, will be cleared
		err  error   // expected error
	}{
		{"register 0", "rs0", 0, 2.0, nil},
		{"register 9", "rs9", 9, 2.0, nil},
		{"empty", "", 0, 3.0, errInvalidRegister},
		{"register too big", "rs10", 0, 4.0, errInvalidRegister},
		{"register NaN", "rsapa", 0, 5.0, errInvalidRegister},
	}

	for _, c := range cases {
		// Create new RpnCalc
		r := New()

		// Store expected value in register
		r.regs[c.reg] = c.val

		// Execute store command
		err := dynOpRegClear(r, c.cmd)
		if err != c.err {
			t.Errorf("Case %v: Expected error %v, but got %v", c.name, c.err, err)
		}

		// Check for value in register
		got := r.stack[0]
		if 0.0 != got {
			t.Errorf("Case %v: Expected value to be cleared in register %v, but got %v", c.name, c.reg, got)
		}
	}

}
