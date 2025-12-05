package piscine

func Max(a []int) int {
	if len(a) < 1 {
		return 0
	}
	max := 0
	for i, num := range a {
		if i == 0 {
			if num > a[i+1] {
				max = num
			} else {
				max = a[i+1]
			}
		} else {
			if num > max {
				max = num
			} else {
				continue
			}
		}
	}
	return max
}
