package module02

import (
	"sort"
)

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for i := range len(list) - 1 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
			BubbleSortInt(list)
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
	for i := range len(list) - 1 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
			BubbleSortString(list)
		}
	}
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
}
