package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	for i := range make([]int, 10) {
		x := i
		fmt.Printf("(%p) (%p) %d: %v\n", &i, &x, i+1, f())
	}
}
