package piscine

func LoafOfBread(str string) string {
	s := []rune(str)
	r := []rune{}

	if len(s) > 0 && len(s) < 5 {
		return "Invalid Output\n"
	}

	var stp bool
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			r = append(r, s[i])
			stp = true
			count++
		}
		if count%5 == 0 && stp && i+1 < len(s) && len(r) != 0 {

			hasFutureLoafContent := false
			for j := i + 2; j < len(s); j++ {
				if s[j] != ' ' {
					hasFutureLoafContent = true
					break
				}
			}

			if s[i+1] == ' ' {
				stp = false
			} else {
				s[i+1] = ' '
				stp = false
			}

			if hasFutureLoafContent {
				r = append(r, ' ')
			}
		}
	}
	return string(r) + "\n"
}
