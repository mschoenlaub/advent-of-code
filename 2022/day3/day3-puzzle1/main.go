package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func priority(r rune) int32 {
	if r > 'Z' {
		return r - 'a' + 1
	} else {
		return r - 'A' + 27
	}
}

func main() {
	f, _ := os.Open("2022/day3/day3-puzzle1/input.txt")
	scanner := bufio.NewScanner(f)
	priorities := int32(0)
	for scanner.Scan() {
		backpack := []rune(scanner.Text())
		left := backpack[0 : len(backpack)/2]
		right := backpack[len(backpack)/2:]
		sort.Slice(left, func(i, j int) bool {
			return left[i] < left[j]
		})
		sort.Slice(right, func(i, j int) bool {
			return right[i] < right[j]
		})

		found := false

		for !found && len(left) > 0 {
			for !found && len(right) > 0 {
				if left[0] == right[0] {
					priorities += priority(left[0])
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
