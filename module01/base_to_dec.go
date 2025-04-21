package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	var res float64
	var pow int
	const charset = "0123456789ABCDEF"
	for i := len(value) - 1; i >= 0; i-- {
		v := value[i]
		char := strings.Index(charset, string(v))
		res += float64(char) * math.Pow(float64(base), float64(pow))
		pow++
	}
	return int(res)
}
