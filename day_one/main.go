package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var spelledOutNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	f, err := os.Open("day_one/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var totalSumDigitOnly int = 0
	var totalSumSpelledInclusive int = 0

	for scanner.Scan() {
		var line = scanner.Text()

		totalSumDigitOnly += findDoubleDigitNumber(line, false)
		totalSumSpelledInclusive += findDoubleDigitNumber(line, true)
	}

	fmt.Printf("Total sum (digits only): %v", totalSumDigitOnly)
	fmt.Print("\n")
	fmt.Printf("Total sum (spelled inclusive): %v", totalSumSpelledInclusive)
	fmt.Print("\n")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findDoubleDigitNumber(line string, checkForSpelledOutNumber bool) int {
	var leftmostRegexExpression = regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|[1-9])")
	var rightmostRegexExpression = regexp.MustCompile("(.*)(one|two|three|four|five|six|seven|eight|nine|[1-9])")

	if checkForSpelledOutNumber == false {
		leftmostRegexExpression = regexp.MustCompile("[1-9]")
		rightmostRegexExpression = regexp.MustCompile("(.*)([1-9])")
	}

	var firstNumberString string = leftmostRegexExpression.FindString(line)
	var lastNumberStringMatchs []string = rightmostRegexExpression.FindStringSubmatch(line)
	var lastNumberString string = lastNumberStringMatchs[len(lastNumberStringMatchs)-1]

	if checkForSpelledOutNumber == false {
		var doubleDigitString = firstNumberString + lastNumberString
		if value, error := strconv.Atoi(doubleDigitString); error == nil {
			return value
		}
	}

	var firstSpelledNumber = spelledOutNumbers[firstNumberString]
	var lastSpelledNumber = spelledOutNumbers[lastNumberString]

	if firstSpelledNumber == 0 {
		if firstValue, error := strconv.Atoi(firstNumberString); error == nil {
			firstSpelledNumber = firstValue
		}
	}

	if lastSpelledNumber == 0 {
		if lastValue, error := strconv.Atoi(lastNumberString); error == nil {
			lastSpelledNumber = lastValue
		}
	}

	var result = firstSpelledNumber*10 + lastSpelledNumber

	return result
}
