package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	names := make(map[string]map[string]int)
	lines := strings.Split(string(file[:len(file)-1]), "\n")
	for _, l := range lines {
		line := strings.Split(l, " ")
		n, err := strconv.Atoi(line[3])
		if err != nil {
			panic(err)
		}
		if _, ok := names[line[0]]; !ok {
			names[line[0]] = make(map[string]int)
		}
		if line[2] == "gain" {
			names[line[0]][line[10][:len(line[10])-1]] += n
		} else {
			names[line[0]][line[10][:len(line[10])-1]] -= n
		}
	}
	fmt.Println(len(names))
	fmt.Println(names)
}
