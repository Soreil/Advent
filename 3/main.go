package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const part = 2

type house struct {
	x, y int
}

type presentCount map[house]int

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	hits := make(presentCount)
	if part == 1 {
		var curpos house
		hits[curpos]++
		for _, v := range input {
			switch v {
			case '<':
				curpos.x--
			case 'v':
				curpos.y++
			case '>':
				curpos.x++
			case '^':
				curpos.y--
			}
			hits[curpos]++
		}
	}
	if part == 2 {
		var curposSanta house
		var curposRobo house
		hits[curposSanta]++
		hits[curposRobo]++
		for i, v := range input {
			if i%2 == 0 {
				switch v {
				case '<':
					curposSanta.x--
				case 'v':
					curposSanta.y++
				case '>':
					curposSanta.x++
				case '^':
					curposSanta.y--
				}
				hits[curposSanta]++
			} else {
				switch v {
				case '<':
					curposRobo.x--
				case 'v':
					curposRobo.y++
				case '>':
					curposRobo.x++
				case '^':
					curposRobo.y--
				}
				hits[curposRobo]++

			}
		}

	}

	fmt.Println(len(hits))
}
