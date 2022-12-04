package main

import (
	"advent-of-code/2022/utils"
	"bufio"
	"fmt"
	"os"
)

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
		first := utils.SortAndDedupe([]rune(scanner.Text()))
		scanner.Scan()
		second := utils.SortAndDedupe([]rune(scanner.Text()))
		scanner.Scan()
		third := utils.SortAndDedupe([]rune(scanner.Text()))

		badge, _, _, _ := findCommon(first, second, third)
		priorities += utils.Priority(badge)
	}
	fmt.Println(priorities)
}
