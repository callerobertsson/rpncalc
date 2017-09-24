// Package main is a simple command line calculator using RpnCalc
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/callerobertsson/rpncalc/rpncalc"
)

func main() {
	fmt.Println("RPN Calc")

	r := rpncalc.New()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	printPrompt(0.0, "enter q to quit")
	for scanner.Scan() {
		var err error
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "":
			continue
		case "q", "quit":
			os.Exit(0)
		case "?", "help":
			printHelp()
		case "rs", "regs", "registers":
			printRegisters(r)
		case "h", "history":
			printLog(r)
		default:
			err = r.Enter(input)
		}

		fmt.Printf("Stack: %v\n", r.Stack())
		msg := ""
		v, _ := r.Val()
		if err != nil {
			msg = err.Error()
		}

		printPrompt(v, msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func printPrompt(v float64, msg string) {
	if msg == "" {
		fmt.Printf("rpncalc %v > ", v)
		return
	}
	fmt.Printf("rpncalc %q > ", msg)
}

func printRegisters(r *rpncalc.RpnCalc) {
	fmt.Printf("Registers:\n")
	for i, r := range r.Regs() {
		fmt.Printf("  %2d: %v\n", i, r)
	}
}

func printLog(r *rpncalc.RpnCalc) {
	fmt.Printf("History (%d items)\n", len(r.Log()))
	for i, l := range r.Log() {
		fmt.Printf("  %4d: %v\n", len(r.Log())-i, l)
	}
}

func printHelp() {
	fmt.Println(`
	RPN Calc Help

	TODO: Write some more help info here!
	`)
}
