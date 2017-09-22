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
		words := strings.TrimSpace(scanner.Text())

		for _, word := range strings.Split(words, " ") {
			switch word {
			case "":
				continue
			case "q", "quit":
				os.Exit(0)
			case "?", "help":
				printHelp()
			case "h", "history":
				printLog(r)
			case "cs":
				r.ClearStack()
			case "cr":
				r.ClearRegs()
			default:
				err = r.Enter(word)
			}
		}

		fmt.Printf("Regs:  %v\n", r.Regs())
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
		fmt.Printf("rpncalc %f > ", v)
		return
	}
	fmt.Printf("rpncalc %q > ", msg)
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
