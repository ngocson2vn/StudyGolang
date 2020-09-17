package main

import "fmt"

type Vertex struct {
	Lat  float64
	Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])

    fmt.Println("====================================")
    for k, v := range m {
        fmt.Println(k)
        fmt.Println(v)
    }
}
