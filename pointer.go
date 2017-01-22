package main

import "fmt"

func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("Address of a variable: %x\n", &a)

	fmt.Printf("Address stored in ip variable: %x\n", ip)

	fmt.Printf("Value of *ip variable: %d\n", *ip)

	// pointer to pointer
	var ptr *int
	var pptr **int
	ptr = &a
	pptr = &ptr

	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *ptr = %d\n", *ptr)
	fmt.Printf("Value available at **pptr = %d\n", **pptr)

	// passing pointers to functions
	a = 100
	var b int = 200

	fmt.Printf("Before swap, value of a: %d\n", a)
	fmt.Printf("Before swap, value of b: %d\n", b)

	swap(&a, &b)

	fmt.Printf("After swap, value of a: %d\n", a)
	fmt.Printf("After swap, value of b: %d\n", b)
}

func swap(x *int, y *int) {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
}