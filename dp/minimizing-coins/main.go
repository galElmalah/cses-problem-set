package main

import (
	"fmt"
	"math"

	"github.com/galElmalah/cses-problem-set/util"
)

func minimizeCoinsRecursive(coins []int, t int) int {
	if t == 0 {
		return 0
	}
	min := math.MaxInt
	for _, c := range coins {
		if t >= c {
			min = util.Min(min,
				1+minimizeCoinsRecursive(coins, t-c))
		}

	}
	return min

}

func minimizeCoinsMemoize(coins []int, t int, m *map[int]int) int {
	if v, ok := (*m)[t]; ok {
		return v
	}
	if t == 0 {
		return 0
	}
	min := math.MaxInt
	for _, c := range coins {
		if t >= c {
			min = util.Min(min,
				1+minimizeCoinsMemoize(coins, t-c, m))
		}

	}

	(*m)[t] = min
	return min
}

func minimizeCoinsBottomUp(coins []int, t int) int {

	dp := make([]int, t+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	dp[0] = 0

	for i := 0; i <= t; i++ {
		for _, c := range coins {
			if i >= c {
				dp[i] = util.Min(dp[i], dp[i-c]+1)
			}
		}
	}

	return dp[t]
}

func main() {
	fmt.Println(minimizeCoinsRecursive([]int{1, 5, 7}, 15))
	fmt.Println(minimizeCoinsMemoize([]int{1, 5, 7}, 15, &map[int]int{}))
	fmt.Println(minimizeCoinsBottomUp([]int{1}, 15))
}
