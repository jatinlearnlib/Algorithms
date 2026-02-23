package main

import "fmt"

func main() {
	arr := []int{10, 20, 5, 3, 6, 1, 2}
	max_profit := 0
	current_profit := 0
	buy := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < buy {
			buy = arr[i]
		} else {
			current_profit = arr[i] - buy
			if current_profit > max_profit {
				max_profit = current_profit
			}
		}

	}
	fmt.Println(max_profit)
}

x := map[int]int{}

x[count] = i