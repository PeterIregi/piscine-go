package piscine

func SplitWhiteSpaces(s string) []string {
	var slice []string
	var word string
	for i, ch := range s {
		if ch != ' ' && ch != '\t' && ch != '\n' {
			word += string(ch)
		}
		if ((ch == ' ' || ch == '\t' || ch == '\n') && word != "") || i == len(s)-1 {
			slice = append(slice, word)
			word = ""
		}
	}
	return slice
}
