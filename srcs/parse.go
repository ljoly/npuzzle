package main

import "strings"

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
