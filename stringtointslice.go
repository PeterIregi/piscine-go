package piscine

func StringToIntSlice(str string) []int {
	new := []int{}
	if str == "" {
		return []int(nil)
	}
	for _, r := range str {
		if r >= '0' || r <= '9' {
			new = append(new, int(r))
		}
	}
	return new
}
