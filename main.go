package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	size := 9
	rand.Seed(time.Now().UTC().UnixNano())
	matrixCell := make([][]int, size)
	for i := range matrixCell {
		matrixCell[i] = make([]int, size)
	}

	sizeSqrt := int(math.Sqrt(float64(size)))
	for i := 0; i < 3; i++ {
		arr := generateRandomBox(size)
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				switch i {
				case 0:
					matrixCell[x][y] = arr[x][y]
				case 1:
					matrixCell[x+sizeSqrt][y+sizeSqrt] = arr[x][y]
				case 2:
					matrixCell[x+(sizeSqrt*2)][y+(sizeSqrt*2)] = arr[x][y]
				}
			}
		}
	}
	printMatrix(matrixCell)
}

func generateRandomBox(size int) [][]int {
	arr := make([][]int, size)
	for i := range arr {
		arr[i] = make([]int, size)
	}

	uniqueArr := rand.Perm(9 - 1 + 1)
	count := 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if uniqueArr[count] == 0 {
				uniqueArr[count] = 9
			}
			arr[x][y] = uniqueArr[count]
			count++
		}
	}

	return arr
}

func printMatrix(matrixCell [][]int) {
	fmt.Println(matrixCell)
}

func checkUnique(matrixCell [][]int, size int) {

}
