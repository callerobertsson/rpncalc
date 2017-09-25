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
	Prec int
}{
	Prec: 2,
}

func main() {
	fmt.Println("Simple RPN Calculator")
	fmt.Println(`enter "h" for help or "q" to quit`)

	r := rpncalc.New()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	printPrompt(r, "")
	for scanner.Scan() {
		var err error
		var msg string

		// TODO: refactor in func, present in both main and rpncalc
		in := func(t string, ms ...string) bool {
			for _, m := range ms {
				if m == t {
					return true
				}
			}
			return false
		}

		input := strings.TrimSpace(scanner.Text())

		switch {
		case input == "":
			continue
		case in(input, "q", "quit"):
			os.Exit(0)
		case in(input, "?", "help"):
			printHelp()
		case in(input, "st", "stack"):
			fmt.Printf("Stack:\n")
			printStack(r)
			fmt.Println("")
		case in(input, "regs", "registers"):
			fmt.Printf("Registers:\n")
			printRegisters(r)
		case in(input, "h", "log", "history"):
			fmt.Printf("History:\n")
			printLog(r)
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
	fmt.Println("==================")
	printStack(r)
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
	f := fmt.Sprintf("%%.%vf", config.Prec)
	return fmt.Sprintf(f, v)
}

func printHelp() {
	fmt.Println(`
	RPN Calc Help

	TODO: Write some more help info here!
	`)
}
