package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	targetMap := map[byte]int{}
	resultMap := map[byte]int{}
	leftPointer := 0
	rightPointer := 0
	minLength := 0
	substr := ""

	for _, x := range t {
		targetMap[byte(x)]++
	}

	for rightPointer < len(s) && leftPointer < len(s) {
		resultMap[s[rightPointer]]++
		rightPointer++
		for compareMaps(targetMap, resultMap) {
			substr = s[leftPointer:rightPointer]
			if minLength > len(substr) {
				minLength = len(substr)
			}
			resultMap[s[leftPointer]]--
			leftPointer++
		}
	}
	return substr
}

func compareMaps(targetMap map[byte]int, resultMap map[byte]int) bool {
	for key, val := range targetMap {
		if resultMap[key] < val {
			return false
		}
	}
	return true
}
