package piscine

func ToUpper(s string) string {
	var product string
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			product += string(ch - 32)
		} else {
			product += string(ch)
		}
	}
	return product
}
