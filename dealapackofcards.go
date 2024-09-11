package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	players := [4][]int{}
	for i, card := range deck {
		playerIndex := i / 3
		players[playerIndex] = append(players[playerIndex], card)
	}
	for i, playerCards := range players {
		playerName := fmt.Sprintf("Player %d:", i+1)
		for _, r := range playerName {
			z01.PrintRune(r)
		}
		z01.PrintRune(' ')

		for j, card := range playerCards {
			if j > 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			for _, r := range fmt.Sprintf("%d", card) {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune('\n')
	}

	z01.PrintRune('\n')
}
