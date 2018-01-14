package main

import (
	"errors"
	"fmt"
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

	// sizeSqrt := int(math.Sqrt(float64(size)))
	// for i := 0; i < 3; i++ {
	// 	arr := generateRandomBox(size)
	// 	for x := 0; x < 3; x++ {
	// 		for y := 0; y < 3; y++ {
	// 			switch i {
	// 			case 0:
	// 				matrixCell[x][y] = arr[x][y]
	// 			case 1:
	// 				matrixCell[x+sizeSqrt][y+sizeSqrt] = arr[x][y]
	// 			case 2:
	// 				matrixCell[x+(sizeSqrt*2)][y+(sizeSqrt*2)] = arr[x][y]
	// 			}
	// 		}
	// 	}
	// }
	//fill the Remaining Box
	for {
		passed, err := fill3x3(matrixCell, 0, 0)
		if passed {
			fmt.Println(err)
			break
		}
	}
	printMatrix(matrixCell)

	for {
		passed, err := fill3x3(matrixCell, 0, 3)
		if passed {
			fmt.Println(err)
			break
		}
	}
	printMatrix(matrixCell)

	for {
		passed, err := fill3x3(matrixCell, 0, 6)
		if passed {
			fmt.Println(err)
			break
		}
	}
	printMatrix(matrixCell)

	for {
		passed, err := fill3x3(matrixCell, 3, 0)
		if passed {
			fmt.Println(err)
			break
		}
	}
	printMatrix(matrixCell)

	for {
		passed, err := fill3x3(matrixCell, 3, 3)
		if passed {
			fmt.Println(err)
			break
		}
	}
	printMatrix(matrixCell)

	for {
		passed, err := fill3x3(matrixCell, 3, 6)
		if passed {
			fmt.Println(err)
			break
		}
	}
	printMatrix(matrixCell)

	for {
		passed, err := fill3x3(matrixCell, 6, 0)
		if passed {
			fmt.Println(err)
			break
		}
	}
	printMatrix(matrixCell)

	for {
		passed, err := fill3x3(matrixCell, 6, 3)
		if passed {
			fmt.Println(err)
			break
		}
	}
	printMatrix(matrixCell)

	for {
		passed, err := fill3x3(matrixCell, 6, 6)
		if passed {
			fmt.Println(err)
			break
		}
	}
	printMatrix(matrixCell)
	//Fill remaining boxes with backtracking
}

func generateRandomBox(size int) [][]int {
	arr := make([][]int, size)
	for i := range arr {
		arr[i] = make([]int, size)
	}

	uniqueArr := uniqueArrGenerator()
	count := 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			arr[x][y] = uniqueArr[count]
			count++
		}
	}

	return arr
}

func uniqueArrGenerator() []int {
	uniqueArr := rand.Perm(9 - 1 + 1)

	for i := 0; i < 9; i++ {
		if uniqueArr[i] == 0 {
			uniqueArr[i] = 9
		}
	}
	return uniqueArr
}

func printMatrix(matrixCell [][]int) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			fmt.Printf("%d  ", matrixCell[x][y])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func checkColifSafe(matrixCell [][]int, num int, y int) bool {
	for x := 0; x < 9; x++ {
		if matrixCell[x][y] == num {
			return false
		}
	}
	return true
}

func checkRowifSafe(matrixCell [][]int, num int, x int) bool {
	for y := 0; y < 9; y++ {
		if matrixCell[x][y] == num {
			return false
		}
	}
	return true
}

func fill3x3(matrixCell [][]int, x int, y int) (bool, error) {
	uniqueArr := uniqueArrGenerator()
	copyUnique := make([]int, len(uniqueArr))
	copy(copyUnique, uniqueArr)
	count := 0
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			for {
				passed := checkColifSafe(matrixCell, uniqueArr[count], j) && checkRowifSafe(matrixCell, uniqueArr[count], i)
				if passed {
					matrixCell[i][j] = uniqueArr[count]
					uniqueArr = append(uniqueArr[:count], uniqueArr[count+1:]...)
					count = 0
					break
				} else {
					if len(uniqueArr) == (count + 1) {
						break
					} else {
						count++
					}
				}
			}
		}
	}
	//put check 3x3 filled here
	filled := check3x3Filled(matrixCell, x, y)
	unique := check3x3Unique(matrixCell, copyUnique, x, y)
	// fmt.Printf("Filled is: %t\n", filled)
	// fmt.Printf("Unique is: %t\n", unique)
	return (unique && filled), errors.New("Array out of bounds")
}

func check3x3Filled(matrixCell [][]int, x int, y int) bool {
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if matrixCell[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func check3x3Unique(matrixCell [][]int, unique []int, x int, y int) bool {
	copyUnique := make([]int, len(unique))
	copy(copyUnique, unique)
	// fmt.Println(matrixCell[0:3][0:3])
	count := 0
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			for {
				if matrixCell[i][j] == copyUnique[count] {
					copyUnique = append(copyUnique[:count], copyUnique[count+1:]...)
					count = 0
					break
				} else {
					if len(copyUnique) == (count + 1) {
						return false
					} else {
						count++
					}
				}
			}
		}
	}

	return true
}
