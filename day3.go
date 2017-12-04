package main

import (
	"fmt"
	"math"
)

func main() {
	input := 368078
	puzzle1(input)
	puzzle2(2)
}

func puzzle1(input int) {
	size := calcMatrixSize(input)
	currentValue := int(math.Pow(float64(size), 2))
	currentX, currentY := size-1, size-1
	centerX, centerY := getCenter(size)
	for ; currentX > 0; {
		if currentValue == input {break}
		currentValue--
		currentX--
	}
	for ; currentY > 0; {
		if currentValue == input {break}
		currentValue--
		currentY--
	}
	for ; currentX < size-1; {
		if currentValue == input {break}
		currentValue--
		currentX++
	}
	for ; currentY < size-1; {
		if currentValue == input {break}
		currentValue--
		currentY++
	}
	dist := math.Abs(float64(currentX) - float64(centerX)) + math.Abs(float64(currentY) - float64(centerY))
	fmt.Println(dist)
}

func getCenter(max int) (x, y int) {
	return int(math.Floor(float64(max)/2.0)), int(math.Floor(float64(max)/2.0))
}

func calcMatrixSize(max int) int {
	size := int(math.Ceil(math.Sqrt(float64(max))))
	if size%2 == 0 {
		size++
	}
	return size
}

func puzzle2(input int) {
	currentValue := 1
	size := calcMatrixSize(int(currentValue))
	matrix := makeMatrix(size)
	currentX, currentY := 0, 0
	matrix[currentX][currentY] = currentValue
	for ; currentValue <= input; {
		if currentX == size-1 && currentY == size-1 {
			matrix = addLayer(matrix)
			currentX++
			currentY++
			size = size + 2
		}
		break
	}
	fmt.Println(matrix)
	fmt.Println(currentX, currentY)
	fmt.Println(currentValue)
}

func makeMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i:= range matrix {
		matrix[i] = make([]int, size)
	}
	return matrix
}

func addLayer(matrix [][]int) [][]int {
	for i:= range matrix {
		matrix[i] = append([]int{0}, matrix[i]...)
		matrix[i] = append(matrix[i], 0)
	}
	base := make([][]int, 1)
	base[0] = make([]int, len(matrix[0]))
	matrix = append(base, matrix...)
	matrix = append(matrix, make([]int, len(matrix[0])))
	return matrix
}