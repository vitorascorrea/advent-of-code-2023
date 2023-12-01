package main

import (
	"testing"
)

func TestFindingDoubleDigitNoSpelling(t *testing.T) {
	var expected = 12
	var actual = findDoubleDigitNumber("1abc2", false)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestFindingDoubleDigitNoSpellingRepeated(t *testing.T) {
	var expected = 12
	var actual = findDoubleDigitNumber("11112", false)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestFindingDoubleDigitNoSpellingOnlyDigits(t *testing.T) {
	var expected = 19
	var actual = findDoubleDigitNumber("123456789", false)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestFindingDoubleDigitWithSpelling(t *testing.T) {
	var expected = 15
	var actual = findDoubleDigitNumber("one234five", true)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestFindingDoubleDigitWithSpellingOnlyDigits(t *testing.T) {
	var expected = 15
	var actual = findDoubleDigitNumber("12345", true)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestFindingDoubleDigitWithSpellingOnlySpelling(t *testing.T) {
	var expected = 15
	var actual = findDoubleDigitNumber("onetwothreefourfive", true)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestFindingDoubleDigitWithSpellingRandomCharsBetween(t *testing.T) {
	var expected = 13
	var actual = findDoubleDigitNumber("hbonefourfour8lsnfgvf3", true)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}

func TestFindingDoubleDigitWithSpellingWordsSharingLetters(t *testing.T) {
	var expected = 12
	var actual = findDoubleDigitNumber("oneeightwo", true)

	if expected != actual {
		t.Fatalf("Expected %v Actual %v", expected, actual)
	}
}
