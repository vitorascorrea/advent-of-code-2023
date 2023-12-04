package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day_4/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	var totalSum int

	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
		var _, cardTotal = computeWinningNumbers(line)
		totalSum += cardTotal
	}

	fmt.Printf("Total sum: %v \n", totalSum)
	fmt.Printf("Total scratchcards sum: %v \n", computeTotalScratchcards(lines))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeWinningNumbers(line string) (int, int) {
	var noCardPrefix = strings.Split(line, ": ")[1]
	var numbers = strings.Split(noCardPrefix, " | ")
	var winningNumbers = deleteEmptyValues(strings.Split(numbers[0], " "))
	var elfNumbers = deleteEmptyValues(strings.Split(numbers[1], " "))

	var totalSum = 0
	var totalMatches = 0

	for _, number := range elfNumbers {
		if arrayContains(winningNumbers, number) {
			totalMatches += 1

			if totalSum == 0 {
				totalSum = 1
			} else {
				totalSum *= 2
			}
		}
	}

	return totalMatches, totalSum
}

func arrayContains(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

func deleteEmptyValues(array []string) []string {
	var newArray []string
	for _, str := range array {
		if str != "" {
			newArray = append(newArray, str)
		}
	}
	return newArray
}

func computeTotalScratchcards(cards []string) int {
	var cardsMap = map[string]int{}
	var cardsResultsMap = map[string]int{}

	for _, card := range cards {
		cardsMap[card] = 1
		var totalMatches, _ = computeWinningNumbers(card)
		cardsResultsMap[card] = totalMatches
	}

	for x, card := range cards {
		var count = cardsMap[card]
		for i := 0; i < count; i++ {
			var totalMatches = cardsResultsMap[card]

			if totalMatches == 0 {
				continue
			}

			for j := 1; j <= totalMatches; j++ {
				var cardToDuplicate = cards[x+j]
				cardsMap[cardToDuplicate] += 1
			}
		}
	}

	var totalCards = 0

	for _, count := range cardsMap {
		totalCards += count
	}

	return totalCards
}
