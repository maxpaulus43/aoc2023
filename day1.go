package main

import (
	_ "embed"
	"fmt"
	_ "fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/maxpaulus43/aoc2023/collections"
)

//go:embed inputs/day1.txt
var input string

func day1Part1() int {
	lines := strings.Split(input, "\n")
	values := make([]int, 0, len(lines))
	for _, l := range lines {
		lenL := len(l)
		if lenL == 0 {
			continue
		}
		_, first := firstDigitIdx(l)
		_, last := lastDigitIdx(l)
		v, _ := strconv.Atoi(first + last)
		values = append(values, v)
	}
	return collections.Sum(values)
}

func day1Part2() int {
	lines := strings.Split(input, "\n")
	values := make([]int, 0, len(lines))
	for _, l := range lines {
		lenL := len(l)
		if lenL == 0 {
			continue
		}
		v, _ := strconv.Atoi(firstDigit2(l))
		values = append(values, v)
	}
	return collections.Sum(values)
}

func firstDigitIdx(s string) (int, string) {
	for i, c := range s {
		if unicode.IsNumber(c) {
			return i, string(c)
		}
	}
	return -1, ""
}
func lastDigitIdx(s string) (int, string) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		c := rune(s[i])
		if unicode.IsNumber(c) {
			return i, string(c)
		}
	}
	return -1, ""
}

func firstDigit2(s string) string {
	fi, f := firstDigitIdx(s)
	li, l := lastDigitIdx(s)

	digitStrs := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, str := range digitStrs {
		idx := strings.Index(s, str)
		lIdx := strings.LastIndex(s, str)
		if idx == -1 {
			continue
		}
		if idx < fi || fi == -1 {
			f = string('0' + i)
			fi = idx
		}
		if lIdx > li || li == -1 {
			l = string('0' + i)
			li = lIdx
		}
	}

	fmt.Printf("%v: %v %v\n", s, f, l)

	return f + l
}
