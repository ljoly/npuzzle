package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	initState [][]int
	file      string
	size      string
)

func main() {
	var arg = os.Args[1:]
	if len(arg) != 1 {
		printError("Wrong number of arguments : need 1")
	}
	file, err := ioutil.ReadFile(arg[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
	tab, size := parseFile(string(file))
	fmt.Println(tab)
	play(size)
}
