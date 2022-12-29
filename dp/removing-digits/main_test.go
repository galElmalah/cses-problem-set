package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	input int
	large bool
}{
	{input: 27, large: false},
	{input: 30, large: false},
	{input: 15, large: false},
	{input: 4434, large: true},
	{input: 37882, large: true},
}

func TestSum(t *testing.T) {

	t.Run("Regular recursive solution", func(t *testing.T) {

		got := removingDigitsRecursive(27)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Memoized recursive solution", func(t *testing.T) {

		got := removingDigitsMemoized(4434, &map[int]int{})
		want := 687

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Bottom up solution", func(t *testing.T) {

		got := removingDigitsBottomUp(4434)
		want := 687

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Greedy solution", func(t *testing.T) {

		got := removingDigitsGreedy(4434)
		want := 687

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

func BenchmarkRecursive(b *testing.B) {
	for _, v := range table {
		if !v.large {
			b.Run(fmt.Sprintf("input_%d", v.input), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					removingDigitsRecursive(v.input)
				}
			})
		}
	}
}

func BenchmarkMemoized(b *testing.B) {
	for _, v := range table {

		b.Run(fmt.Sprintf("input_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				removingDigitsMemoized(v.input, &map[int]int{})
			}
		})
	}

}

func BenchmarkBottomUp(b *testing.B) {
	for _, v := range table {

		b.Run(fmt.Sprintf("input_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				removingDigitsBottomUp(v.input)
			}
		})

	}
}

func BenchmarkGreedy(b *testing.B) {
	for _, v := range table {

		b.Run(fmt.Sprintf("input_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				removingDigitsGreedy(v.input)
			}
		})

	}
}
