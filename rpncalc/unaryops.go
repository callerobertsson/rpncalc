// Package rpncalc operators
package rpncalc

import "math"

// Binary operators
var unaryOps = []struct {
	names       []string
	handler     func(*RpnCalc) error
	description string
}{
	{[]string{"!", "neg"}, opNegate, "Negates (-x) first value on stack"},
	{[]string{"inv"}, opInverse, "Inverts (1/x) first value on stack"},
	{[]string{"sq", "square"}, opSquare, "Squares (x^2) first value on stack"},
	{[]string{"sqrt", "root"}, opSquareRoot, "Calulates the square root"},
}

func (r *RpnCalc) unaryOp(f func(float64) (float64, error)) error {
	v, err := f(r.stack[0])
	if err != nil {
		return err
	}
	r.stack[0] = v
	return nil
}

func opNegate(r *RpnCalc) error {
	return r.unaryOp(func(x float64) (float64, error) {
		return x * (-1.0), nil
	})
}

func opInverse(r *RpnCalc) error {
	return r.unaryOp(func(x float64) (float64, error) {
		if x == 0.0 {
			return 0.0, errDivisionByZero
		}

		return 1 / x, nil
	})
}

func opSquare(r *RpnCalc) error {
	return r.unaryOp(func(x float64) (float64, error) {
		if x > math.Sqrt(math.MaxFloat64) {
			return 0.0, errOverflow
		}
		return x * x, nil
	})
}

func opSquareRoot(r *RpnCalc) error {
	return r.unaryOp(func(x float64) (float64, error) {
		r := math.Sqrt(x)
		if math.IsNaN(r) {
			return 0.0, errNaN
		}
		return r, nil
	})
}
