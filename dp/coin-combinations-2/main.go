package main

import (
	"fmt"

	"github.com/galElmalah/cses-problem-set/util"
)

func coinCombinationsBottomUp(coins []int, t int) int {

	dp := make([]int, t+1)

	dp[0] = 1
	for _, c := range coins {
		for i := 1; i <= t; i++ {
			if i >= c {
				dp[i] = util.AddModulo(dp[i], dp[i-c])
			}
		}

	}

	return dp[t]
}

func main() {
	fmt.Println(coinCombinationsBottomUp([]int{1, 3}, 9))
}
