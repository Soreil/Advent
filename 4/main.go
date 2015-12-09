package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

const part = 2

func main() {
	input := "bgvyzdsv"
	prefix := []byte{0, 0, 0}
	for i := 0; i < 1000000000; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		//heh, funny thing is I already finished part two here without knowing
		if reflect.DeepEqual(hash[:3], prefix) || reflect.DeepEqual(hash[:2], prefix[:2]) && hash[2]&0xF0 == 0 {
			if part == 2 {
				if reflect.DeepEqual(hash[:3], prefix) {
					fmt.Println(i)
				} else {
					continue
				}
			} else {
				fmt.Println(i)
			}
			os.Exit(0)
		}
	}
}
