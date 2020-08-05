package main

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestGame(t *testing.T) {

	output = new(bytes.Buffer)

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
