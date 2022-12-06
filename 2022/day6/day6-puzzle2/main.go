package main

import (
	"bufio"
	"fmt"
	"os"
)

type indices []int

func hash(r rune) byte {
	return (byte(r) - byte('a')) % 13
}

func isAllDiff(a []rune) bool {
	bins := make([]uint, len(a)-1)
	for i := 0; i < len(a); i++ {
		bins[hash(a[i])] |= 1 << i
	}
	for _, bin := range bins {
		switch bin {
		case 0b0001, 0b0010, 0b0100, 0b1000, 0b0001_0000, 0b0010_0000, 0b0100_0000, 0b1000_0000, 0b0001_0000_0000, 0b0010_0000_0000, 0b0100_0000_0000, 0b1000_0000_0000:
			continue // only one position has hashed value
		default:
			indicesToCheck := make([]int, 0)
			for i := 0; i < len(a); i++ {
				if bin&(1<<i) != 0 {
					indicesToCheck = append(indicesToCheck, i)
				}
			}
			for x := 0; x < len(indicesToCheck); x++ {
				for y := x + 1; y < len(indicesToCheck); y++ {
					if a[indicesToCheck[x]] == a[indicesToCheck[y]] {
						return false
					}
				}
			}
		}
	}
	return true
}

func findDifferent(a []rune) int {
	for i := 0; i < len(a)-14; i++ {
		if isAllDiff(a[i : i+14]) {
			return i + 14
		}
	}
	return -1
}

func main() {
	f, _ := os.Open("2022/day6/input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(findDifferent([]rune(scanner.Text())))
	}
}
