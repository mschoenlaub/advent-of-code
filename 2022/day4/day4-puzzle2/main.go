package main

import (
	"bufio"
	"fmt"
	"os"
)

func overlapRange(a, b, c, d int) bool {
	return c <= b && c >= a
}

func main() {
	f, _ := os.Open("2022/day4/input.txt")
	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		var a, b, c, d int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		if overlapRange(a, b, c, d) || overlapRange(c, d, a, b) {
			score++
		}
	}
	fmt.Println(score)
}
