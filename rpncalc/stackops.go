// Package rpncalc stack operations
package rpncalc

func opClearStack(r *RpnCalc, _ string) error {
	r.ClearStack()
	return nil
}

func opSwap(r *RpnCalc, _ string) error {
	return r.stackSwap(0, 1)
}

func (r *RpnCalc) stackSwap(i, j int) error {
	if i < 0 || j < 0 || i >= len(r.stack) || j >= len(r.stack) {
		return errIndexOutOfRange
	}
	tmp := r.stack[i]
	r.stack[i] = r.stack[j]
	r.stack[j] = tmp

	return nil
}
