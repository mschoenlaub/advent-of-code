package main

import (
	"bufio"
	"fmt"
	"os"
)

func containsSubRange(a, b, c, d int) bool {
	return a <= c && d <= b
}

func main() {
	f, _ := os.Open("2022/day4/input.txt")
	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		var a, b, c, d int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		if containsSubRange(a, b, c, d) || containsSubRange(c, d, a, b) {
			score++
		}
	}
	fmt.Println(score)
}
