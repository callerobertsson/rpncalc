// Package rpncalc operators
package rpncalc

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
	{[]string{"sqrt", "root"}, opSquareRoot, "Calulates the square root"},
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

func opClearStack(r *RpnCalc) error {
	r.ClearStack()
	return nil
}

func opClearRegs(r *RpnCalc) error {
	r.ClearRegs()
	return nil
}
