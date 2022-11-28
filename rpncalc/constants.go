// Package rpncalc operators
package rpncalc

import "math"

// Constant type storing names, value and description of a constant
type Constant struct {
	Names       []string
	Value       float64
	Description string
}

var constants = []Constant{
	// Math
	{[]string{"e"}, 2.718281828459, "natural logarithm base"},
	{[]string{"phi"}, 1.61803398874989484820, "golden ratio"},
	{[]string{"pi"}, 3.1415926535897932, "pi"},
	{[]string{"tau"}, 6.28318530717958623200, "2 *pi"},
	// Units
	{[]string{"k", "kilo"}, 1000, "kilo"},
	{[]string{"M", "mega"}, 1000000, "mega"},
	{[]string{"G", "giga"}, 1000000000, "tera"},
	{[]string{"T", "tera"}, 1000000000000, "giga"},
	{[]string{"kb", "kilobyte"}, 1024, "kilo byte"},
	{[]string{"Mb", "megabyte"}, 1048576, "mega byte"},
	{[]string{"Gb", "gigabyte"}, 1073741824, "tera byte"},
	{[]string{"Tb", "terabyte"}, 1099511627776, "giga byte"},
	// Physics
	{[]string{"sol"}, 299792458, "m/s speed of light in vacuum"},
	// Maxima
	{[]string{"maxf"}, math.MaxFloat64, "maximum size of values in rpn"},

	// TODO: add constants
}

// Constants returns a string float map of constant name and value.
func Constants() []Constant {
	return constants
}

func (r *RpnCalc) pushConstant(name string) bool {

	for _, c := range constants {
		for _, n := range c.Names {
			if n == name {
				r.stack = enter(r.stack, c.Value)
				return true
			}
		}
	}

	return false
}
