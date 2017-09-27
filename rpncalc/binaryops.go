// Package rpncalc operators
package rpncalc

import "math"

// Static operators can have different names but no postfix values
var binaryOps = []struct {
	names       []string
	handler     func(*RpnCalc) error
	description string
}{
	{[]string{"!", "neg"}, opNegate, "Negates (-x) first value on stack"},
	{[]string{"inv"}, opInverse, "Inverts (1/x) first value on stack"},
	{[]string{"sq", "square"}, opSquare, "Squares (x^2) first value on stack"},
	{[]string{"sqrt", "root"}, opSquareRoot, "Calulates the square root"},
	{[]string{"+", "add"}, opAddition, "Adds (x+y) first two values on stack"},
	{[]string{"-", "sub"}, opSubtraction, "Subtracts (y-x) first two values on stack"},
	{[]string{"*", "mul"}, opMultiplication, "Multiplies (y*x) first two values on stack"},
	{[]string{"/", "div"}, opDivision, "Divides (y/x) first two values on stack"},
	{[]string{"cs", "clearstack"}, opClearStack, "Clears all values on stack"},
	{[]string{"cr", "clearregs"}, opClearRegs, "Clears all register values"},
	{[]string{"**", "pow"}, opPower, "Calculates y to the power of x (y**x)"},
	{[]string{"sw", "swap"}, opSwap, "Swap first and second value on stack"},
}

func (r *RpnCalc) binaryOp(f func(float64, float64) (float64, error)) error {
	v, err := f(r.stack[1], r.stack[0])
	if err != nil {
		return err
	}

	r.stack = rolldown(r.stack)
	r.stack[0] = v
	return nil
}

func opAddition(r *RpnCalc) error {
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

func opSubtraction(r *RpnCalc) error {
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

func opMultiplication(r *RpnCalc) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		return x * y, nil
	})
}

func opDivision(r *RpnCalc) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		if y == 0.0 {
			return 0.0, errDivisionByZero
		}

		return x / y, nil
	})
}

func opPower(r *RpnCalc) error {
	return r.binaryOp(func(x, y float64) (float64, error) {
		return math.Pow(x, y), nil
	})
}

func opSwap(r *RpnCalc) error {
	return r.stackSwap(0, 1)
}
