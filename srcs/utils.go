package main

import (
	"fmt"
	"os"
	"strconv"
)

func printError(err string) {
	fmt.Println(err)
	os.Exit(0)
}

func atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}
