package main

import (
	"bufio"
	"fmt"
	"os"
)

func hash(r rune) byte {
	return (byte(r) - byte('a')) % 3
}

func isAllDiff(a [4]rune) bool {
	bins := make([]uint, 4)
	for i := 0; i < len(a); i++ {
		bins[hash(a[i])] |= 1 << i
	}
	for _, bin := range bins {
		switch bin {
		case 0b0001, 0b0010, 0b0100, 0b1000:
			continue // only one position has hashed value
		default:
			indicesToCheck := make([]int, 0)
			for i := 0; i < 4; i++ {
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
	for i := 0; i < len(a)-4; i++ {
		if isAllDiff([4]rune{a[i], a[i+1], a[i+2], a[i+3]}) {
			return i + 4
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
