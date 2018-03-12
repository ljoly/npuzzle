package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

func parse(file string) ([][]int) {
	var size, x = 0, 0
	var tab [][] int
	var lines = strings.Split(file, "\n")
	for i := 0; i < len(lines) - 1; i++ {
		if (lines[i][0] != '#') {
			var line = strings.Fields(lines[i])
			if (len(line) == 1) {
				val, err := strconv.Atoi(line[0])
				if (err != nil) {
					panic(err)
				}
				size = val
				tab = make([][]int, size)
			} else if (len(line) == size) {
				tab[x] = make([]int, size)
				for y := 0; y < size; y++ {
					val, err := strconv.Atoi(line[y])
					if (err != nil) {
						panic(err)
					}
					tab[x][y] = val
				}
				x += 1
			}
		}
	}
	return (tab)
}

func main() {
	var arg = os.Args[1:]
	if len(arg) != 1 {
		panic("Wrong number of arguments : need 1")
	}
	file, err := ioutil.ReadFile(arg[0])
    if err != nil {
        panic(err)
	}
	fmt.Println(string(file))
	tab := parse(string(file))
	fmt.Println(tab)
}