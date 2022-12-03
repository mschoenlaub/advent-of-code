package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	elveSum := 0
	currentMax := 0

	f, _ := os.Open("2022/day1/input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if elveSum > currentMax {
				currentMax = elveSum
			}
			elveSum = 0
			continue
		}
		calories, _ := strconv.Atoi(line)
		elveSum += calories
	}
	fmt.Println(currentMax)
}
