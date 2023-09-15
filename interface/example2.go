package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// var period = flag.Duration("period", 1*time.Second, "Sleep period")
	// flag.Parse()
	// fmt.Printf("Sleep for %v...", *period)
	// time.Sleep(*period)
	// fmt.Println()

	var w io.Writer
	fmt.Printf("%T\n", w)

	w = os.Stdout
	fmt.Printf("%T\n", w)

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)

	fmt.Println()

	w.Write([]byte("Hello"))
	var buf = w.(*bytes.Buffer)
	fmt.Println(buf.String())
}
