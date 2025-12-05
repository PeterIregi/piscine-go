package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}
	Asc := true
	Desc := true
	for i := 0; i < len(a)-1; i++ {
		comp := f(a[i], a[i+1])
		if comp > 0 {
			Asc = false
		} else if comp < 0 {
			Desc = false
		}
	}
	return Asc || Desc
}
