package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//	Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//
// Examples:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//	Fibonacci(2) => 1
//	Fibonacci(3) => 2
//	Fibonacci(4) => 3
//	Fibonacci(5) => 5
//	Fibonacci(6) => 8
//	Fibonacci(7) => 13
//	Fibonacci(14) => 377

var cache = make(map[int]int)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	}
	res := Fibonacci(n-1) + Fibonacci(n-2)
	cache[n] = res
	return res

	//Iterative solution
	// if n <= 1 {
	// 	return n
	// }

	// nM2 := 0
	// nM1 := 1
	// var cur int
	// for i := 2; i <= n; i++ {
	// 	cur = nM2 + nM1
	// 	nM2 = nM1
	// 	nM1 = cur
	// }
	// return cur

	// Hyperbolic solution
	// if n <= 1 {
	// 	return n
	// }
	// return Fibonacci(n-1) + Fibonacci(n-2)
}
