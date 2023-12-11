package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	row int
	col int
}

func main() {
	f, err := os.Open("day_10/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var matrix [][]string
	var startingPoint Coordinate

	var row = 0

	for scanner.Scan() {
		var line = scanner.Text()
		var splitLine = strings.Split(line, "")

		for col, c := range splitLine {
			if c == "S" {
				startingPoint = Coordinate{row: row, col: col}
			}
		}

		matrix = append(matrix, splitLine)
		row += 1
	}

	var numStepsFarthestPoint = computeNumberOfStepsToFarthestPoint(matrix, startingPoint)
	var tilesEnclosedByLoop = computeNumberOfTilesEnclosedByLoop(matrix, startingPoint)
	fmt.Printf("Number of steps: %v \n", numStepsFarthestPoint)
	fmt.Printf("Number of tiles enclosed: %v \n", tilesEnclosedByLoop)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeNumberOfStepsToFarthestPoint(matrix [][]string, startingPoint Coordinate) int {
	var stepsByCoordinate = make(map[Coordinate]int)
	var currentCoord = startingPoint
	var connectedCoordinates = getConnectedCoordinates(matrix, currentCoord)
	var prevCoordDist = 0
	var farthestDist = 0

	for len(connectedCoordinates) > 0 {
		var newConnectedCoordinates []Coordinate

		for _, coord := range connectedCoordinates {
			var _, present = stepsByCoordinate[coord]

			if !present {
				stepsByCoordinate[coord] = prevCoordDist + 1

				if stepsByCoordinate[coord] > farthestDist {
					farthestDist = stepsByCoordinate[coord]
				}

				newConnectedCoordinates = append(
					newConnectedCoordinates,
					getConnectedCoordinates(matrix, coord)...,
				)
			}
		}

		prevCoordDist += 1
		connectedCoordinates = newConnectedCoordinates
	}

	return farthestDist
}

func computeNumberOfTilesEnclosedByLoop(matrix [][]string, originalStartingPoint Coordinate) int {
	var loopCoordinates = make(map[Coordinate]int)
	var insideLoopCount = 0
	var startingPointReplacement = findSuitableSReplacement(matrix, originalStartingPoint)
	var enlargedMatrix, startingPoint = enlargeMatrix(matrix, startingPointReplacement)
	var currentCoord = startingPoint

	loopCoordinates[startingPoint] = 1

	var connectedCoordinates = getConnectedCoordinates(enlargedMatrix, currentCoord)

	for len(connectedCoordinates) > 0 {
		var newConnectedCoordinates []Coordinate

		for _, coord := range connectedCoordinates {
			var _, present = loopCoordinates[coord]

			if !present {
				loopCoordinates[coord] = 1

				newConnectedCoordinates = append(
					newConnectedCoordinates,
					getConnectedCoordinates(enlargedMatrix, coord)...,
				)
			}
		}

		connectedCoordinates = newConnectedCoordinates
	}

	var guaranteedInsidePoint = findGuaranteedInsidePointAfterEnlarge(startingPoint, startingPointReplacement)

	floodFill(enlargedMatrix, guaranteedInsidePoint, loopCoordinates, "C")

	for i := 0; i < len(enlargedMatrix); i += 2 {
		for j := 0; j < len(enlargedMatrix[i]); j += 2 {
			if enlargedMatrix[i][j] == "C" {
				insideLoopCount += 1
			}
		}
	}

	return insideLoopCount
}

func enlargeMatrix(matrix [][]string, startingPointReplacement string) ([][]string, Coordinate) {
	var enlargedMatrix [][]string
	var startingPoint Coordinate

	for i := 0; i < len(matrix); i++ {
		enlargedMatrix = append(enlargedMatrix, matrix[i])
		var newRow []string

		for j := 0; j < len(matrix[i]); j++ {
			newRow = append(newRow, "|")
		}

		enlargedMatrix = append(enlargedMatrix, newRow)
	}

	for i := 0; i < len(enlargedMatrix); i++ {
		var newRow []string

		for j := 0; j < len(enlargedMatrix[i]); j++ {
			var previousSymbol = enlargedMatrix[i][j]

			newRow = append(newRow, previousSymbol)
			newRow = append(newRow, "-")
		}

		enlargedMatrix[i] = newRow
	}

	for i := 0; i < len(enlargedMatrix); i++ {
		for j := 0; j < len(enlargedMatrix[i]); j++ {
			if enlargedMatrix[i][j] == "S" {
				enlargedMatrix[i][j] = startingPointReplacement
				startingPoint = Coordinate{row: i, col: j}
			}
		}
	}

	return enlargedMatrix, startingPoint
}

func floodFill(matrix [][]string, currentPoint Coordinate, loopCoordinates map[Coordinate]int, characterToFill string) {
	if _, ok := loopCoordinates[currentPoint]; !inBounds(matrix, currentPoint) || ok || matrix[currentPoint.row][currentPoint.col] == characterToFill {
		return
	}

	matrix[currentPoint.row][currentPoint.col] = characterToFill

	// up
	floodFill(matrix, Coordinate{row: currentPoint.row - 1, col: currentPoint.col}, loopCoordinates, characterToFill)
	// down
	floodFill(matrix, Coordinate{row: currentPoint.row + 1, col: currentPoint.col}, loopCoordinates, characterToFill)
	// right
	floodFill(matrix, Coordinate{row: currentPoint.row, col: currentPoint.col - 1}, loopCoordinates, characterToFill)
	// right
	floodFill(matrix, Coordinate{row: currentPoint.row, col: currentPoint.col + 1}, loopCoordinates, characterToFill)
}

func getConnectedCoordinates(matrix [][]string, currentCoord Coordinate) []Coordinate {
	var connectedCoordinates []Coordinate

	var up = Coordinate{row: currentCoord.row - 1, col: currentCoord.col}
	var down = Coordinate{row: currentCoord.row + 1, col: currentCoord.col}
	var left = Coordinate{row: currentCoord.row, col: currentCoord.col - 1}
	var right = Coordinate{row: currentCoord.row, col: currentCoord.col + 1}

	var currentCell = matrix[currentCoord.row][currentCoord.col]

	if currentCell == "S" {
		currentCell = findSuitableSReplacement(matrix, currentCoord)
	}

	if currentCell == "7" {
		connectedCoordinates = append(connectedCoordinates, left)
		connectedCoordinates = append(connectedCoordinates, down)
	}

	if currentCell == "|" {
		connectedCoordinates = append(connectedCoordinates, up)
		connectedCoordinates = append(connectedCoordinates, down)
	}

	if currentCell == "F" {
		connectedCoordinates = append(connectedCoordinates, right)
		connectedCoordinates = append(connectedCoordinates, down)
	}

	if currentCell == "L" {
		connectedCoordinates = append(connectedCoordinates, right)
		connectedCoordinates = append(connectedCoordinates, up)
	}

	if currentCell == "J" {
		connectedCoordinates = append(connectedCoordinates, left)
		connectedCoordinates = append(connectedCoordinates, up)
	}

	if currentCell == "-" {
		connectedCoordinates = append(connectedCoordinates, left)
		connectedCoordinates = append(connectedCoordinates, right)
	}

	return connectedCoordinates
}

func findSuitableSReplacement(matrix [][]string, sCoord Coordinate) string {
	// have to connect to two symbols
	var up = Coordinate{row: sCoord.row - 1, col: sCoord.col}
	var down = Coordinate{row: sCoord.row + 1, col: sCoord.col}
	var left = Coordinate{row: sCoord.row, col: sCoord.col - 1}
	var right = Coordinate{row: sCoord.row, col: sCoord.col + 1}

	var possibleSymbols = map[string]bool{
		"J": true,
		"L": true,
		"7": true,
		"F": true,
		"|": true,
		"-": true,
	}

	if inBounds(matrix, up) && strings.Contains("|F7", matrix[up.row][up.col]) {
		possibleSymbols["-"] = false
		possibleSymbols["7"] = false
		possibleSymbols["F"] = false
	}

	if inBounds(matrix, down) && strings.Contains("|JL", matrix[down.row][down.col]) {
		possibleSymbols["-"] = false
		possibleSymbols["J"] = false
		possibleSymbols["L"] = false
	}

	if inBounds(matrix, left) && strings.Contains("-LF", matrix[left.row][left.col]) {
		possibleSymbols["|"] = false
		possibleSymbols["L"] = false
		possibleSymbols["F"] = false
	}

	if inBounds(matrix, right) && strings.Contains("-J7", matrix[right.row][right.col]) {
		possibleSymbols["|"] = false
		possibleSymbols["J"] = false
		possibleSymbols["7"] = false
	}

	for key := range possibleSymbols {
		if possibleSymbols[key] {
			return key
		}
	}

	return "-"
}

func findGuaranteedInsidePointAfterEnlarge(sCoord Coordinate, sReplacement string) Coordinate {
	switch sReplacement {
	case "J":
		return Coordinate{row: sCoord.row - 1, col: sCoord.col - 1}
	case "F":
		return Coordinate{row: sCoord.row + 1, col: sCoord.col + 1}
	case "L":
		return Coordinate{row: sCoord.row - 1, col: sCoord.col + 1}
	case "7":
		return Coordinate{row: sCoord.row + 1, col: sCoord.col - 1}
	default:
		// this would be a guess between left or right I think
		return Coordinate{row: sCoord.row, col: sCoord.col + 1}
	}
}

func inBounds(matrix [][]string, coordinate Coordinate) bool {
	if coordinate.row < 0 || coordinate.row >= len(matrix) {
		return false
	}

	if coordinate.col < 0 || coordinate.col >= len(matrix[0]) {
		return false
	}

	return true
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
