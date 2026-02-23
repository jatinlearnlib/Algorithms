package main

import (
	"fmt"
)

func Intersection(arr1 []int, arr2 []int) ([]int, error) {
	result := make([]int, 0)
	resultMap := make(map[int]bool, 0)
	for _, val := range arr1 {
		resultMap[val] = true
	}

	for _, val := range arr2 {
		if _, ok := resultMap[val]; ok {
			delete(resultMap, val)
			result = append(result, val)
		}
	}

	return result, nil
}

func main() {
	arr1 := []int{4, 9, 5}
	arr2 := []int{9, 4, 9, 8, 4}

	result, _ := Intersection(arr1, arr2)

	fmt.Println(result)
}
