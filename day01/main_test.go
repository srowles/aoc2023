package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineParse(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		input    string
		expected int64
	}{
		"ex1": {
			input:    "1abc2",
			expected: 12,
		},
		"ex2": {
			input:    "pqr3stu8vwx",
			expected: 38,
		},
		"ex3": {
			input:    "a1b2c3d4e5f",
			expected: 15,
		},
		"ex4": {
			input:    "treb7uchet",
			expected: 77,
		},
		"first2": {
			input:    "12dfjvbasjdfgasdhfgewjhfad",
			expected: 12,
		},
		"last2": {
			input:    "asvjbakjdsfbsahdfg12",
			expected: 12,
		},
		"3digits": {
			input:    "716",
			expected: 76,
		},
		"2digits": {
			input:    "99",
			expected: 99,
		},
		"only1": {
			input:    "9",
			expected: 99,
		},
		"last1only": {
			input:    "cvbkasdfkasdfbahsdbfhxcavhjvaskdjfhgasdjfg8",
			expected: 88,
		},
		"first1only": {
			input:    "7cvbkasdfkasdfbahsdbfhxcavhjvaskdjfhgasdjfg",
			expected: 77,
		},
	}

	// Testing
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := calibrationValue(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestCalibrationWithWords(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		input    string
		expected int64
	}{
		"ex1": {
			input:    "two1nine",
			expected: 29,
		},
		"ex2": {
			input:    "eightwothree",
			expected: 83,
		},
		"ex3": {
			input:    "abcone2threexyz",
			expected: 13,
		},
		"ex4": {
			input:    "xtwone3four",
			expected: 24,
		},
		"ex5": {
			input:    "4nineeightseven2",
			expected: 42,
		},
		"ex6": {
			input:    "zoneight234",
			expected: 14,
		},
		"ex7": {
			input:    "7pqrstsixteen",
			expected: 76,
		},
	}

	// Testing
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := calibrationValueWithWords(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSum(t *testing.T) {
	actual := part1(`1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
	`)
	assert.Equal(t, int64(142), actual)
}
