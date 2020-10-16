package helper

func IntSlice(val, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = val
	}
	return s
}
