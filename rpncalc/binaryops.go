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
		// both negative
		if x < 0 && y < 0 {
			if math.MaxFloat64+x < -y || math.MaxFloat64+y < -x {
				return 0.0, errOverflow
			}
		}
		// both positive
		if x > 0 && y > 0 {
			if math.MaxFloat64-x < y || math.MaxFloat64-y < x {
				return 0.0, errOverflow
			}
		}
		return x + y, nil
	})
}

func opSubtraction(r *RpnCalc, _ string) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		if x < 0 && y > 0 {
			if math.MaxFloat64-x < y || math.MaxFloat64-y > x {
				return 0.0, errOverflow
			}
		}
		if x > 0 && y < 0 {
			if math.MaxFloat64-x < y || math.MaxFloat64-y < x {
				return 0.0, errOverflow
			}
		}
		return x - y, nil
	})
}

func opMultiplication(r *RpnCalc, _ string) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		return x * y, nil
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
