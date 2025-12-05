package piscine

func StrRev(s string) string {
	stringChangable := []byte(s)
	j := len(s) - 1
	for i := 0; i < len(s); i++ {
		stringChangable[i] = s[j]
		j--
	}
	finalString := string(stringChangable)
	return finalString
}
