package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var engineMatrix [][]string

	for scanner.Scan() {
		var lineSplit = strings.Split(scanner.Text(), "")
		engineMatrix = append(engineMatrix, lineSplit)
	}

	var totalEngineSum = computeTotalEngineSum(engineMatrix)
	var totalGearSum = computeTotalGearSum(engineMatrix)

	fmt.Printf("Total engine sum: %v \n", totalEngineSum)
	fmt.Printf("Total gear sum: %v \n", totalGearSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeTotalEngineSum(engineMatrix [][]string) int {
	var result = 0

	for i := 0; i < len(engineMatrix); i++ {
		var hasAdjacentSymbol = false

		for j := 0; j < len(engineMatrix[i]); j++ {
			cell := engineMatrix[i][j]

			if isNumber(cell) && hasAdjacentSymbol == false && isNumberAdjacentToSymbol(engineMatrix, i, j) {
				hasAdjacentSymbol = true
			} else {
				if hasAdjacentSymbol {
					var fullNumber, newJ = getFullNumber(engineMatrix, i, j-1)
					result += fullNumber
					j = newJ + 1
				}

				hasAdjacentSymbol = false
			}
		}

		if hasAdjacentSymbol {
			var fullNumber, _ = getFullNumber(engineMatrix, i, len(engineMatrix[i])-1)
			result += fullNumber
		}
	}

	return result
}

func computeTotalGearSum(engineMatrix [][]string) int {
	var result = 0

	for i := 0; i < len(engineMatrix); i++ {
		for j := 0; j < len(engineMatrix[i]); j++ {
			cell := engineMatrix[i][j]

			if cell == "*" {
				var adjacentNumbers = isSymbolAdjacentToNumber(engineMatrix, i, j)
				if len(adjacentNumbers) == 2 {
					var gearRatio = 1

					for key := range adjacentNumbers {
						gearRatio *= key
					}

					result += gearRatio
				}
			}
		}
	}

	return result
}

func isNumberAdjacentToSymbol(matrix [][]string, cellI int, cellJ int) bool {
	surroundingCoordinates := [][]int{
		{cellI - 1, cellJ - 1},
		{cellI - 1, cellJ - 1},
		{cellI - 1, cellJ},
		{cellI - 1, cellJ + 1},
		{cellI, cellJ - 1},
		{cellI, cellJ + 1},
		{cellI + 1, cellJ - 1},
		{cellI + 1, cellJ},
		{cellI + 1, cellJ + 1},
	}

	for _, coord := range surroundingCoordinates {
		var inBound, value = getCellValue(matrix, coord[0], coord[1])

		if inBound && value != "." && !isNumber(value) {
			return true
		}
	}

	return false
}

func isSymbolAdjacentToNumber(matrix [][]string, cellI int, cellJ int) map[int]bool {
	surroundingCoordinates := [][]int{
		{cellI - 1, cellJ - 1},
		{cellI - 1, cellJ - 1},
		{cellI - 1, cellJ},
		{cellI - 1, cellJ + 1},
		{cellI, cellJ - 1},
		{cellI, cellJ + 1},
		{cellI + 1, cellJ - 1},
		{cellI + 1, cellJ},
		{cellI + 1, cellJ + 1},
	}

	adjacentNumbers := make(map[int]bool)

	for _, coord := range surroundingCoordinates {
		var inBound, value = getCellValue(matrix, coord[0], coord[1])

		if inBound && isNumber(value) {
			var fullNumber, _ = getFullNumber(matrix, coord[0], coord[1])
			adjacentNumbers[fullNumber] = true
		}
	}

	return adjacentNumbers
}

func getFullNumber(matrix [][]string, i int, j int) (int, int) {
	var fullNumber = matrix[i][j]

	// check left first
	for x := j - 1; x >= 0; x-- {
		if isNumber(matrix[i][x]) {
			fullNumber = matrix[i][x] + fullNumber
		} else {
			break
		}
	}

	var newJ = j

	for y := j + 1; y < len(matrix[i]); y++ {
		if isNumber(matrix[i][y]) {
			fullNumber = fullNumber + matrix[i][y]
			newJ += 1
		} else {
			break
		}
	}

	return tryConvertStrToInt(fullNumber), newJ
}

func getCellValue(matrix [][]string, i int, j int) (bool, string) {
	if i < 0 || i >= len(matrix) {
		return false, ""
	}

	if j < 0 || j >= len(matrix[i]) {
		return false, ""
	}

	return true, matrix[i][j]
}

func isNumber(char string) bool {
	return tryConvertStrToInt(char) != -1
}

func tryConvertStrToInt(value string) int {
	if intValue, error := strconv.Atoi(value); error == nil {
		return intValue
	}

	return -1
}
