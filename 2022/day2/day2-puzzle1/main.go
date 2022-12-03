package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scoreMap := map[rune]map[rune]int{
		'A': {
			'X': 3,
			'Y': 6,
			'Z': 0,
		},
		'B': {
			'X': 0,
			'Y': 3,
			'Z': 6,
		},
		'C': {
			'X': 6,
			'Y': 0,
			'Z': 3,
		},
	}

	choiceMap := map[rune]int{
		'X': 1,
		'Y': 2,
		'Z': 3,
	}

	f, _ := os.Open("2022/day2/input.txt")
	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		var other, me rune
		fmt.Sscanf(scanner.Text(), "%c %c", &other, &me)
		score += scoreMap[other][me]
		score += choiceMap[me]
	}
	fmt.Println(score)
}
