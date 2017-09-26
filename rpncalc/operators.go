// Package rpncalc operators
package rpncalc

import "math"

// OpInfo contains displayable operation information
type OpInfo struct {
	Names       []string
	Description string
}

// StaticOpsInfo returns an OpInfo slice with all supported static operators
func StaticOpsInfo() []OpInfo {
	ois := []OpInfo{}
	for _, o := range staticOps {
		ois = append(ois, OpInfo{o.names, o.description})
	}
	return ois
}

// DynamicOpsInfo returns an OpInfo slice with all supported dynamic operators
func DynamicOpsInfo() []OpInfo {
	ois := []OpInfo{}
	for _, o := range dynamicOps {
		ois = append(ois, OpInfo{[]string{o.prefix + "X"}, o.description})
	}
	return ois
}

// Static operators can have different names but no postfix values
var staticOps = []struct {
	names       []string
	handler     func(*RpnCalc) error
	description string
}{
	{[]string{"!", "neg"}, opNegate, "Negates (-x) first value on stack"},
	{[]string{"inv"}, opInverse, "Inverts (1/x) first value on stack"},
	{[]string{"sq", "square"}, opSquare, "Squares (x^2) first value on stack"},
	{[]string{"+", "add"}, opAddition, "Adds (x+y) first two values on stack"},
	{[]string{"-", "sub"}, opSubtraction, "Subtracts (y-x) first two values on stack"},
	{[]string{"*", "mul"}, opMultiplication, "Multiplies (y*x) first two values on stack"},
	{[]string{"/", "div"}, opDivision, "Divides (y/x) first two values on stack"},
	{[]string{"cs", "clearstack"}, opClearStack, "Clears all values on stack"},
	{[]string{"cr", "clearregs"}, opClearRegs, "Clears all register values"},
	{[]string{"sw", "swap"}, opSwap, "Swap first and second value on stack"},
}

// Dynamic operators is a prefix string with a postfix number
var dynamicOps = []struct {
	prefix      string
	handler     func(*RpnCalc, string) error
	description string
}{
	{"rs", dynOpRegStore, "Store (rsX) value in register X"},
	{"rr", dynOpRegRestore, "Restore (rrX) value from register X"},
	{"rc", dynOpRegClear, "Clear (rcX) value from register X"},
}

func (r *RpnCalc) unaryOp(f func(float64) (float64, error)) error {
	v, err := f(r.stack[0])
	if err != nil {
		return err
	}
	r.stack[0] = v
	return nil
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

func opClearStack(r *RpnCalc) error {
	r.ClearStack()
	return nil
}

func opClearRegs(r *RpnCalc) error {
	r.ClearRegs()
	return nil
}

func opSwap(r *RpnCalc) error {
	return r.stackSwap(0, 1)
}
