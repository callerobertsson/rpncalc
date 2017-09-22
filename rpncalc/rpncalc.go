// Package rpncalc implements a RPN calculator
package rpncalc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// RpnCalcer defines the interface for a RpnCalc
type RpnCalcer interface {
	Enter(string) error
	Val() (float64, error)
	Stack() []float64
	Regs() []float64
	Log() []string
	ClearVal()
	ClearStack()
	ClearReg(i int) error
	ClearRegs()
	ClearLog()
}

const (
	newStackSize = 4
	newRegsSize  = 10
	newLogSize   = 0
)

var (
	errDivisionByZero   = errors.New("division by zero")
	errUnknownOperation = errors.New("unknown operation")
)

// RpnCalc implements a RPN calculator adhering to the RpnCalcer interface
type RpnCalc struct {
	stack []float64
	regs  []float64
	log   []string
	err   error
}

// New creates a new RpnCalc with default settings
func New() *RpnCalc {
	r := &RpnCalc{}

	r.stack = make([]float64, newStackSize)
	r.regs = make([]float64, newRegsSize)
	r.log = []string{}

	return r
}

// Enter takes some input, number, operator, or command, and tries to parse it
func (r *RpnCalc) Enter(expr string) error {
	expr = strings.ToLower(strings.TrimSpace(expr))

	// Try to parse a float64
	val, err := strconv.ParseFloat(expr, 64)
	if err == nil {
		r.log = append(r.log, fmt.Sprintf("%v", val))
		r.stack = enter(r.stack, val)
		return nil
	}
	err = nil

	// Try to find a matching operator
	switch expr {
	case "!", "neg":
		err = r.unaryOp(opNegate)
	case "inv":
		err = r.unaryOp(opInverse)
	case "+", "add":
		err = r.binaryOp(opAddition)
	case "-", "sub":
		err = r.binaryOp(opSubtraction)
	case "*", "mul", "mult":
		err = r.binaryOp(opMultiplication)
	case "/", "div":
		err = r.binaryOp(opDivision)
	default:
		r.log = append(r.log, "Unknown operation: "+expr)
		return errUnknownOperation
	}

	r.log = append(r.log, expr)
	if err != nil {
		r.log = append(r.log, "Error: "+err.Error())
	}

	return err
}

// Val gets the first value on the stack, the display value
func (r *RpnCalc) Val() (float64, error) {
	if r.err != nil {
		return 0.0, r.err
	}
	return r.stack[0], nil
}

// Stack returns the current stack of values
func (r *RpnCalc) Stack() []float64 {
	return r.stack
}

// Regs returns the registers
func (r *RpnCalc) Regs() []float64 {
	return r.regs
}

// Log returns the calculation log
func (r *RpnCalc) Log() []string {
	return r.log
}

// ClearVal puts a zero value in the first position of the stack
func (r *RpnCalc) ClearVal() {
	r.stack[0] = 0.0
}

// ClearStack puts zero values in all positons of the stack
func (r *RpnCalc) ClearStack() {
	for i := range r.stack {
		r.stack[i] = 0.0
	}
}

// ClearReg puts a zero value in the i:th position of the registers
func (r *RpnCalc) ClearReg(i int) error {
	if i < 0 || i > len(r.regs)-1 {
		return fmt.Errorf("index out of range")
	}
	r.regs[i] = 0.0
	return nil
}

// ClearRegs clears all register values
func (r *RpnCalc) ClearRegs() {
	for i := range r.regs {
		r.regs[i] = 0.0
	}
}

// ClearLog clears the log
func (r *RpnCalc) ClearLog() {
	r.log = []string{}
}

// Helper functions

func enter(s []float64, v float64) []float64 {
	s = rollup(s)
	s[0] = v
	return s
}

func rollup(s []float64) []float64 {
	for i := len(s) - 1; i > 0; i-- {
		s[i] = s[i-1]
	}
	return s
}

func rolldown(s []float64) []float64 {
	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	return s
}
