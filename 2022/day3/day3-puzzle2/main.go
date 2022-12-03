package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func priority(r rune) int {
	if r > 'Z' {
		return int(r - 'a' + 1)
	} else {
		return int(r - 'A' + 27)
	}
}

func findCommon(a, b, c []rune) (int32, []rune, []rune, []rune) {
	if a[0] == b[0] && b[0] == c[0] {
		return a[0], []rune{}, []rune{}, []rune{}
	}
	if a[0] <= b[0] && a[0] <= c[0] {
		return findCommon(a[1:], b, c)
	} else if b[0] <= a[0] && b[0] <= c[0] {
		return findCommon(a, b[1:], c)
	} else {
		return findCommon(a, b, c[1:])
	}
}

func main() {
	f, _ := os.Open("2022/day3/day3-puzzle2/input.txt")
	scanner := bufio.NewScanner(f)
	priorities := 0
	for scanner.Scan() {
		first := []rune(scanner.Text())
		scanner.Scan()
		second := []rune(scanner.Text())
		scanner.Scan()
		third := []rune(scanner.Text())

		sort.Slice(first, func(i, j int) bool {
			return first[i] < first[j]
		})
		sort.Slice(second, func(i, j int) bool {
			return second[i] < second[j]
		})
		sort.Slice(third, func(i, j int) bool {
			return third[i] < third[j]
		})
		badge, _, _, _ := findCommon(first, second, third)
		priorities += priority(badge)
	}
	fmt.Println(priorities)
}
