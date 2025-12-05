package piscine

func Map(f func(int) bool, a []int) []bool {
	out := make([]bool, len(a))
	for i, num := range a {
		out[i] = f(num)
	}
	return out
}
