package piscine

func DescendAppendRange(max, min int) []int {
	intSlice := []int{}
	if min >= max {
		return intSlice
	}
	for i := max; i > min; i-- {
		intSlice = append(intSlice, i)
	}
	return intSlice
}
