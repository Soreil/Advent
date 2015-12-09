package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input[:len(input)-1]), "\n")

	var total int
	for _, line := range lines {
		total -= len(line)
		l := strconv.Quote(line)
		if err != nil {
			panic(err)
		}
		total += len(l)
	}
	fmt.Println(total)
}
