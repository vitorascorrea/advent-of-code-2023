package main

import (
	"strings"
	"testing"
)

func TestComputeSameStrengthHandsNoJoker(t *testing.T) {
	var expected = true
	var actual = computeSameStrengthHandsTieBreaker(hand("22345"), hand("33A24"), false)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestComputeSameStrengthHandsNoJokerDifferentPosition(t *testing.T) {
	var expected = true
	var actual = computeSameStrengthHandsTieBreaker(hand("TKKKK"), hand("QQQQJ"), false)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestComputeSameStrengthHandsWithJoker(t *testing.T) {
	var expected = true
	var actual = computeSameStrengthHandsTieBreaker(hand("JKKKK"), hand("2222J"), true)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestComputeHandStrengthHighestCardNoJoker(t *testing.T) {
	var expected = HANDS_VALUE["HIGHEST_CARD"]
	var actual = computeHandStrength("A2345", false)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestComputeHandStrengthHighestCardWithJoker(t *testing.T) {
	var expected = HANDS_VALUE["ONE_PAIR"]
	var actual = computeHandStrength("J2345", true)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func hand(str string) []string {
	return strings.Split(str, "")
}
