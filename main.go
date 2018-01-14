package main

import (
	"errors"
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
	//fill the Remaining Box
	// fillRemainingBox(matrixCell)
	passed, err := fill3x3(matrixCell, 3, 0)
	// fill3x3(matrixCell, 6, 0)
	// fill3x3(matrixCell, 6, 3)
	// fill3x3(matrixCell, 6, 0)
	for err != nil {
		passed, err = fill3x3(matrixCell, 3, 0)
		printMatrix(matrixCell)
		if passed {
			break
		}
		// fmt.Println(passed)
	}
	// for x := 0; x < 3; x++ {
	// 	for y := 0; y < 3; y++ {
	// 		fmt.Printf("Hi from, %d\n", y)
	// 		if x == 1 && y == 1 {
	// 			continue
	// 		}
	// 	}
	// }
	printMatrix(matrixCell)
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
			fmt.Printf("%d", matrixCell[x][y])
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

// func fillRemainingBox(matrixCell [][]int) {
// 	var randNo int
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			if matrixCell[i][j] == 0 {
// 				for {
// 					randNo = rand.Intn(10-1) + 1
// 					if checkColifSafe(matrixCell, randNo, j) && checkRowifSafe(matrixCell, randNo, i) {
// 						matrixCell[i][j] = randNo
// 						break
// 					}
// 					fmt.Println("Hi")
// 					printMatrix(matrixCell)
// 					fmt.Println(i)
// 					fmt.Println(j)
// 					fmt.Println(randNo)
// 				}
// 			} else {
// 				// fmt.Println("Hi")
// 				// fmt.Println(i)
// 				// fmt.Println(j)
// 				// printMatrix(matrixCell)
// 				fmt.Printf("\n")
// 				// fmt.Println(matrixCell[i][j])
// 				// fmt.Println(randNo)
// 			}
// 		}
// 	}
// }

func fill3x3(matrixCell [][]int, x int, y int) (bool, error) {
	uniqueArr := uniqueArrGenerator()
	count := 0
	pass := false
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			for {
				passed := checkColifSafe(matrixCell, uniqueArr[count], j) && checkRowifSafe(matrixCell, uniqueArr[count], i)
				if passed {
					matrixCell[i][j] = uniqueArr[count]
					uniqueArr = append(uniqueArr[:count], uniqueArr[count+1:]...)
					fmt.Println(uniqueArr)
					printMatrix(matrixCell)
					count = 0
					break
				} else {
					if len(uniqueArr) == (count + 1) {
						break
					} else {
						count++
					}
					// if count == 9 {
					// 	count = 0
					// 	printMatrix(matrixCell)
					// }
				}
			}
		}
	}
	//put check 3x3 filled here
	pass = check3x3Filled(matrixCell, x, y)
	// pass = true
	return pass, errors.New("sdasdas")
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
