package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Colleaction of 5 numers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		went := 15

		if got != went {
			t.Errorf("got %d went %d given %v", got, went, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	went := []int{3, 9}

	if !reflect.DeepEqual(got, went) {
		t.Errorf("got %d went %v", got, went)
	}
}

func TestSumAllTails(t *testing.T) {

	checksums := func(t testing.TB, got, went []int) {
		t.Helper()
		if !reflect.DeepEqual(got, went) {
			t.Errorf("got %d went %v", got, went)
		}
	}

	t.Run("make sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		went := []int{2, 9}

		checksums(t, got, went)
	})

	t.Run("safly sum empty of slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		went := []int{0, 9}

		checksums(t, got, went)
	})
}
