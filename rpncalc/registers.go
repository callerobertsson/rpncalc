// Package rpncalc register functions
package rpncalc

import "strconv"

func dynOpRegStore(r *RpnCalc, t string) error {
	reg, err := parseReg(t)
	if err != nil {
		return errInvalidRegister
	}

	return r.regStore(reg)
}

func dynOpRegRestore(r *RpnCalc, t string) error {
	reg, err := parseReg(t)
	if err != nil {
		return errInvalidRegister
	}

	return r.regRetrieve(reg)
}

func dynOpRegClear(r *RpnCalc, t string) error {
	reg, err := parseReg(t)
	if err != nil {
		return errInvalidRegister
	}

	return r.ClearReg(reg)
}

func parseReg(t string) (int, error) {
	if len(t) < 2 {
		return 0, errInvalidRegister
	}
	t = t[2:]
	if val, err := strconv.Atoi(t); err == nil {
		return val, nil
	}

	return 0, errInvalidRegister
}

func (r *RpnCalc) regStore(i int) error {
	if i < 0 || i >= len(r.regs) {
		return errInvalidRegister
	}
	r.regs[i] = r.stack[0]

	return nil
}

func (r *RpnCalc) regRetrieve(i int) error {
	if i < 0 || i >= len(r.regs) {
		return errInvalidRegister
	}
	enter(r.stack, r.regs[i])
	return nil
}
