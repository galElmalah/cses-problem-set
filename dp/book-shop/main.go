package main

import (
	"fmt"

	"github.com/galElmalah/cses-problem-set/util"
)

//https://cses.fi/problemset/task/1158

func main() {
	prices, pages := []int{1, 2, 10, 6, 5, 1, 7, 4, 10, 4}, []int{6, 3, 8, 1, 7, 3, 8, 6, 5, 6}
	fmt.Println(bookShopRecursive(prices, pages, 10, 0))
	fmt.Println(bookShopBottomUp(prices, pages, 10))
}

func bookShopRecursive(prices, pages []int, toSpend int, bookNo int) int {

	if toSpend == 0 {
		return 0
	}

	if bookNo > len(prices)-1 {
		return 0
	}

	if toSpend-prices[bookNo] >= 0 {
		return util.Max(pages[bookNo]+bookShopRecursive(prices, pages, toSpend-prices[bookNo], bookNo+1), bookShopRecursive(prices, pages, toSpend, bookNo+1))
	}

	return bookShopRecursive(prices, pages, toSpend, bookNo+1)
}

func bookShopBottomUp(prices, pages []int, toSpend int) int {
	dp := make([][]int, len(pages))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, toSpend+1)
		for j := 0; j < toSpend+1; j++ {
			dp[i][j] = 0
		}

	}

	for i := 0; i <= toSpend; i++ {
		if i >= prices[0] {
			dp[0][i] = pages[0]
		}
	}

	for i := 1; i < len(pages); i++ {
		for j := 1; j <= toSpend; j++ {
			if dp[i][j] = dp[i-1][j]; j >= prices[i] {
				dp[i][j] = util.Max(dp[i-1][j], dp[i-1][j-prices[i]]+pages[i])
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}

	return 1
}
