// Package rpncalc operators
package rpncalc

import (
	"math"
)

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

func opDecToBin(r *RpnCalc, _ string) error {
	return r.unaryOp(func(x float64, _ string) (float64, error) {
		dec := int64(x)
		bin := int64(0)
		i := 0

		for dec > 0 {
			lastBin := bin
			bin += int64(math.Pow10(i)) * (dec % int64(2))
			dec = dec / 2
			i++
			if lastBin > bin {
				return 0.0, errOverflow
			}
		}

		return float64(bin), nil
	})
}
