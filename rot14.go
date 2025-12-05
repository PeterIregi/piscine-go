package piscine

func Rot14(s string) string {
	strSlice := []byte(s)
	for i, ch := range s {
		if (ch >= 'a' && ch <= 'l') || (ch >= 'A' && ch <= 'L') {
			strSlice[i] = strSlice[i] + 14
		} else if (ch > 'l' && ch <= 'z') || (ch > 'L' && ch <= 'Z') {
			strSlice[i] = strSlice[i] - 12
		}
	}
	return string(strSlice)
}
