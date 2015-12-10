package main

import "fmt"
import "math"

func main() {
	val := int64(3113322113)
	digs := digits(val)
	for i := 0; i < 50; i++ {
		digs = lookAndSee(digs)
	}
	fmt.Println(digs)
	fmt.Println(len(digs))
}

func digits(n int64) []int64 {
	l := int(math.Log10(float64(n)))
	d := make([]int64, l+1)
	for i := 0; n > 0; i++ {
		d[len(d)-1-i] = n % 10
		n /= 10
	}
	return d
}

func join(s []int64) int64 {
	var sum int64
	for i := len(s); i > 0; i-- {
		sum += s[i-1] * int64(math.Pow10(len(s)-i))
	}
	return sum
}

func lookAndSee(s []int64) []int64 {

	var output []int64
	for i := 0; i < len(s); i++ {
		length := int64(1)
		for i < len(s)-1 && s[i] == s[i+1] {
			length++
			i++
		}
		output = append(output, length)
		output = append(output, s[i])
	}
	return output
}
