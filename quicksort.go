package main

import "fmt"

func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func arrange(a []int, left int, right int, p int, pivot int) int {

	if a[left] <= pivot {
		p = p + 1
		swap(a, p, left)
	}

	left = left + 1
	if left < right {
		p = arrange(a, left, right, p, pivot)
	}

	return p
}

func partition(a []int, left int, right int) int {

	// pick up pivot value
	pivot := a[right]

	// initialize pivot index
	var p = left - 1

	// 1. scan the array from left to right - 1
	// 2. move element (which is less than or equal the pivot) to the left
	// 3. save index of the last element (which is less than or equal the pivot)

	// ****************************************
	// *     Using for loop                   *
	// ****************************************
	// for i := left; i <= right-1; i++ {
	// 	if a[i] <= pivot {
	// 		p = p + 1

	// 		if i > p {
	// 			swap(a, p, i)
	// 		}
	// 	}
	// }

	// Using recursive function
	p = arrange(a, left, right, p, pivot)

	p = p + 1
	swap(a, p, right)

	return p
}

func quicksort(a []int, left int, right int) {
	if left < right {
		var p = partition(a, left, right)
		quicksort(a, left, p-1)
		quicksort(a, p+1, right)
	}
}

func main() {
	var a = []int{49, 70, 97, 38, 57, 21, 85, 68, 76, 9, 81, 36, 55, 79, 74, 85, 16, 61, 77, 24, 63}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)
}
