// Package main commands
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/callerobertsson/rpncalc/rpncalc"
)

type command struct {
	names       []string
	handler     func(*rpncalc.RpnCalc, ...string) error
	description string
}

var commands []command

func init() {
	commands = []command{
		{[]string{"q", "quit"}, cmdQuit, "Exits RpnCalc"},
		{[]string{"s", "stack"}, cmdStack, "Show stack values"},
		{[]string{"r", "regs"}, cmdRegs, "Show registers"},
		{[]string{"l", "history"}, cmdHistory, "Show calculation history"},
		{[]string{"h", "help"}, cmdHelp, "Show RpnCalc help"},
	}
}

func doCommand(r *rpncalc.RpnCalc, in string) error {
	in = strings.TrimSpace(strings.TrimPrefix(in, ":"))
	for _, cmd := range commands {
		if member(in, cmd.names...) {
			return cmd.handler(r, in)
		}
	}
	return fmt.Errorf("unknown command %q", in)
}

func cmdQuit(_ *rpncalc.RpnCalc, _ ...string) error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil // :-)
}

func cmdStack(r *rpncalc.RpnCalc, _ ...string) error {
	fmt.Printf("Stack:\n")
	for i := len(r.Stack()) - 1; i >= 0; i-- {
		fmt.Printf("%3d: %10v", i, formatVal(r.Stack()[i]))
		if i != 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Println("")
	return nil
}

func cmdRegs(r *rpncalc.RpnCalc, _ ...string) error {
	fmt.Printf("Registers:\n")
	for i, v := range r.Regs() {
		fmt.Printf("  %2d: %v\n", i, formatVal(v))
	}
	return nil
}

func cmdHistory(r *rpncalc.RpnCalc, _ ...string) error {
	fmt.Printf("History:\n")
	if len(r.Log()) < 1 {
		fmt.Println("  history is empty")
		return nil
	}
	for i, l := range r.Log() {
		fmt.Printf("  %4d: %v\n", len(r.Log())-i, l)
	}
	return nil
}

func cmdHelp(_ *rpncalc.RpnCalc, _ ...string) error {
	format := "  %20v: %v\n"
	cmds := fmt.Sprintf(format, "Command", "Desciption")
	for _, cmd := range commands {
		cmds += fmt.Sprintf(format, strings.Join(cmd.names, ", "), cmd.description)
	}

	ops := "TBD"

	fmt.Printf(`
RPN Calc Help

COMMANDS

Commands are prefixed with a colon, ex ":quit" and must be entered first on a line.

List of commands:

%v

OPERATORS AND VALUES

Operators and values can be entered one per line or as a sequence of tokens separated by space.

Caluclations are performed using Reverse Polish Notation (RPN).

Unary operators will act on the first element in the stack, binary on the first two elements.

TODO: Enter more help information...

List of operators:

%v

`, cmds, ops)

	return nil
}
