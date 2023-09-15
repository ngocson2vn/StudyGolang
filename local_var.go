package main

import "fmt"

func main() {
	x := 100
	fmt.Printf("x is at %p; x = %d\n\n", &x, x)
	for x := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("x is at %p, x = %d\n", &x, x)
	}
	fmt.Println()
	fmt.Printf("x is at %p; x = %d\n", &x, x)
}
