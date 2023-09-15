package main

import "fmt"

func sumRecursive(n int) int {
	if n == 1 {
		return n
	}

	s := n + sumRecursive(n - 1)
	return s
}

func main() {
	n := 100
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("sum =", sum)

	sum2 := 0
	sum2 = sumRecursive(n)
	
	fmt.Println("sum2 =", sum2)	
}