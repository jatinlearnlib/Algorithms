package main

import "fmt"

func main() {
	arr := []int{2, 2, 1, 3, 1, 2, 2}
	count := 1
	majority := arr[0]
	for i := 1; i < len(arr); i++ {
		if majority == arr[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majority = arr[i]
			count++
		}
	}
	fmt.Println(majority)
}
