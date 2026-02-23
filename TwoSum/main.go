package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{16, 4, 23, 8, 15, 42, 1, 2}
	target := 19

	sort.Ints(arr)

	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			fmt.Println(arr[left], arr[right])
			return
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	fmt.Println("No match found")
}
