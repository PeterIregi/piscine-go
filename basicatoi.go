package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, n := range s {
		num = num*10 + int(n-'0')
	}
	return num
}
