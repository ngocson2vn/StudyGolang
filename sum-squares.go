package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func sumSquare(sum []int, step int, Yn []string, index int) {
	if index >= (len(Yn)) {
		return
	}

	y, err := strconv.Atoi(Yn[index])
	if err != nil {
		os.Exit(1)
	}

	if y >= -100 && y <= 100 {
		if y > 0 {
			sum[step] += (y * y)
		}
	} else {
		os.Exit(1)
	}

	index = index + 1
	sumSquare(sum, step, Yn, index)
}

func process(sum []int, step int, N int) {
	if step >= N {
		return
	}

	textX, _ := reader.ReadString('\n')
	textX = strings.TrimRight(textX, "\n")
	X, err := strconv.Atoi(textX)
	if err != nil {
		os.Exit(1)
	}

	if X < 1 || X > 100 {
		os.Exit(1)
	}

	textYn, _ := reader.ReadString('\n')
	textYn = strings.TrimRight(textYn, "\n")
	Yn := strings.Split(textYn, " ")
	if len(Yn) != X {
		os.Exit(1)
	}

	sumSquare(sum, step, Yn, 0)

	step = step + 1
	process(sum, step, N)
}

func display(sum []int, index int, N int) {
	if index >= N {
		return
	}

	fmt.Println(sum[index])
	index = index + 1
	display(sum, index, N)
}

func main() {

	textN, _ := reader.ReadString('\n')
	textN = strings.TrimRight(textN, "\n")
	N, err := strconv.Atoi(textN)

	if err != nil {
		os.Exit(1)
	}

	if N < 1 || N > 100 {
		os.Exit(1)
	}

	var sum = make([]int, N)

	process(sum, 0, N)

	display(sum, 0, N)
}
