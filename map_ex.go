package main

import "fmt"

func main() {
	m := make(map[string]*string)
	v := m["name"]
	if v != nil {
		fmt.Println("OK")
		fmt.Printf("%v\n", v)
	} else {
		fmt.Println("Not found")
	}
}
