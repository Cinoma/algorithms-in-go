package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	if len(list) == 0 {
		return false
	}
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false

	// Alt Go specific way - potentially less efficient
	// return slices.Contains(list, num)
}
