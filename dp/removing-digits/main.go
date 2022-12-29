package main

import (
	"fmt"
	"math"

	"github.com/galElmalah/cses-problem-set/util"
)

func getDigits(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}
	return digits
}
func removingDigitsRecursive(number int) int {
	if number == 0 {
		return 0
	}
	if number < 0 {
		return math.MaxInt
	}
	min := math.MaxInt
	for _, d := range getDigits(number) {
		if d != 0 {
			min = util.Min(min, 1+removingDigitsRecursive(number-d))
		}
	}

	return min
}

func removingDigitsMemoized(number int, m *map[int]int) int {
	if val, ok := (*m)[number]; ok {
		return val
	}
	if number == 0 {
		return 0
	}
	if number < 0 {
		return math.MaxInt
	}
	min := math.MaxInt
	for _, d := range getDigits(number) {
		if d != 0 {
			min = util.Min(min, 1+removingDigitsMemoized(number-d, m))
		}
	}
	(*m)[number] = min
	return min
}

func removingDigitsBottomUp(number int) int {
	dp := make([]int, number+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt - 10

	}
	dp[0] = 0
	dp[1] = 1
	for i := 0; i <= number; i++ {
		for _, d := range getDigits(i) {

			dp[i] = util.Min(dp[i], dp[i-d]+1)

		}
	}

	return dp[number]
}

func removingDigitsGreedy(number int) int {
	c := 0
	for number != 0 {
		c++
		subtract := 0
		for _, d := range getDigits(number) {
			subtract = util.Max(subtract, d)
		}
		number -= subtract
	}
	return c
}

func main() {
	fmt.Println(removingDigitsRecursive(27))
	fmt.Println(removingDigitsMemoized(4434, &map[int]int{}))
	fmt.Println(removingDigitsBottomUp(4434))
	fmt.Println(removingDigitsGreedy(4434))
}
