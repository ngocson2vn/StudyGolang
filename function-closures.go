package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	n := 101
	pos, neg := adder(), adder()
	for i:= 0; i < n; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}