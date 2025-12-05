package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		if i == 1 {
			result = nb
		} else {
			result = result * nb
		}
	}
	return result
}
