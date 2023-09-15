package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func play() {
	result := 0
	defer func() {
		fmt.Printf("Result: %d\n", result)
	}()
	result = add(10, 20)
}

func main() {
	// fmt.Println(add(42, 13))
	play()
	fmt.Println("The end of main")
}
