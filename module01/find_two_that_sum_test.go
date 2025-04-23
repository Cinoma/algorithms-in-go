package module01

import (
	"fmt"
	"testing"
)

func TestFindTwoThatSum(t *testing.T) {
	tests := []struct {
		numbers        []int
		sum            int
		expectedI      int
		expectedJ      int
		expectSolution bool
	}{
		{[]int{1, 2, 3, 4}, 7, 2, 3, true},
		{[]int{0, -1, 1}, 0, 0, 2, true},
		{[]int{0, 1, 1}, 0, -1, -1, false},
		{[]int{10, 1, 12, 3, 7, 2, 2, 1}, 4, 1, 7, true},
		{[]int{10, 1, 12, 3, 7, 2, 2,
			1}, 20, -1, -1, false},
		{[]int{}, 5, -1, -1, false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v with sum %v", tc.numbers, tc.sum), func(t *testing.T) {
			numbersCopy := append([]int{}, tc.numbers...)
			i, j := FindTwoThatSum(numbersCopy, tc.sum)

			if tc.expectSolution {
				if i == -1 || j == -1 {
					t.Fatalf("Expected a solution for %v summing to %v, but no solution was found", tc.numbers, tc.sum)
				}
				//only run these checks if there is a solution
				if i < 0 || j < 0 || i >= len(tc.numbers) || j >= len(tc.numbers) {
					t.Fatalf("Indices out of bounds: i = %d, j = %d, length of array %d", i, j, len(tc.numbers))
				}

				if numbersCopy[i]+numbersCopy[j] != tc.sum {
					t.Fatalf("Wrong solution found. Expected sum of %v, but got %v + %v = %v with indices %d and %d",
						tc.sum, numbersCopy[i], numbersCopy[j], numbersCopy[i]+numbersCopy[j], i, j)
				}
			} else {
				if i != -1 || j != -1 {
					t.Fatalf("Did not expect a solution for %v summing to %v, but a solution was found at indices %d and %d", tc.numbers, tc.sum, i, j)
				}
			}
		})
	}
}

func intSlicesEqual(got, want []int) error {
	if len(got) != len(want) {
		return fmt.Errorf("len(got) = %v; len(want) = %v", len(got), len(want))
	}
	for i, gv := range got {
		wv := want[i]
		if gv != wv {
			return fmt.Errorf("got[%v] = %v; want[%v] = %v", i, gv, i, wv)
		}
	}
	return nil
}
