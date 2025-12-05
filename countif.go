package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, word := range tab {
		if !f(string(word)) {
			continue
		} else {
			count++
		}
	}
	return count
}
