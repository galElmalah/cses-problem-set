package main

import (
	"fmt"

	"github.com/galElmalah/cses-problem-set/util"
)

func coinCombinationsRecursive(coins []int, t int) int {
	if t < 0 {
		return 0
	}
	if t == 0 {
		return 1
	}
	sum := 0
	for _, c := range coins {
		sum = util.AddModulo(sum, coinCombinationsRecursive(coins, t-c))
	}

	return sum

}

func coinCombinationsMemoized(coins []int, t int, m *map[int]int) int {
	if v, ok := (*m)[t]; ok {
		return v
	}
	if t < 0 {
		return 0
	}
	if t == 0 {
		return 1
	}

	sum := 0
	for _, c := range coins {
		sum = util.AddModulo(sum, coinCombinationsMemoized(coins, t-c, m))
	}

	(*m)[t] = sum
	return sum
}

func coinCombinationsBottomUp(coins []int, t int) int {

	dp := make([]int, t+1)

	dp[0] = 1
	for i := 0; i <= t; i++ {
		for _, c := range coins {
			if i >= c {
				dp[i] = util.AddModulo(dp[i], dp[i-c])
			}
		}
	}

	return dp[t]
}

func main() {
	fmt.Println(coinCombinationsRecursive([]int{2, 3, 5}, 9))
	fmt.Println(coinCombinationsMemoized([]int{1, 1500, 1000}, 2000, &map[int]int{}))
	fmt.Println(coinCombinationsBottomUp([]int{27, 69, 68, 13, 1, 63, 28, 44, 45, 67, 57, 11, 8, 85, 42, 20, 32, 77, 39, 52, 70, 24, 4, 79, 7, 15, 54, 88, 51, 73, 89, 66, 48, 56, 47, 18, 2, 30, 21, 49, 74, 9, 99, 83, 55, 95, 62, 90, 22, 31, 71, 98, 43, 75, 25, 16, 12, 64, 61, 38, 40, 26, 3, 96, 41, 86, 5, 14, 91, 33, 78, 50, 23, 84, 94, 36, 46, 97, 53, 81, 65, 34, 93, 59, 6, 35, 72, 10, 82, 60, 19, 92, 29, 87, 76, 100, 80, 17, 58, 37}, 1000000))
}
