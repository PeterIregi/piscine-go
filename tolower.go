package piscine

func ToLower(s string) string {
	var product string
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			product += string(ch + 32)
		} else {
			product += string(ch)
		}
	}
	return product
}
