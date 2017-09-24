// Package rpncalc stack operations
package rpncalc

func (r *RpnCalc) stackSwap(i, j int) error {
	if i < 0 || j < 0 || i >= len(r.stack) || j >= len(r.stack) {
		return errIndexOutOfRange
	}
	tmp := r.stack[i]
	r.stack[i] = r.stack[j]
	r.stack[j] = tmp

	return nil
}
