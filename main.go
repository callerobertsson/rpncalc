// Package main is a simple command line calculator using RpnCalc
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/callerobertsson/rpncalc/rpncalc"
)

var config = struct {
	DisplayPrecision int
	PrintStack       bool
}{
	DisplayPrecision: 2,
	PrintStack:       false,
}

func main() {
	fmt.Println("Simple RPN Calculator")
	fmt.Println(`enter ":h" for help or ":q" to quit`)

	r := rpncalc.New()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	printPrompt(r, "")
	for scanner.Scan() {
		var err error
		var msg string

		input := strings.TrimSpace(scanner.Text())

		switch {
		case input == "":
			continue
		case strings.HasPrefix(input, ":"):
			err = doCommand(r, input)
		default:
			err = r.Enter(input)
		}

		if err != nil {
			msg = err.Error()
		}

		printPrompt(r, msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func printPrompt(r *rpncalc.RpnCalc, msg string) {
	if config.PrintStack {
		fmt.Println("==================")
		printStack(r)
	} else {
		fmt.Printf("%v", formatVal(r.Val()))
	}

	if msg == "" {
		fmt.Printf(" > ")
		return
	}
	fmt.Printf(" [%v] > ", msg)
}

func printStack(r *rpncalc.RpnCalc) {
	for i := len(r.Stack()) - 1; i >= 0; i-- {
		fmt.Printf("%3d: %10v", i, formatVal(r.Stack()[i]))
		if i != 0 {
			fmt.Printf("\n")
		}
	}
}

func printRegisters(r *rpncalc.RpnCalc) {
	for i, v := range r.Regs() {
		fmt.Printf("  %2d: %v\n", i, formatVal(v))
	}
}

func printLog(r *rpncalc.RpnCalc) {
	if len(r.Log()) < 1 {
		fmt.Println("  history is empty")
		return
	}
	for i, l := range r.Log() {
		fmt.Printf("  %4d: %v\n", len(r.Log())-i, l)
	}
}

func formatVal(v float64) string {
	f := fmt.Sprintf("%%.%vf", config.DisplayPrecision)
	return fmt.Sprintf(f, v)
}

func printHelp() {
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

}

func member(t string, ms ...string) bool {
	for _, m := range ms {
		if m == t {
			return true
		}
	}
	return false
}
