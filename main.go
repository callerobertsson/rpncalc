// Package main is a simple command line calculator using RpnCalc
package main

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("Simple RPN Calculator")
	fmt.Println(`enter "help" for help or "quit" to quit`)

	r := rpncalc.New()

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
			fmt.Fprintln(os.Stderr, "Reading input line failed:", err)
			os.Exit(1)
		}
		line = strings.TrimSpace(line)
		args := strings.Split(line, " ")
		args = filter(args, func(x string) bool { return x != "" })

		// Choose what to do
		switch {
		case len(args) < 1:
			continue
		case isCommand(args[0]):
			err = doCommand(r, args)
		default:
			err = r.Enter(line)
		}

		// Add error message to prompt, if it exists
		msg := ""
		if err != nil {
			msg = err.Error()
		}

		rl.SetPrompt(prompt(r, msg))
	}
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
