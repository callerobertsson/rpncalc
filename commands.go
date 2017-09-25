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
		{[]string{"h", "help"}, cmdHelp, "Show RpnCalc help"},
		{[]string{"s", "stack"}, cmdStack, "Show stack values"},
		{[]string{"r", "regs"}, cmdRegs, "Show registers"},
		{[]string{"l", "history"}, cmdHistory, "Show calculation history"},
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

func cmdHelp(_ *rpncalc.RpnCalc, _ ...string) error {
	printHelp()
	return nil
}

func cmdStack(r *rpncalc.RpnCalc, _ ...string) error {
	fmt.Printf("Stack:\n")
	printStack(r)
	fmt.Println("")
	return nil
}

func cmdRegs(r *rpncalc.RpnCalc, _ ...string) error {
	fmt.Printf("Registers:\n")
	printRegisters(r)
	return nil
}

func cmdHistory(r *rpncalc.RpnCalc, _ ...string) error {
	fmt.Printf("History:\n")
	printLog(r)
	return nil
}
