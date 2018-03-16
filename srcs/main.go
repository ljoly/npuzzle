package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Env struct {
	initState  [][]int
	finalState [][]int
	file       string
	boardSize  int
}

func main() {
	var arg = os.Args[1:]
	if len(arg) != 1 {
		printError("Wrong number of arguments : need 1")
	}
	e := Env{}
	file, err := ioutil.ReadFile(arg[0])
	e.file = string(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(e.file))
	tab, size := parseFile(string(e.file))
	e.initState = tab
	e.boardSize = size
	fmt.Println(tab)
	play(e)
}
