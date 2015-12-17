package main

import (
	"fmt"
	"strings"
)

var input = "hepxcrrq"

func main() {
	for s := input; len(s) == 8 && s != "zzzzzzzz"; s = inc(s) {
		if !strings.ContainsAny(s, "iol") && hasDoubleDouble(s) && hasStraight(s) {
			fmt.Println("Solution found", s)
			break
		} else {
			fmt.Println(s)
		}
	}
}

func hasDoubleDouble(s string) bool {
	for c := "a"; c <= "z"; c = inc(c) {
		if n := strings.Index(s, c+c); n >= 0 {
			s = s[n:]
			for c := "a"; c <= "z"; c = inc(c) {
				if n := strings.Index(s, c+c); n >= 0 {
					return true
				}
			}
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
	for i := range s {
		if s[len(s)-i-1] != 'z' {
			var tmp string
			tmp = s[:len(s)-i-1] + string(s[len(s)-i-1]+1) + s[len(s)-i:]
			s = tmp
			break
		}
	}
	return s
}
