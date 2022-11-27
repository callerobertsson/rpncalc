// Package rpncalc operators
package rpncalc

var constants = map[string]float64{
	// Math
	"e":   2.718281828459,         // base natural logarithm
	"phi": 1.61803398874989484820, // golden ratio
	"pi":  3.1415926535897932,     //
	"tau": 6.28318530717958623200, // pi * 2
	// Units
	"k":  1000,    // kilo
	"kb": 1024,    // kilo
	"M":  1000000, // mega
	// Physics
	"sol": 299792458, // in vacuum

	// TODO: add constants
}

// Constants returns a string float map of constant name and value.
func Constants() map[string]float64 {
	return constants
}

func (r *RpnCalc) pushConstant(name string) bool {

	if val, ok := constants[name]; ok {
		r.stack = enter(r.stack, val)
		return true
	}

	return false
}
