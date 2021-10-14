package server

import (
	"bytes"
	"casinowar/client"
	"casinowar/game"
	"strings"
	"testing"
)

func TestServer_Play(t *testing.T) {
	desk := game.FakeDesk()
	gameServer := NewServer(client.Chips{Amount: 20}, desk)
	var stdin bytes.Buffer

	err := gameServer.Play(
		func() game.Card {
			next, err := desk.Next()
			if err != nil {
				t.Fatal(err)
			}
			return next
		},
		&stdin,
		func() bool {
			return false
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	output := stdin.String()
	outputParts := strings.Split(output, "\n")
	if !strings.Contains(outputParts[0], "Dealer got two hearts, User got three hearts") {
		t.Fatalf("Wrong first line of output,it should show two cards,but got '%s' instead", outputParts[0])
	}
	if !strings.Contains(outputParts[1], "user won 40 chips") {
		t.Fatalf("Wrong second line of output,it should show who won,but got '%s' instead", outputParts[1])
	}
}

func TestTie(t *testing.T) {
	desk := game.FakeDesk()
	gameServer := NewServer(client.Chips{Amount: 20}, desk)
	var stdin bytes.Buffer

	err := gameServer.Play(
		func() game.Card {
			//always return the same card
			return game.TestCard
		},
		&stdin,
		func() bool {
			return false
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	output := stdin.String()
	outputParts := strings.Split(output, "\n")
	if !strings.Contains(outputParts[0], "Dealer got two diamonds, User got two diamonds") {
		t.Fatalf("Wrong first line of output,it should show two cards,but got '%s' instead", outputParts[0])
	}
	if !strings.Contains(outputParts[1], "it's a tie") {
		t.Fatalf("Wrong second line of output,it should show the tie,but got '%s' instead", outputParts[1])
	}
}
