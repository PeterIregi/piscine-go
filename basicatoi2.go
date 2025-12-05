package piscine

func BasicAtoi2(s string) int {
	num := 0
	for _, n := range s {
		if !(n >= '0' && n <= '9') {
			return 0
		}
		num = num*10 + int(n-'0')
	}
	return num
}
