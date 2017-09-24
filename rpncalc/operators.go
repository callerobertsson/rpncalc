// Package rpncalc operators
package rpncalc

import (
	"fmt"
	"math"
)

func (r *RpnCalc) unaryOp(f func(float64) (float64, error)) error {
	v, err := f(r.stack[0])
	if err != nil {
		return err
	}
	r.stack[0] = v
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

func opSquare(x float64) (float64, error) {
	if x > math.Sqrt(math.MaxFloat64) {
		fmt.Printf("opSquare(%v) = overflow\n", x)
		return 0.0, errOverflow
	}
	fmt.Printf("opSquare(%v) = %v\n", x, x*x)
	return x * x, nil
}

func opAddition(x, y float64) (float64, error) {
	if x < 0 && y < 0 {
		// both negative
		if math.MaxFloat64+x < -y || math.MaxFloat64+y < -x {
			return 0.0, errOverflow
		}
	}
	if x > 0 && y > 0 {
		if math.MaxFloat64-x < y || math.MaxFloat64-y < x {
			return 0.0, errOverflow
		}
	}
	return x + y, nil
}

func opSubtraction(x, y float64) (float64, error) {
	return opAddition(x, -y)
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
