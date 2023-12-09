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
	f, err := os.Open("day_9/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var histories [][]int

	for scanner.Scan() {
		var line = scanner.Text()
		var splitLine = strings.Split(line, " ")
		histories = append(histories, parseArrayOfStringsToIntegers(splitLine))
	}

	var sumOfExtrapolatedNextValues = computeSumOfExtrapolatedValues(histories, true)
	var sumOfExtrapolatedPreviousValues = computeSumOfExtrapolatedValues(histories, false)
	fmt.Printf("Sum of extrapolated next values: %v \n", sumOfExtrapolatedNextValues)
	fmt.Printf("Sum of extrapolated previous values: %v \n", sumOfExtrapolatedPreviousValues)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeSumOfExtrapolatedValues(histories [][]int, next bool) int {
	var total = 0

	for _, history := range histories {
		if next {
			total += computeNextValue(history)
		} else {
			total += computePreviousValue(history)
		}
	}

	return total
}

func computeNextValue(history []int) int {
	var nextValue = 0
	var historyDifferences = computeHistoryOfDifferences(history)

	for _, differences := range historyDifferences {
		nextValue += differences[len(differences)-1]
	}

	return nextValue
}

func computePreviousValue(history []int) int {
	var historyDifferences = computeHistoryOfDifferences(history)
	var currentPreviousValue = historyDifferences[len(historyDifferences)-1][0]

	for i := len(historyDifferences) - 2; i >= 0; i-- {
		currentPreviousValue = historyDifferences[i][0] - currentPreviousValue
	}

	return currentPreviousValue
}

func computeHistoryOfDifferences(history []int) [][]int {
	var foundAllEqual = false
	var historyDifferences = [][]int{history}

	for !foundAllEqual {
		var differences, allEqual = computeArrayOfDifferences(
			historyDifferences[len(historyDifferences)-1],
		)

		historyDifferences = append(historyDifferences, differences)
		foundAllEqual = allEqual
	}

	return historyDifferences
}

func computeArrayOfDifferences(array []int) ([]int, bool) {
	var differences []int
	var allEqual = true

	for i := 0; i < len(array)-1; i++ {
		var currentValue = array[i]
		var nextValue = array[i+1]
		var difference = nextValue - currentValue

		if len(differences) > 0 && allEqual {
			var latestDifference = differences[len(differences)-1]
			allEqual = difference == latestDifference
		}

		differences = append(differences, difference)
	}

	return differences, allEqual
}

func parseArrayOfStringsToIntegers(array []string) []int {
	var intArray []int

	for _, c := range array {
		if c != "" {
			var intValue = tryConvertStrToInt(c)
			intArray = append(intArray, intValue)
		}
	}

	return intArray
}

func tryConvertStrToInt(value string) int {
	if intValue, error := strconv.Atoi(value); error == nil {
		return intValue
	}

	return -1
}
