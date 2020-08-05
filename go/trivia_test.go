package main

import (
	"testing"
)

func TestGame(t *testing.T) {
	notAWinner := false

	game := NewGame()

	game.Add("Chet")
	game.Add("Pat")
	game.Add("Sue")

	rand.Seed(42)

	for {
		game.Roll(rand.Intn(5) + 1)

		if rand.Intn(9) == 7 {
			notAWinner = game.WrongAnswer()
		} else {
			notAWinner = game.WasCorrectlyAnswered()
		}

		if !notAWinner {
			break
		}
	}
}
