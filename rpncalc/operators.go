// Package rpncalc operators
package rpncalc

import "fmt"

func (r *RpnCalc) unaryOp(f func(float64) (float64, error)) error {
	v, err := f(r.stack[0])
	if err != nil {
		return err
	}
	r.stack[0] = v
	r.log = append(r.log, fmt.Sprintf(">> %v", v))

	return nil
}

func (r *RpnCalc) binaryOp(f func(float64, float64) (float64, error)) error {
	s1 := r.stack[0]
	s2 := r.stack[1]
	v, err := f(s2, s1)
	if err != nil {
		return err
	}

	r.stack = rolldown(r.stack)
	r.stack[0] = v
	r.log = append(r.log, fmt.Sprintf(">> %v", v))

	return nil
}

func opNegate(x float64) (float64, error) {
	return x * (-1.0), nil
}

func opInverse(x float64) (float64, error) {
	if x == 0.0 {
		return 0.0, errDivisionByZero
	}

	return 1 / x, nil
}

func opAddition(x, y float64) (float64, error) {
	// TODO: Check for MaxFloat64?
	return x + y, nil
}

func opSubtraction(x, y float64) (float64, error) {
	return x - y, nil
}

func opMultiplication(x, y float64) (float64, error) {
	return x * y, nil
}

func opDivision(x, y float64) (float64, error) {
	if y == 0.0 {
		return 0.0, errDivisionByZero
	}

	return x / y, nil
}
