package slices

// Sum returns the sum of the elements of an int slice
func Sum(s []int) int {
	result := 0
	for _, v := range s {
		result += v
	}
	return result
}
