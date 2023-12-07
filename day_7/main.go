package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardValueMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var HANDS_VALUE = map[string]int{
	"FIVE_OF_A_KIND":  7,
	"FOUR_OF_A_KIND":  6,
	"FULL_HOUSE":      5,
	"THREE_OF_A_KIND": 4,
	"TWO_PAIRS":       3,
	"ONE_PAIR":        2,
	"HIGHEST_CARD":    1,
}

func main() {
	f, err := os.Open("day_7/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var hands []string
	var bids []int

	for scanner.Scan() {
		var line = scanner.Text()
		var splitLine = strings.Split(line, " ")
		hands = append(hands, splitLine[0])
		bids = append(bids, tryConvertStrToInt(splitLine[1]))
	}

	var totalWinnings = computeTotalWinnings(hands, bids, false)
	var totalWinningsWithJoker = computeTotalWinnings(hands, bids, true)

	fmt.Printf("Total winnings: %v \n", totalWinnings)
	fmt.Printf("Total winnings with joker: %v \n", totalWinningsWithJoker)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeTotalWinnings(originalHands []string, bids []int, considerJoker bool) int {
	var total = 0
	var hands []string
	var handsToBidsMap = make(map[string]int)

	for i, hand := range originalHands {
		hands = append(hands, hand)
		handsToBidsMap[hand] = bids[i]
	}

	sort.SliceStable(hands, func(i, j int) bool {
		var iValue = computeHandStrength(hands[i], considerJoker)
		var jValue = computeHandStrength(hands[j], considerJoker)

		if iValue == jValue {
			return computeSameStrengthHandsTieBreaker(
				strings.Split(hands[i], ""),
				strings.Split(hands[j], ""),
				considerJoker,
			)
		}

		return iValue < jValue
	})

	for rank, hand := range hands {
		var handBid = handsToBidsMap[hand]
		total += handBid * (rank + 1)
	}

	return total
}

func computeHandStrength(hand string, considerJoker bool) int {
	var cardsMap = make(map[string]int)

	for _, card := range strings.Split(hand, "") {
		cardsMap[card] += 1
	}

	var hasFourOfAKind = false
	var hasThreeOfAKind = false
	var hasOnePair = false
	var hasAnotherPair = false
	var jokerCount = cardsMap["J"]

	for _, value := range cardsMap {
		// Five of a kind
		if value == 5 {
			return HANDS_VALUE["FIVE_OF_A_KIND"]
		}

		// Four of a kind
		if value == 4 {
			hasFourOfAKind = true
			continue
		}

		if value == 3 {
			hasThreeOfAKind = true
			continue
		}

		if value == 2 {
			if hasOnePair {
				hasAnotherPair = true
			} else {
				hasOnePair = true
			}

			continue
		}
	}

	if hasFourOfAKind {
		if considerJoker && jokerCount > 0 {
			// Becomes a five of a kind, joker is the last card left OR is the four of a kind
			return HANDS_VALUE["FIVE_OF_A_KIND"]
		}

		return HANDS_VALUE["FOUR_OF_A_KIND"]
	}

	// Full house
	if hasThreeOfAKind && hasOnePair {
		if considerJoker && jokerCount > 0 {
			// Becomes a five of a kind: joker could be only the three of a kind or the pair
			return HANDS_VALUE["FIVE_OF_A_KIND"]
		}

		return HANDS_VALUE["FULL_HOUSE"]
	}

	// Three of a kind
	if hasThreeOfAKind {
		if considerJoker {
			switch jokerCount {
			case 1:
				// four of a kind. The joker can be a lone card or it can be the three of a kind
				return HANDS_VALUE["FOUR_OF_A_KIND"]
			case 3:
				// four of a kind. The joker can be a lone card or it can be the three of a kind
				return HANDS_VALUE["FOUR_OF_A_KIND"]
			}
		}

		return HANDS_VALUE["THREE_OF_A_KIND"]
	}

	// Two pairs
	if hasOnePair && hasAnotherPair {
		if considerJoker {
			switch jokerCount {
			case 1:
				// joker is a solo card, turns one of the pairs into a three of a kind
				return HANDS_VALUE["FULL_HOUSE"]
			case 2:
				// joker is one of the pairs, becomes four of a kind
				return HANDS_VALUE["FOUR_OF_A_KIND"]
			}
		}

		return HANDS_VALUE["TWO_PAIRS"]
	}

	// One pair
	if hasOnePair {
		if considerJoker {
			switch jokerCount {
			case 1:
				// joker is a solo card, turns the pair into a three of a kind
				return HANDS_VALUE["THREE_OF_A_KIND"]
			case 2:
				// joker is the pair, becomes three of a kind (since all the other cards are different)
				return HANDS_VALUE["THREE_OF_A_KIND"]
			}
		}

		return HANDS_VALUE["ONE_PAIR"]
	}

	if considerJoker && jokerCount > 0 {
		// if there is one joker, it must form a pair
		return HANDS_VALUE["ONE_PAIR"]
	}

	return HANDS_VALUE["HIGHEST_CARD"]
}

// Will return true if iHand is less valueable than jHand
func computeSameStrengthHandsTieBreaker(iHand []string, jHand []string, considerJoker bool) bool {
	for i := 0; i < len(iHand); i++ {
		var iCard = iHand[i]
		var jCard = jHand[i]

		var iCardValue = cardValueMap[iCard]
		var jCardValue = cardValueMap[jCard]

		if considerJoker && iCard == "J" {
			iCardValue = 1
		}

		if considerJoker && jCard == "J" {
			jCardValue = 1
		}

		if iCardValue == jCardValue {
			continue
		}

		return iCardValue < jCardValue
	}

	return false
}

func tryConvertStrToInt(value string) int {
	if intValue, error := strconv.Atoi(value); error == nil {
		return intValue
	}

	return -1
}
