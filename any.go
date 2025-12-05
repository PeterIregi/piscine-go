package piscine

func Any(f func(string) bool, a []string) bool {
	for _, word := range a {
		if !f(word) {
			continue
		} else {
			return true
		}
	}
	return false
}
