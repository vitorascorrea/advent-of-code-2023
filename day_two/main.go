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
	f, err := os.Open("day_two/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var possibleGameIdsSum int = 0
	var minSetPowerSum int = 0
	var maxReds int = 12
	var maxGreens int = 13
	var maxBlues int = 14

	for scanner.Scan() {
		var line = scanner.Text()

		if id, possible := isPossibleGame(line, maxReds, maxGreens, maxBlues); possible == true {
			possibleGameIdsSum += id
		}

		minSetPowerSum += calculateMinSetPowerSum(line)
	}

	fmt.Printf("Total sum: %v", possibleGameIdsSum)
	fmt.Print("\n")
	fmt.Printf("Min Set Power sum: %v", minSetPowerSum)
	fmt.Print("\n")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isPossibleGame(line string, maxReds int, maxGreens int, maxBlues int) (int, bool) {
	var id int = extractIdFromLine(line)

	var revealedSetsArray = strings.Split(strings.Split(line, ": ")[1], "; ")

	for _, set := range revealedSetsArray {
		var reds, greens, blues = extractColorValuesFromSet(set)

		if reds > maxReds || greens > maxGreens || blues > maxBlues {
			return id, false
		}
	}

	return id, true
}

func extractIdFromLine(line string) int {
	var gameIdString string = strings.Split(line, ":")[0]
	var idString string = strings.Split(gameIdString, "Game ")[1]

	return tryConvertStrToInt(idString)
}

func extractColorValuesFromSet(set string) (int, int, int) {
	var redValue = -1
	var greenValue = -1
	var blueValue = -1

	var splitSet = strings.Split(set, ", ")

	for _, s := range splitSet {
		if strings.Contains(s, "red") {
			var stringValue = strings.Split(s, " ")[0]
			redValue = tryConvertStrToInt(stringValue)
		}

		if strings.Contains(s, "green") {
			var stringValue = strings.Split(s, " ")[0]
			greenValue = tryConvertStrToInt(stringValue)
		}

		if strings.Contains(s, "blue") {
			var stringValue = strings.Split(s, " ")[0]
			blueValue = tryConvertStrToInt(stringValue)
		}
	}

	return redValue, greenValue, blueValue
}

func calculateMinSetPowerSum(line string) int {
	var revealedSetsArray = strings.Split(strings.Split(line, ": ")[1], "; ")

	var minReds = 1
	var minGreens = 1
	var minBlues = 1

	for _, set := range revealedSetsArray {
		var reds, greens, blues = extractColorValuesFromSet(set)

		if reds > minReds {
			minReds = reds
		}

		if blues > minBlues {
			minBlues = blues
		}

		if greens > minGreens {
			minGreens = greens
		}
	}

	return minReds * minBlues * minGreens
}

func tryConvertStrToInt(strValue string) int {
	if value, error := strconv.Atoi(strValue); error == nil {
		return value
	}

	return 0
}
