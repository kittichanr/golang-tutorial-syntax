package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// step run this file
	// go build clitest.go
	// ./clitest 24 43 12 9 10

	fmt.Println(os.Args)
	args := os.Args[1:]
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max value :", max)
}
