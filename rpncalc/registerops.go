// Package rpncalc register functions
package rpncalc

import "strconv"

func opClearRegs(r *RpnCalc, _ string) error {
	r.ClearRegs()
	return nil
}

func dynOpRegStore(r *RpnCalc, t string) error {
	reg, err := parseReg(t)
	if err != nil {
		return errInvalidRegister
	}

	if reg < 0 || reg >= len(r.regs) {
		return errInvalidRegister
	}
	r.regs[reg] = r.stack[0]

	return nil
}

func dynOpRegRestore(r *RpnCalc, t string) error {
	reg, err := parseReg(t)
	if err != nil {
		return errInvalidRegister
	}

	if reg < 0 || reg >= len(r.regs) {
		return errInvalidRegister
	}

	enter(r.stack, r.regs[reg])
	return nil
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
