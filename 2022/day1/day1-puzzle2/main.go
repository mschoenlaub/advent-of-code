package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func heapUp(arr []int, i int) {
	if i == 0 {
		return
	}
	parent := (i - 1) / 2
	if arr[i] > arr[parent] {
		swap(arr, i, parent)
		heapUp(arr, parent)
	}
}

func heapDown(arr []int, i, l int) {
	leftChild := i*2 + 1
	rightChild := i*2 + 2
	largest := i
	if leftChild < l && arr[leftChild] > arr[largest] {
		largest = leftChild
	}
	if rightChild < l && arr[rightChild] > arr[largest] {
		largest = rightChild
	}
	if largest != i {
		swap(arr, i, largest)
		heapDown(arr, largest, l)
	}
}

func pop(arr []int, c int) int {
	swap(arr, 0, len(arr)-c)
	heapDown(arr, 0, len(arr)-c)
	return arr[len(arr)-c]
}

func main() {
	elveSum := 0
	heap := make([]int, 0)

	f, _ := os.Open("2022/day1/input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			heap = append(heap, elveSum)
			heapUp(heap, len(heap)-1)
			elveSum = 0
			continue
		}
		calories, _ := strconv.Atoi(line)
		elveSum += calories
	}
	topThree := 0
	for i := 0; i < 3; i++ {
		topThree += pop(heap, i+1)
	}
	fmt.Println(topThree)
}
