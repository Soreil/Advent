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
	NOT    operator = "NOT"
	LSHIFT operator = "LSHIFT"
	RSHIFT operator = "RSHIFT"
	STORE  operator = ""
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	m = make(map[string]uint16)
	lines := strings.Split(string(input[:len(input)-1]), "\n")
	m["b"] = 16076

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
	for i := 0; i < len(moves); {
		if move(moves[i]) {
			m["b"] = 16076 //remove this line to get Solution A instead
			moves = append(moves[:i], moves[i+1:]...)
			i = 0
		} else {
			i++
		}
	}

	for _, v := range moves {
		if v.op == AND {
			fmt.Println(v)
		}
	}
	fmt.Println()

	fmt.Println(m)
	fmt.Println(m["a"])
}

func move(g gate) bool {
	switch g.op {
	case AND:
		if len(g.r) == 2 {
			r1, ok := m[g.r[0]]
			if ok {
				m[g.r[1]] = r1 & g.n
				return true
			}
		} else {
			r1, ok := m[g.r[0]]
			r2, ok2 := m[g.r[1]]
			if ok && ok2 {
				m[g.r[2]] = r1 & r2
				return true
			}
		}
	case OR:
		r1, ok := m[g.r[0]]
		r2, ok2 := m[g.r[1]]
		if ok && ok2 {
			m[g.r[2]] = r1 | r2
			return true
		}
	case NOT:
		r1, ok := m[g.r[0]]
		if ok {
			m[g.r[1]] = ^r1
			return true
		}
	case LSHIFT:
		r1, ok := m[g.r[0]]
		if ok {
			m[g.r[1]] = r1 << g.n
			return true
		}
	case RSHIFT:
		r1, ok := m[g.r[0]]
		if ok {
			m[g.r[1]] = r1 >> g.n
			return true
		}
	case STORE:
		if len(g.r) == 1 {
			m[g.r[0]] = g.n
			return true
		} else {
			r1, ok := m[g.r[0]]
			if ok {
				m[g.r[1]] = r1
				return true
			}
		}
	default:
		return false
	}
	return false
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
