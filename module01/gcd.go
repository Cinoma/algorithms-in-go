package module01

func GCD(a, b int) int {
	for a != b {
		if b == 0 {
			return a
		}
		a, b = b, a%b
	}
	return 0

	// Alt recursive solution
	// if b == 0 {
	// 	return a
	// }
	// a, b = b, a%b
	// return GCD(a, b)
}
