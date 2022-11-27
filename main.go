// Package main is a simple command line calculator using RpnCalc
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/callerobertsson/rpn/rpncalc"
	"github.com/chzyer/readline"
)

var config = struct {
	DisplayPrecision int  `json:"prec"`
	ShowStack        bool `json:"showstack"`
}{
	DisplayPrecision: 2,
	ShowStack:        false,
}

func main() {
	r := rpncalc.New()

	if len(os.Args) > 1 {
		err := calculate(r, strings.Join(os.Args[1:], " "))
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("%f\n", r.Stack()[0])
		os.Exit(0)
	}

	fmt.Println("Simple RPN Calculator")
	fmt.Println(`enter "help" for help or "quit" to quit`)

	// Create input line reader
	rl, err := readline.New(prompt(r, ""))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create line reader:", err)
		os.Exit(1)
	}

	for {
		// Read input line
		line, err := rl.Readline()
		if err != nil {
			switch err {
			case io.EOF:
				line = ""
			case readline.ErrInterrupt:
				line = "quit"
			default:
				fmt.Fprintln(os.Stderr, "Reading input line failed:", err)
				os.Exit(1)
			}
		}

		err = calculate(r, line)

		// Add error message to prompt, if it exists
		msg := ""
		if err != nil {
			msg = err.Error()
		}

		rl.SetPrompt(prompt(r, msg))
	}
}

func calculate(r *rpncalc.RpnCalc, line string) (err error) {
	line = strings.TrimSpace(line)
	args := strings.Split(line, " ")
	args = filter(args, func(x string) bool { return x != "" })

	// Choose what to do
	switch {
	case len(args) < 1:
		return err
	case isCommand(args[0]):
		err = doCommand(r, args)
	default:
		err = r.Enter(line)
	}

	return err
}

func prompt(r *rpncalc.RpnCalc, msg string) (p string) {
	if config.ShowStack {
		cmdStack(r, []string{"s"}) // reuse stack command
	}
	p = fmt.Sprintf("%v", formatVal(r.Val()))

	if msg == "" {
		p += " > "
		return p
	}
	p += fmt.Sprintf(" [%v] > ", msg)

	return p
}

func formatVal(v float64) string {
	f := fmt.Sprintf("%%.%vf", config.DisplayPrecision)
	return fmt.Sprintf(f, v)
}

func jsonConfig() string {
	bs, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return "could not display config"
	}

	return string(bs)
}
