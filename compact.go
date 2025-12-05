package piscine

func Compact(ptr *[]string) int {
	if ptr == nil {
		return 0
	}
	s := *ptr
	newSlice := make([]string, 0, len(s))

	for _, v := range s {
		if v != "" {
			newSlice = append(newSlice, v)
		}
	}
	*ptr = newSlice
	return len(newSlice)
}
