package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("[main] Hello")
	go doSomething()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("[main] Hello again")
	}

	fmt.Println("[main] Exit")
}

func doSomething() {
	fmt.Println("[doSomething] Hello")
	go doSomethingElse()
	fmt.Println("[doSomething] Exit")
}

func doSomethingElse() {
	fmt.Println("[doSomethingElse] Now:", time.Now())
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("[doSomethingElse] Now:", time.Now())
	}
}
