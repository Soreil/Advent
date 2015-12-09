package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type operator string

type gate struct {
	r  []string
	op operator
	n  uint16
}

var m map[string]uint16
var moves []gate

const (
	AND    operator = "AND"
	OR     operator = "OR"
	NOT    operator = "OR"
	LSHIFT operator = "LSHIFT"
	RSHIFT operator = "RSHIFT"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	m = make(map[string]uint16)
	lines := strings.Split(string(input[:len(input)-1]), "\n")

	for _, l := range lines {
		line := strings.Split(l, " ")
		var g gate
		for _, v := range line {
			if isLower(v) {
				g.r = append(g.r, v)
			}
			if isUpper(v) {
				g.op = operator(v)
			}
			if isNumeric(v) {
				n, _ := strconv.Atoi(v)
				g.n = uint16(n)
			}
		}
		moves = append(moves, g)
	}
	fmt.Println(moves)
	fmt.Println(m)
}

func isLower(s string) bool {
	for _, v := range s {
		if v < 'a' || v > 'z' {
			return false
		}
	}
	return true
}

func isUpper(s string) bool {
	for _, v := range s {
		if v < 'A' || v > 'Z' {
			return false
		}
	}
	return true
}

func isNumeric(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
