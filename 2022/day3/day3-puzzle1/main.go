package main

import (
	"advent-of-code/2022/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("2022/day3/day3-puzzle1/input.txt")
	scanner := bufio.NewScanner(f)
	priorities := 0
	for scanner.Scan() {
		backpack := []rune(scanner.Text())
		left := utils.SortAndDedupe(backpack[0 : len(backpack)/2])
		right := utils.SortAndDedupe(backpack[len(backpack)/2:])

		found := false

		for !found && len(left) > 0 {
			for !found && len(right) > 0 {
				if left[0] == right[0] {
					priorities += utils.Priority(left[0])
					found = true
				}
				if left[0] < right[0] {
					left = left[1:]
				} else {
					right = right[1:]
				}
			}
		}
	}
	fmt.Println(priorities)
}
