package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/srowles/aoc2023"
)

func main() {
	vals := aoc2023.InputFromWebsite("1")
	fmt.Println(part1(vals))
	fmt.Println(part2(vals))
}

func part1(input string) int64 {
	sum := int64(0)
	for _, v := range strings.Split(input, "\n") {
		sum += calibrationValue(v)
	}
	return sum
}

func part2(input string) int64 {
	sum := int64(0)
	for _, v := range strings.Split(input, "\n") {
		sum += calibrationValueWithWords(v)
	}
	return sum
}

func calibrationValue(input string) int64 {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0
	}
	var val int64
	for _, r := range input {
		num := int64(r) - int64('0')
		if num < 0 || num > 9 {
			continue
		}
		val = num
		break
	}
	first := val
	for i := len(input) - 1; i > 0; i-- {
		num := int64(input[i]) - int64('0')
		if num < 0 || num > 9 {
			continue
		}
		val = num
		break
	}
	last := val
	return (first * 10) + last
}

var lookup map[string]int64 = map[string]int64{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
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

type found struct {
	idx int
	val int64
}

func calibrationValueWithWords(input string) int64 {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0
	}
	var start, end []found
	for k, v := range lookup {
		idx := strings.Index(input, k)
		if idx != -1 {
			start = append(start, found{idx: idx, val: v})
		}
	}

	for k, v := range lookup {
		idx := strings.LastIndex(input, k)
		if idx != -1 {
			end = append(end, found{idx: idx, val: v})
		}
	}

	sort.Slice(start, func(i, j int) bool {
		return start[i].idx < start[j].idx
	})
	sort.Slice(end, func(i, j int) bool {
		return end[i].idx > end[j].idx
	})
	return (start[0].val * 10) + end[0].val
}
