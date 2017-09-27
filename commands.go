// Package main commands
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/callerobertsson/rpncalc/rpncalc"
)

type command struct {
	names       []string
	handler     func(*rpncalc.RpnCalc, []string) error
	description string
}

var commands []command

func init() {
	commands = []command{
		{[]string{"q", "quit"}, cmdQuit, "Exits RpnCalc"},
		{[]string{"s", "stack"}, cmdStack, "Show or clear stack values"},
		{[]string{"r", "regs"}, cmdRegs, "Show or clear registers"},
		{[]string{"l", "log"}, cmdLog, "Show calculation history"},
		{[]string{"set"}, cmdSetting, "Show or set configuration settings"},
		{[]string{"h", "help"}, cmdHelp, "Show RpnCalc help"},
	}
}

func doCommand(r *rpncalc.RpnCalc, in string) error {
	in = strings.TrimSpace(strings.TrimPrefix(in, ":"))
	args := strings.Split(in, " ")
	args = filter(args, func(x string) bool { return x != "" })
	for _, cmd := range commands {
		if member(args[0], cmd.names...) {
			return cmd.handler(r, args)
		}
	}
	return fmt.Errorf("unknown command %q", in)
}

func cmdQuit(_ *rpncalc.RpnCalc, _ []string) error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil // :-)
}

func cmdStack(r *rpncalc.RpnCalc, args []string) error {
	if len(args) == 1 {
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

	if len(args) == 2 && args[1] == "clear" {
		r.ClearStack()
		return nil
	}

	return fmt.Errorf("unknown command: %q", strings.Join(args, " "))
}

func cmdRegs(r *rpncalc.RpnCalc, args []string) error {
	if len(args) == 1 {
		fmt.Printf("Registers:\n")
		for i, v := range r.Regs() {
			fmt.Printf("  %2d: %v\n", i, formatVal(v))
		}
		return nil
	}

	if len(args) == 2 && args[1] == "clear" {
		r.ClearRegs()
		return nil
	}

	return fmt.Errorf("unknown command: %q", strings.Join(args, " "))
}

func cmdSetting(r *rpncalc.RpnCalc, args []string) error {
	if len(args) < 2 {
		// show all configuration
		fmt.Printf("%v\n", jsonConfig())
		return nil
	}

	if len(args) >= 2 {
		// show one configuaration
		f := "  %v: %v\n"
		switch args[1] {
		case "prec":
			if len(args) > 2 {
				p, err := strconv.Atoi(args[2])
				if err != nil {
					return fmt.Errorf("precision value is not a number")
				}
				config.DisplayPrecision = p
			}
			fmt.Printf(f, "prec", config.DisplayPrecision)
		case "showstack":
			if len(args) > 2 {
				t, err := strconv.ParseBool(args[2])
				if err != nil {
					return fmt.Errorf("%q is not a boolean value", args[2])
				}
				config.ShowStack = t
			}
			fmt.Printf(f, "showstack", config.ShowStack)
		default:
			return fmt.Errorf("unknown setting: %q", args[1])
		}

		return nil
	}

	return fmt.Errorf("partially implemented")
}

func cmdLog(r *rpncalc.RpnCalc, _ []string) error {
	fmt.Printf("Log:\n")
	if len(r.Log()) < 1 {
		fmt.Println("  log is empty")
		return nil
	}
	for i, l := range r.Log() {
		fmt.Printf("  %4d: %v\n", len(r.Log())-i, l)
	}
	return nil
}

func cmdHelp(r *rpncalc.RpnCalc, _ []string) error {
	format := "  %20v: %v\n"
	cmds := fmt.Sprintf(format, "Command", "Desciption")
	for _, cmd := range commands {
		cmds += fmt.Sprintf(format, strings.Join(cmd.names, ", "), cmd.description)
	}

	ops := ""
	for _, op := range rpncalc.OpsInfo() {
		ops += fmt.Sprintf(format, strings.Join(op.Names, ", "), op.Description)
	}

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
