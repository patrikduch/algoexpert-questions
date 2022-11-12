package main

import "sort"

func main() {
	coins := []int{1, 1, 1, 1, 1}

	coinsSum := NonConstructibleChange(coins)
	println("result",coinsSum)
}

func NonConstructibleChange(coins []int) int {

	currentChangeCreated := 0;

	// no coins returns minimum change 1
	if len(coins) == 0 {
		return 1
	}

	sort.Slice(coins, func(i, j int) bool {
		return coins[j] > coins[i]
	})

	for _, coin := range coins {
		if coin > currentChangeCreated +1 {
			return currentChangeCreated +1
		}

		currentChangeCreated += coin
	}

	return currentChangeCreated + 1;
}