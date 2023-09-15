package main

import (
    "fmt"
)

func main() {
    s := "INIT"
    fmt.Printf("%p\n", &s)
    fmt.Println(s)

    s = "Son"
    fmt.Printf("%p\n", &s)
    fmt.Println(s)
}
