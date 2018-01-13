package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	matrixCell := make([][]int, 9)
	for i := range matrixCell {
		matrixCell[i] = make([]int, 9)
	}

	for i := 0; i < 3; i++ {
		arr := generateRandomBox()
		for x := 0; x < x+3; x++ {
			for y := 0; y < y+3; y++ {
				switch i {
				case 0:
					matrixCell[x][y] = arr[y]
				case 1:
					x = 3
					y = 3
					matrixCell[x][y] = arr[y]
				case 2:
					x = 6
					y = 6
					matrixCell[x][y] = arr[y]
				}
			}
		}
	}

	printMatrix(matrixCell)
}

func generateRandomBox() []int {
	arr := make([]int, 9)
	for i := 0; i < 9; i++ {
		arr[i] = rand.Intn(10)
	}

	return arr
}

func printMatrix([][]int) {

}
