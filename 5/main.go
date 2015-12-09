package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const step = 2

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input[:len(input)-1]), "\n")

	var count int
	for _, line := range lines {
		if niceString(line) {
			count++
		}
	}
	fmt.Println(count)
}

func niceString(s string) bool {
	if step != 2 {
		return threeVowel(s) && !illegals(s) && DoubleLetter(s)
	} else {
		return doublePair(s) && oneLetterApart(s)
	}
}

func doublePair(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func oneLetterApart(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i+2] == s[i] {
			return true
		}
	}
	return false
}

func threeVowel(s string) bool {
	var count int
	count += strings.Count(s, "a")
	count += strings.Count(s, "e")
	count += strings.Count(s, "i")
	count += strings.Count(s, "o")
	count += strings.Count(s, "u")
	return count >= 3
}
func illegals(s string) bool {
	return strings.Contains(s, "ab") ||
		strings.Contains(s, "cd") ||
		strings.Contains(s, "pq") ||
		strings.Contains(s, "xy")
}

func DoubleLetter(s string) bool {
	for c := 'a'; c <= 'z'; c++ {
		if strings.Contains(s, string(c)+string(c)) {
			return true
		}
	}
	return false
}
