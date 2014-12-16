package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() float64 {
	prev := float64(0)
	next := float64(0)
	return func() float64 {
		
		if next == 0 {
			next = float64(1)
			return next
		}
		
		if prev == 0 && next == 1 {
			prev = float64(1)
			return next
		}
		
		tmp := next
	    next = prev + next
		prev = tmp
		
		return next
	}
}

func main() {
	n := 100
	f := fibonacci()
	for i := 0; i < n; i++ {
		fmt.Print(" ", f())
		if i > 0 && i % 5 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
}