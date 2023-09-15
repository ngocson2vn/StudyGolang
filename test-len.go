package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	output1 := []string{}
	fmt.Println(len(output1))

	output2 := []string{"E1", "E2"}
	fmt.Println(len(output2))

	testPSMs := ""
	fmt.Printf("testPSMs = %v\n", testPSMs)
	results := strings.Split(strings.Trim(testPSMs, " "), ",")
	fmt.Printf("results = %v\n", results)
	fmt.Printf("len(results) = %v\n", len(results))

	crontaskPSM := "data.reckon.crontask"
	if v := os.Getenv("TCE_PSM"); v != "" {
		crontaskPSM = v
	}
	fmt.Printf("crontaskPSM = %v\n", crontaskPSM)
}
