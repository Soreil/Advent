package main

import (
	"fmt"
	"strings"
)

var input = "hepxxzaa"

func main() {
	for s := input; len(s) == 8 && s != "zzzzzzzz"; s = inc(s) {
		if !strings.ContainsAny(s, "iol") && hasDoubleDouble(s) && hasStraight(s) {
			fmt.Println("Solution found", s)
			break
		}
	}
}

func hasDoubleDouble(s string) bool {
	for c := "a"; c <= "z"; c = inc(c) {
		if n := strings.Index(s, c+c); n >= 0 {
			for c = inc(c); c <= "z"; c = inc(c) {
				if n := strings.Index(s, c+c); n >= 0 {
					return true
				}
			}
			return false
		}
	}
	return false
}

func hasStraight(s string) bool {
	for c := 'a'; c <= 'z'-2; c++ {
		if strings.Contains(s, string(c)+string(c+1)+string(c+2)) {
			return true
		}
	}
	return false
}

func inc(s string) string {
	if s[len(s)-1] == 'z' {
		if len(s) != 1 {
			return inc(s[:len(s)-1]) + "a"
		} else {
			return "az"
		}
	} else {
		return s[:len(s)-1] + string(s[len(s)-1]+1)
	}
}
