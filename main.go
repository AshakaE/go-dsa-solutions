package main

import (
	"fmt"
	"os"
	"github.com/AshakaE/go-dsa-solutions/samples"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <algorithm>")
		fmt.Println("Available: arraysum, diagonaldiff")
		return
	}

	switch os.Args[1] {
	case "arraySum":
		samples.SimpleArraySumTest()
	case "diagonalDiff":
		samples.DiagonalDifferenceTest()
	case "minMaxSum":
		samples.MiniMaxSumTest()
	default:
		fmt.Println("Unknown algorithm:", os.Args[1])
	}
}
