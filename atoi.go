package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi(os.Getenv("RETRY_COUNT"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(i)
}
