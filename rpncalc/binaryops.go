// Package rpncalc operators
package rpncalc

import "math"

func (r *RpnCalc) binaryOp(f func(float64, float64) (float64, error)) error {
	v, err := f(r.stack[1], r.stack[0])
	if err != nil {
		return err
	}

	r.stack = rolldown(r.stack)
	r.stack[0] = v
	return nil
}

func opAddition(r *RpnCalc, _ string) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		z := x + y
		if math.IsInf(z, 1) || math.IsInf(z, -1) {
			return 0.0, errOverflow
		}
		return z, nil
	})
}

func opSubtraction(r *RpnCalc, _ string) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		z := x - y
		if math.IsInf(z, 1) || math.IsInf(z, -1) {
			return 0.0, errOverflow
		}
		return z, nil
	})
}

func opMultiplication(r *RpnCalc, _ string) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		z := x * y
		if math.IsInf(z, 1) || math.IsInf(z, -1) {
			return 0.0, errOverflow
		}
		return z, nil
	})
}

func opDivision(r *RpnCalc, _ string) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		if y == 0.0 {
			return 0.0, errDivisionByZero
		}

		return x / y, nil
	})
}

func opPower(r *RpnCalc, _ string) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		return math.Pow(x, y), nil
	})
}
