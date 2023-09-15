package main

import (
	"fmt"
	"flag"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Arg(0))
	fmt.Println(flag.Arg(1))
}
