// Package main is a simple command line calculator using RpnCalc
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/callerobertsson/rpncalc/rpncalc"
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

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	prompt(r, "")
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

		prompt(r, msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func prompt(r *rpncalc.RpnCalc, msg string) {
	if config.ShowStack {
		cmdStack(r, []string{})
	}
	fmt.Printf("%v", formatVal(r.Val()))

	if msg == "" {
		fmt.Printf(" > ")
		return
	}
	fmt.Printf(" [%v] > ", msg)
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
