package piscine

func Abort(a, b, c, d, e int) int {
	new := []int{a, b, c, d, e}

	for i := 0; i < len(new); i++ {
		for j := i + 1; j < len(new); j++ {
			if new[j] < new[i] {
				new[i], new[j] = new[j], new[i]
			}
		}
	}

	return new[2]
}
