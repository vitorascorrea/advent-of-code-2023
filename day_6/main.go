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
	f, err := os.Open("day_6/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var times []int
	var accTime int
	var distances []int
	var accDistance int

	for scanner.Scan() {
		var line = scanner.Text()
		var splitLine = strings.Split(strings.Split(line, ": ")[1], " ")

		if strings.Contains(line, "Time:") {
			times = parseArrayOfStringsToIntegers(splitLine)
			accTime = parseAndConcatenateArrayOfStrings(splitLine)
			continue
		}

		if strings.Contains(line, "Distance:") {
			distances = parseArrayOfStringsToIntegers(splitLine)
			accDistance = parseAndConcatenateArrayOfStrings(splitLine)
			continue
		}
	}

	var multipliedNumberOfWaysToWin = computeMultipliedNumberOfWaysToWin(times, distances)
	var multipliedNumberOfWaysToWinBigRace = computeMultipliedNumberOfWaysToWin([]int{accTime}, []int{accDistance})

	fmt.Printf("Multiplied number of ways to win: %v \n", multipliedNumberOfWaysToWin)
	fmt.Printf("Multiplied number of ways to win big race: %v \n", multipliedNumberOfWaysToWinBigRace)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeMultipliedNumberOfWaysToWin(times []int, distances []int) int {
	var result = 1

	for race := 0; race < len(times); race++ {
		var raceTime = times[race]
		var raceDistance = distances[race]
		var waysToWin = 0

		for possibleTime := 0; possibleTime < raceTime; possibleTime++ {
			var timeLeft = raceTime - possibleTime

			if timeLeft*possibleTime > raceDistance {
				waysToWin += 1
			}
		}

		result *= waysToWin
	}

	return result
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

func parseAndConcatenateArrayOfStrings(array []string) int {
	var concatenatedNumber string

	for _, c := range array {
		if c != "" {
			concatenatedNumber += c
		}
	}

	return tryConvertStrToInt(concatenatedNumber)
}

func tryConvertStrToInt(value string) int {
	if intValue, error := strconv.Atoi(value); error == nil {
		return intValue
	}

	return -1
}
