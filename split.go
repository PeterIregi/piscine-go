package piscine

func Split(s, sep string) []string {
	sepLen := len(sep)
	if sepLen == 0 {
		return []string{s}
	}
	count := 1
	for i := 0; i+sepLen <= len(s); {
		if s[i:i+sepLen] == sep {
			count++
			i += sepLen
		} else {
			i++
		}
	}
	result := make([]string, 0, count)
	start := 0
	for i := 0; i+sepLen <= len(s); {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			i += sepLen
			start = i
		} else {
			i++
		}
	}
	result = append(result, s[start:])

	return result
}
