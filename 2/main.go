package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type box struct {
	l, w, h int
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input[:len(input)-1]), "\n")

	boxes := make([]box, len(lines))

	for i, line := range lines {
		vals := strings.Split(line, "x")
		boxes[i].l, err = strconv.Atoi(vals[0])
		boxes[i].w, err = strconv.Atoi(vals[1])
		boxes[i].h, err = strconv.Atoi(vals[2])
	}
	var wrapsize int
	var ribbonsize int
	for _, v := range boxes {
		wrapsize += wrapSize(v)
		ribbonsize += ribbonSize(v)
	}
	fmt.Println(wrapsize)
	fmt.Println(ribbonsize)
}
func wrapSize(b box) int {
	l := b.l
	w := b.w
	h := b.h

	size := 2*l*w + 2*w*h + 2*h*l
	switch {
	case l <= w && l <= h:
		if w <= h {
			size += l * w
		} else {
			size += l * h
		}
	case w <= l && w <= h:
		if l <= h {
			size += w * l
		} else {
			size += w * h
		}
	case h <= l && h <= w:
		if l <= w {
			size += h * l
		} else {
			size += h * w
		}
	}
	return size
}

func ribbonSize(b box) int {
	l := b.l
	w := b.w
	h := b.h

	size := l * w * h
	switch {
	case l <= w && l <= h:
		if w <= h {
			size += l + l + w + w
		} else {
			size += l + l + h + h
		}
	case w <= l && w <= h:
		if l <= h {
			size += w + w + l + l
		} else {
			size += w + w + h + h
		}
	case h <= l && h <= w:
		if l <= w {
			size += h + h + l + l
		} else {
			size += h + h + w + w
		}
	}
	return size
}
