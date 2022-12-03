package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	decisionMap := map[rune]map[rune]rune{
		'A': {
			'X': 'Z',
			'Y': 'X',
			'Z': 'Y',
		},
		'B': {
			'X': 'X',
			'Y': 'Y',
			'Z': 'Z',
		},
		'C': {
			'X': 'Y',
			'Y': 'Z',
			'Z': 'X',
		},
	}

	pointsMap := map[rune]int{
		'X': 0,
		'Y': 3,
		'Z': 6,
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
		var other, outcome rune
		fmt.Sscanf(scanner.Text(), "%c %c", &other, &outcome)
		choice := decisionMap[other][outcome]
		score += choiceMap[choice]
		score += pointsMap[outcome]
	}
	fmt.Println(score)
}
