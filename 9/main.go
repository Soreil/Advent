package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input[:len(input)-1]), "\n")

	places := make(map[string][]struct {
		string
		int
	})
	var placeNames []string
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		n, _ := strconv.Atoi(tokens[4])
		places[tokens[0]] = append(places[tokens[0]], struct {
			string
			int
		}{tokens[2], n})
		places[tokens[2]] = append(places[tokens[2]], struct {
			string
			int
		}{tokens[0], n})
	}
	for k, _ := range places {
		placeNames = append(placeNames, k)
	}
	heapPermute(placeNames, len(placeNames))

	var minimum int = 1000000
	var max int = 0
	for _, perm := range output {
		var sum int
		for _, v := range places[perm[0]] {
			if v.string == perm[1] {
				sum += v.int
			}
		}
		for _, v := range places[perm[1]] {
			if v.string == perm[2] {
				sum += v.int
			}
		}
		for _, v := range places[perm[2]] {
			if v.string == perm[3] {
				sum += v.int
			}
		}
		for _, v := range places[perm[3]] {
			if v.string == perm[4] {
				sum += v.int
			}
		}
		for _, v := range places[perm[4]] {
			if v.string == perm[5] {
				sum += v.int
			}
		}
		for _, v := range places[perm[5]] {
			if v.string == perm[6] {
				sum += v.int
			}
		}
		for _, v := range places[perm[6]] {
			if v.string == perm[7] {
				sum += v.int
			}
		}
		if sum < minimum {
			minimum = sum
		}
		if sum > max {
			max = sum
		}
	}
	fmt.Println(minimum)
	fmt.Println(max)
}

var output [][]string

func heapPermute(s []string, n int) {
	if n == 1 {
		tmp := make([]string, len(s))
		copy(tmp, s)
		output = append(output, tmp)
	} else {
		for i := 0; i < n; i++ {
			heapPermute(s, n-1)
			if n%2 == 1 {
				s[0], s[n-1] = s[n-1], s[0]
			} else {
				s[i], s[n-1] = s[n-1], s[i]
			}
		}
	}
}
