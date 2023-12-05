package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("day_5/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var seeds []int
	var mapNames []string
	var maps = make(map[string][][]int)
	var currentMapName string
	var currentMapConversionTables [][]int
	var parsingMap = false

	for scanner.Scan() {
		var line = scanner.Text()

		if line == "" {
			parsingMap = false

			if len(currentMapConversionTables) > 0 {
				maps[currentMapName] = currentMapConversionTables
			}

			continue
		}

		if strings.Contains(line, "seeds:") {
			seeds = parseArrayOfStringsToIntegers(strings.Split(strings.Split(line, ": ")[1], " "))
			continue
		}

		if strings.Contains(line, "map:") {
			parsingMap = true
			currentMapName = line
			mapNames = append(mapNames, line)
			currentMapConversionTables = [][]int{}
			continue
		}

		if parsingMap {
			var parsedRanges = parseArrayOfStringsToIntegers(strings.Split(line, " "))
			currentMapConversionTables = append(currentMapConversionTables, parsedRanges)
		}
	}

	if len(currentMapConversionTables) > 0 {
		maps[currentMapName] = currentMapConversionTables
	}

	part1Start := time.Now()
	var lowestLocationNumber, _ = findLowestLocationNumber(seeds, mapNames, maps)
	part1End := time.Since(part1Start)

	part2Start := time.Now()
	var lowestLocationNumberBySeedRange = findLowestLocationNumberBySeedRange(seeds, mapNames, maps)
	part2End := time.Since(part2Start)

	fmt.Printf("Lowest location number %v. Took %s \n", lowestLocationNumber, part1End)
	fmt.Printf("Lowest location number by range %v. Took %s \n", lowestLocationNumberBySeedRange, part2End)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findLowestLocationNumber(seeds []int, mapNames []string, maps map[string][][]int) (int, int) {
	var lowestLocationNumber = 100000000000
	var lowestLocationSeed = 0

	for _, seed := range seeds {
		var currentValue = seed

		for _, currentMapName := range mapNames {
			for _, conversionTable := range maps[currentMapName] {
				var destination = conversionTable[0]
				var source = conversionTable[1]
				var mapRange = conversionTable[2]

				if currentValue >= source && currentValue <= source+mapRange {
					// inside range, let's find out the destination
					var currentValPosition = currentValue - source
					currentValue = destination + currentValPosition
					break
				}
			}
		}

		if currentValue < lowestLocationNumber {
			lowestLocationNumber = currentValue
			lowestLocationSeed = seed
		}
	}

	return lowestLocationNumber, lowestLocationSeed
}

func findLowestLocationNumberBySeedRange(seeds []int, mapNames []string, maps map[string][][]int) int {
	var reverseMapNames = []string{}

	for i := len(mapNames) - 1; i >= 0; i-- {
		reverseMapNames = append(reverseMapNames, mapNames[i])
	}

	var locationNumber = 0

	for true {
		var currentValue = locationNumber

		for _, currentMapName := range reverseMapNames {
			for _, conversionTable := range maps[currentMapName] {
				// Reverse destination and source
				var destination = conversionTable[1]
				var source = conversionTable[0]
				var mapRange = conversionTable[2]

				if currentValue >= source && currentValue <= source+mapRange {
					var currentValPosition = currentValue - source
					currentValue = destination + currentValPosition
					break
				}
			}
		}

		if inSeedsRange(currentValue, seeds) {
			return locationNumber
		}

		locationNumber += 1
	}

	return 0
}

func inSeedsRange(value int, seeds []int) bool {
	for i, seed := range seeds {
		if i%2 != 0 {
			continue
		}

		if value >= seed && value <= seed+seeds[i+1] {
			return true
		}
	}

	return false
}

func parseArrayOfStringsToIntegers(array []string) []int {
	var intArray []int

	for _, c := range array {
		var intValue = tryConvertStrToInt(c)
		intArray = append(intArray, intValue)
	}

	return intArray
}

func tryConvertStrToInt(value string) int {
	if intValue, error := strconv.Atoi(value); error == nil {
		return intValue
	}

	return -1
}
