package main

import "fmt"

func main() {
	var i, j = 0, 0
	var min = 0
	var iter = 1
	var tab [10][10]int
	var N = 10
	var init = 10
	for {
		
		if (iter > init * init) {
			fmt.Println(iter)
			fmt.Println("END ?")
			break
		}
		
		if (i == min) {
			if (j != N - 1) {
				tab[i][j] = iter
				j++
			} else {
				tab[i][j] = iter
				i++
			}
		} else if (j == N - 1) {
			if (i != N - 1) {
				tab[i][j] = iter
				i++
			} else {
				tab[i][j] = iter
				j--
			}
		} else if (i == N - 1) {
			if (j != min) {
				tab[i][j] = iter
				j--
			} else {
				tab[i][j] = iter
				i--
			}
		} else if (j == min) {
			if (i != min) {
				tab[i][j] = iter
				i--
			} else {
				tab[i][j] = iter
				j++
			}
		}
		if (iter == 4 * N - 4) {
			N--
			min++
			i++
		}
		iter++
		fmt.Println(tab[0])
		fmt.Println(tab[1])
		fmt.Println(tab[2])
		fmt.Println(tab[3])
		fmt.Println(tab[4])
		fmt.Println(tab[5])
		fmt.Println(tab[6])
		fmt.Println(tab[7])
		fmt.Println(tab[8])
		fmt.Println(tab[9])
		fmt.Println("-----------------")
	}
	fmt.Println(tab)
}