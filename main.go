// Package main is a simple command line calculator using RpnCalc
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/callerobertsson/rpncalc/rpncalc"

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
	fmt.Println(`enter ":h" for help or ":q" to quit`)

	r := rpncalc.New()

	rl, err := readline.New(prompt(r, ""))
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create line reader:", err)
		os.Exit(1)
	}

	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Fprintln(os.Stderr, "reading input failed:", err)
			os.Exit(1)
		}

		input := strings.TrimSpace(line)

		switch {
		case input == "":
			continue
		case strings.HasPrefix(input, ":"):
			err = doCommand(r, input)
		default:
			err = r.Enter(input)
		}

		msg := ""
		if err != nil {
			msg = err.Error()
		}

		rl.SetPrompt(prompt(r, msg))
	}

}

func prompt(r *rpncalc.RpnCalc, msg string) (p string) {
	if config.ShowStack {
		cmdStack(r, []string{})
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

// TODO: Move member func to util package?
func member(t string, ms ...string) bool {
	for _, m := range ms {
		if m == t {
			return true
		}
	}
	return false
}

// TODO: Move filter func to util package?
func filter(ss []string, p func(s string) bool) []string {
	rs := []string{}

	for _, s := range ss {
		if p(s) {
			rs = append(rs, s)
		}
	}

	return rs
}

func jsonConfig() string {
	bs, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return "could not display config"
	}

	return string(bs)
}
