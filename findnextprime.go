package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	} else {
		if IsPrime(nb) {
			return nb
		}
		return FindNextPrime(nb + 2)
	}
}
