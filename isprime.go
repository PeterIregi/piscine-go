package piscine

func IsPrime(nb int) bool {
	if nb < 0 || nb <= 1 {
		return false
	} else if nb == 2 || nb == 3 {
		return true
	} else {
		for i := 2; i < nb; i++ {
			if nb%i != 0 {
				continue
			} else {
				return false
			}
		}
		return true
	}
}
