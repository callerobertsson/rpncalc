// Package rpncalc operators. Operators modifies the stack or the registers.
package rpncalc

// OperatorType defines an operator to be static (exact match) or dynamic (postfixed with a value)
type OperatorType int

const (
	// StaticOp defines the static operation type
	StaticOp = iota
	// DynamicOp defines the dynamic operation type
	DynamicOp
)

// OpInfo contains displayable operation information
type OpInfo struct {
	Type        OperatorType
	Names       []string
	Prefix      string
	Description string
}

// Operator defines data needed for one operator
type Operator struct {
	Type        OperatorType
	Names       []string // used by static ops
	Prefix      string   // used by dynamic ops
	Handler     func(*RpnCalc, string) error
	Description string
}

// OpsInfo returns an OpInfo slice with all supported static operators
func OpsInfo() []OpInfo {
	ois := []OpInfo{}
	for _, o := range operators {
		ois = append(ois, OpInfo{o.Type, o.Names, o.Prefix, o.Description})
	}
	return ois
}

var operators = []Operator{
	// Unary
	{StaticOp, []string{"neg"}, "", opNegate, "Negates (-x) first value on stack"},
	{StaticOp, []string{"inv"}, "", opInverse, "Inverts (1/x) first value on stack"},
	{StaticOp, []string{"sq", "square"}, "", opSquare, "Squares (x^2) first value on stack"},
	{StaticOp, []string{"sqrt", "root"}, "", opSquareRoot, "Calulates the square root"},
	{StaticOp, []string{"bin", "b"}, "", opDecToBin, "Converts decimal to binary"},
	{StaticOp, []string{"dec", "d"}, "", opBinToDec, "Converts binary to decimal"},
	// Binary
	{StaticOp, []string{"+", "add"}, "", opAddition, "Adds (x+y) first two values on stack"},
	{StaticOp, []string{"-", "sub"}, "", opSubtraction, "Subtracts (y-x) first two values on stack"},
	{StaticOp, []string{"*", "mul"}, "", opMultiplication, "Multiplies (y*x) first two values on stack"},
	{StaticOp, []string{"/", "div"}, "", opDivision, "Divides (y/x) first two values on stack"},
	{StaticOp, []string{"**", "pow"}, "", opPower, "Calculates y to the power of x (y**x)"},
	{StaticOp, []string{"%", "mod"}, "", opModulus, "Calculates x modulus y"},
	// Stack
	{StaticOp, []string{"sw", "swap"}, "", opSwap, "Swap pos 0 and pos 1 on the stack"},
	// Register
	{DynamicOp, []string{}, "rs", dynOpRegStore, "Store (rsX) value in register X"},
	{DynamicOp, []string{}, "rr", dynOpRegRestore, "Restore (rrX) value from register X"},
	{DynamicOp, []string{}, "rc", dynOpRegClear, "Clear (rcX) value from register X"},

	// TODO: Add more operators
}
