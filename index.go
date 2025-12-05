package piscine

func Index(s string, toFind string) int {
	if len(s) == 0 || len(toFind) == 0 {
		return -1
	}
	if len(toFind) > len(s) {
		return -1
	} else {
		for i := 0; i <= len(s)-len(toFind); i++ {
			if s[i:i+len(toFind)] == toFind {
				return i
			}
		}
		return -1
	}
}
