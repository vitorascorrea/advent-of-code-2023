package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_11/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var originalMatrix [][]string

	for scanner.Scan() {
		var line = scanner.Text()
		var splitLine = strings.Split(line, "")

		originalMatrix = append(originalMatrix, splitLine)
	}

	fmt.Printf("Total sum for 1: %v \n", calculateDistancesSum(originalMatrix, 2))
	fmt.Printf("Total sum for 1000000: %v \n", calculateDistancesSum(originalMatrix, 1000000))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculateDistancesSum(matrix [][]string, fillSize int) uint64 {
	var galaxiesCoordinates [][]int
	var sum = 0.0

	var emptyRows []int
	var emptyCols []int

	for i := 0; i < len(matrix); i++ {
		var emptyRow = true

		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != "." {
				emptyRow = false
			}
		}

		if emptyRow {
			emptyRows = append(emptyRows, i)
		}
	}

	for col := 0; col < len(matrix[0]); col++ {
		var emptyCol = true

		for row := 0; row < len(matrix); row++ {
			if matrix[row][col] != "." {
				emptyCol = false
				break
			}
		}

		if emptyCol {
			emptyCols = append(emptyCols, col)
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "#" {
				var calculatedI = i
				var calculatedJ = j

				for _, row := range emptyRows {
					if i > row {
						calculatedI += fillSize - 1
					}
				}

				for _, col := range emptyCols {
					if j > col {
						calculatedJ += fillSize - 1
					}
				}

				galaxiesCoordinates = append(galaxiesCoordinates, []int{calculatedI, calculatedJ})
			}
		}
	}

	for i := 0; i < len(galaxiesCoordinates); i++ {
		for j := i + 1; j < len(galaxiesCoordinates); j++ {
			var firstGalaxy = galaxiesCoordinates[i]
			var secondGalaxy = galaxiesCoordinates[j]

			var manhattanDistance = math.Abs(float64(secondGalaxy[0])-float64(firstGalaxy[0])) + math.Abs(float64(secondGalaxy[1])-float64(firstGalaxy[1]))
			sum += manhattanDistance
		}
	}

	return uint64(sum)
}

func expandMatrix(originalMatrix [][]string) [][]string {
	var expandedMatrix [][]string

	for i := 0; i < len(originalMatrix); i++ {
		expandedMatrix = append(expandedMatrix, originalMatrix[i])
		var emptyRow = true

		for j := 0; j < len(originalMatrix[i]); j++ {
			if originalMatrix[i][j] != "." {
				emptyRow = false
				break
			}
		}

		if emptyRow {
			expandedMatrix = append(expandedMatrix, originalMatrix[i])
		}
	}

	var emptyColIndexes []int

	for col := 0; col < len(originalMatrix[0]); col++ {
		var emptyCol = true

		for row := 0; row < len(originalMatrix); row++ {
			if originalMatrix[row][col] != "." {
				emptyCol = false
				break
			}
		}

		if emptyCol {
			emptyColIndexes = append(emptyColIndexes, col)
		}
	}

	for row := 0; row < len(expandedMatrix); row++ {
		var indexAcc = 0

		for _, col := range emptyColIndexes {
			var indexToAdd = col + indexAcc
			expandedMatrix[row] = append(expandedMatrix[row][:indexToAdd+1], expandedMatrix[row][indexToAdd:]...)
			expandedMatrix[row][indexToAdd] = "."

			indexAcc += 1
		}
	}

	return expandedMatrix
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
}

func tryConvertStrToInt(value string) int {
	if intValue, error := strconv.Atoi(value); error == nil {
		return intValue
	}

	return -1
}
