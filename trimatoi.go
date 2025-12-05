package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var str string
	for _, ch := range s {
		if ch == '-' && len(str) == 0 {
			str += "-"
			continue
		}
		if ch >= '0' && ch <= '9' {
			str += string(ch)
		}
	}
	if len(str) == 0 || str == "-" {
		return 0
	}

	if str[0] == '-' {
		return -(BasicAtoi2(str[1:]))
	}
	return BasicAtoi2(str)
}
