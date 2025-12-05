package piscine

func NRune(s string, n int) rune {
	sr := []rune(s)
	if n < 0 || n > len(sr) {
		return 0
	} else if n == 0 {
		return 0
	}
	return sr[n-1]
}
