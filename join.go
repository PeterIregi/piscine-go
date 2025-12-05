package piscine

func Join(strs []string, sep string) string {
	new := ""
	for i, ch := range strs {
		if i != len(strs)-1 {
			new += ch
			new += sep

		} else {
			new += ch
		}
	}
	return new
}
