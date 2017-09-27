// Package rpncalc operators
package rpncalc

import "math"

func (r *RpnCalc) unaryOp(f func(float64, string) (float64, error)) error {
	v, err := f(r.stack[0], "")
	if err != nil {
		return err
	}
	r.stack[0] = v
	return nil
}

func opNegate(r *RpnCalc, _ string) error {
	return r.unaryOp(func(x float64, _ string) (float64, error) {
		return x * (-1.0), nil
	})
}

func opInverse(r *RpnCalc, _ string) error {
	return r.unaryOp(func(x float64, _ string) (float64, error) {
		if x == 0.0 {
			return 0.0, errDivisionByZero
		}

		return 1 / x, nil
	})
}

func opSquare(r *RpnCalc, _ string) error {
	return r.unaryOp(func(x float64, _ string) (float64, error) {
		if x > math.Sqrt(math.MaxFloat64) {
			return 0.0, errOverflow
		}
		return x * x, nil
	})
}

func opSquareRoot(r *RpnCalc, _ string) error {
	return r.unaryOp(func(x float64, _ string) (float64, error) {
		r := math.Sqrt(x)
		if math.IsNaN(r) {
			return 0.0, errNaN
		}
		return r, nil
	})
}
