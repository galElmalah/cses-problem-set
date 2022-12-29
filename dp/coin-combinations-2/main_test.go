package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	coins  []int
	target int
	large  bool
}{
	{coins: []int{1, 5, 7}, target: 15, large: false},
	{coins: []int{1, 5, 7}, target: 25, large: false},
	{coins: []int{1, 5, 7}, target: 30, large: false},
	{coins: []int{1, 5, 7}, target: 100, large: true},
	{coins: []int{1, 5, 7}, target: 120, large: true},
	{coins: []int{27, 69, 68, 13, 1, 63, 28, 44, 45, 67, 57, 11, 8, 85, 42, 20, 32, 77, 39, 52, 70, 24, 4, 79, 7, 15, 54, 88, 51, 73, 89, 66, 48, 56, 47, 18, 2, 30, 21, 49, 74, 9, 99, 83, 55, 95, 62, 90, 22, 31, 71, 98, 43, 75, 25, 16, 12, 64, 61, 38, 40, 26, 3, 96, 41, 86, 5, 14, 91, 33, 78, 50, 23, 84, 94, 36, 46, 97, 53, 81, 65, 34, 93, 59, 6, 35, 72, 10, 82, 60, 19, 92, 29, 87, 76, 100, 80, 17, 58, 37}, target: 1000000, large: true},
}

func TestSum(t *testing.T) {

	t.Run("Regular recursive solution", func(t *testing.T) {

		got := coinCombinationsRecursive([]int{2, 3, 5}, 9)
		want := 8

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Memoized recursive solution", func(t *testing.T) {

		got := coinCombinationsMemoized([]int{2, 3, 5}, 9, &map[int]int{})
		want := 8

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Bottom up solution", func(t *testing.T) {

		got := coinCombinationsBottomUp([]int{2, 3, 5}, 9)
		want := 8

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Bottom up solution - large input", func(t *testing.T) {
		largeInput := table[len(table)-1]
		got := coinCombinationsBottomUp(largeInput.coins, largeInput.target)
		want := 851260131

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func BenchmarkRecursive(b *testing.B) {
	for _, v := range table {
		if !v.large {
			b.Run(fmt.Sprintf("target_size_%d_number_of_coins_%d", v.target, len(v.coins)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					coinCombinationsRecursive(v.coins, v.target)
				}
			})
		}
	}
}

func BenchmarkMemoized(b *testing.B) {
	for _, v := range table {

		b.Run(fmt.Sprintf("target_size_%d_number_of_coins_%d", v.target, len(v.coins)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				coinCombinationsMemoized(v.coins, v.target, &map[int]int{})
			}
		})
	}

}

func BenchmarkBottomUp(b *testing.B) {
	for _, v := range table {

		b.Run(fmt.Sprintf("target_size_%d_number_of_coins_%d", v.target, len(v.coins)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				coinCombinationsBottomUp(v.coins, v.target)
			}
		})

	}
}
