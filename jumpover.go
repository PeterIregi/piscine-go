package piscine

func JumpOver(str string) string {
	// If empty or too short, return newline
	if len(str) < 3 {
		return "\n"
	}

	result := ""

	// Every third character → index 2, 5, 8, ...
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	return result + "\n"
}
