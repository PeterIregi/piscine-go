package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	for p := 0; p < 4; p++ {
		fmt.Printf("Player %d:", p+1)

		for i := 0; i < 3; i++ {
			card := deck[p*3+i]
			if i < 2 {
				fmt.Printf(" %d,", card)
			} else {
				fmt.Printf(" %d", card)
			}
		}
		fmt.Printf("\n")
	}
}
