package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

func getSize(line []string) (int, [][]int) {
	if len(line) > 1 && line[1][0] != '#' {
		printError("Error in file")
	}
	var size = atoi(line[0])
	var tab = make([][]int, size)
	return size, tab
}

func parseFile(file string) ([][]int, int) {
	var size, x = 0, 0
	var tab [][]int
	var lines = strings.Split(file, "\n")
	for i := 0; i < len(lines)-1; i++ {
		if lines[i][0] != '#' {
			var line = strings.Fields(lines[i])
			if len(line) == 1 || size == 0 {
				size, tab = getSize(line)
			} else if len(line) == size { // check que des nombres
				tab[x] = make([]int, size)
				for y := 0; y < size; y++ {
					tab[x][y] = atoi(line[y])
				}
				x++
			} else if len(line) > size {
				if line[size][0] == '#' {
					tab[x] = make([]int, size)
					for y := 0; y < size; y++ {
						tab[x][y] = atoi(line[y])
					}
					x++
				} else {
					printError("Error in file")
				}
			} else {
				printError("Error in file")
			}
		}
	}
	return tab, size
}

func getAnswer(size int) [][]int {
	var puzzle [][]int
	var cursor = 1
	var x = 0
	var ix = 1
	var y = 0
	var iy = 0

	puzzle = make([][]int, size)
	for i := 0; i < len(puzzle); i++ {
		puzzle[i] = make([]int, size)
		for j := 0; j < len(puzzle); j++ {
			puzzle[i][j] = -1
		}
	}
	for {
		puzzle[y][x] = cursor
		if cursor == 0 {
			break
		}
		cursor++
		if x+ix == size || x+ix < 0 || (ix != 0 && puzzle[y][x+ix] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == size || y+iy < 0 || (iy != 0 && puzzle[y+iy][x] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
		if cursor == size*size {
			cursor = 0
		}
	}
	return (puzzle)
}

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
	answer := getAnswer(size)
	fmt.Println(tab)
	for i := 0; i < size; i++ {
		fmt.Println(answer[i])
	}
	// fmt.Println(answer)
}
