package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	arr := []int(nil)
	for n > 0 {
		arr = append(arr, n%10)
		n = n / 10
	}
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	for _, v := range arr {
		z01.PrintRune(rune(v + '0'))
	}
}
