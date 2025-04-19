package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total

	// Recursive method - unnecesary and much slower for large slices
	// if len(numbers) == 0{
	// 	return 0
	// }
	// return numbers[0] + Sum(numbers[1:])
}
