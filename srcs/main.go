package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

func print_error(err string) {
	fmt.Println(err)
	os.Exit(0)
}

func atoi (str string) (int) {
	val, err := strconv.Atoi(str)
	if (err != nil) {
		panic(err)
	}
	return val
}

func get_size(line []string) (int, [][]int) {
	if (len(line) > 1 && line[1][0] != '#') {
		print_error("Error in file")
	}
	var size = atoi(line[0])
	var tab = make([][]int, size)
	return size, tab
}

func parse(file string) ([][]int, int) {
	var size, x = 0, 0
	var tab [][] int
	var lines = strings.Split(file, "\n")
	for i := 0; i < len(lines) - 1; i++ {
		if (lines[i][0] != '#') {
			var line = strings.Fields(lines[i])
			if (len(line) == 1 || size == 0) {
				size, tab = get_size(line)
			} else if (len(line) == size) {
				tab[x] = make([]int, size)
				for y := 0; y < size; y++ {
					tab[x][y] = atoi(line[y])
				}
				x += 1
			} else if (len(line) >= size) {
				if (line[size][0] == '#') {
					tab[x] = make([]int, size)
					for y := 0; y < size; y++ {
						tab[x][y] = atoi(line[y])
					}
					x += 1
				} else {
					print_error("Error in file")
				}
			} else {
				print_error("Error in file")
			}
		}
	}
	return tab, size
}

func main() {
	var arg = os.Args[1:]
	if len(arg) != 1 {
		print_error("Wrong number of arguments : need 1")
	}
	file, err := ioutil.ReadFile(arg[0])
    if err != nil {
        panic(err)
	}
	fmt.Println(string(file))
	tab, size := parse(string(file))
	fmt.Println(tab, size)
}