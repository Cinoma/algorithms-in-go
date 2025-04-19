package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	var rev string
	for i := len(word) - 1; i >= 0; i-- {
		rev += string(word[i])
	}
	return rev

	// Alt Strings Builder way for bigger strings
	// var sb strings.Builder
	// for i :=  len(word) - 1; i>= 0; i--{
	// 	sb.WriteByte(word[i])
	// }
	// return sb.String()

	// Alt Create new array with the letters all pushed over by one
	// var rev string
	// for i := range len(word) {
	// 	rev = string(word[i]) + rev
	// }
	// return rev

	// Alt rune method that accounts for unique characters in other countries
	// var rev string
	// for _, rune := range word {
	// 	rev = string(rune) + rev
	// }
	// return rev
}
