package arraynslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func checkSums(got []int, want []int, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestSumAll(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}

		got := SumAll(numbers)
		want := []int{6, 15}
		checkSums(got, want, t)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}

		got := SumAllTails(numbers)
		want := []int{5, 11}

		checkSums(got, want, t)
	})

	t.Run("collection with an empty lists", func(t *testing.T) {
		numbers := [][]int{
			{},
			{1, 2, 3},
		}

		got := SumAllTails(numbers)
		want := []int{
			0,
			5,
		}

		checkSums(got, want, t)
	})

	t.Run("empty collection", func(t *testing.T) {
		numbers := [][]int{}

		got := SumAllTails(numbers)
		var want []int

		checkSums(got, want, t)
	})
}
