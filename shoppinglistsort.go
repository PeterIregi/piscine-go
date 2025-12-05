package piscine

func ShoppingListSort(slice []string) []string {
	// Bubble Sort based on string length
	n := len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// If the length of slice[j] is greater than the length of slice[j+1], swap them
			if len(slice[j]) > len(slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}
