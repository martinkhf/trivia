package main

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"testing"
)

var golden []byte

func init() {
	var err error

	golden, err = ioutil.ReadFile("golden.txt")
	if err != nil {
		panic(err)
	}
}

func TestGame(t *testing.T) {
	buf := new(bytes.Buffer)
	output = buf

	for i := int64(0); i < 256; i++ {
		notAWinner := false

		game := NewGame()

		game.Add("Chet")
		game.Add("Pat")
		game.Add("Sue")

		rand.Seed(i)

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

	if !bytes.Equal(buf.Bytes(), golden) {
		t.Fatal("output does not match golden")
	}
}
