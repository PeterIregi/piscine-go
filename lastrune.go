package piscine

func LastRune(s string) rune {
	sr := []rune(s)
	return sr[(len(s) - 1)]
}
