package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func doShell(l string) error {
	l = strings.TrimLeft(l, "!")

	// TODO: this will split uncorrectly when quoutations are used
	words := strings.Split(l, " ")
	cmd := words[0]
	args := []string{}
	if len(words) > 1 {
		args = words[1:]
	}

	fmt.Printf("Command %v, Args %#v\n", cmd, args)

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", string(out))
	return nil
}
