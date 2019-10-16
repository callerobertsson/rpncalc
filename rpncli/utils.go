// Package main utility funcs
package main

func member(t string, ms ...string) bool {
	for _, m := range ms {
		if m == t {
			return true
		}
	}
	return false
}

func filter(ss []string, p func(s string) bool) []string {
	rs := []string{}

	for _, s := range ss {
		if p(s) {
			rs = append(rs, s)
		}
	}

	return rs
}
