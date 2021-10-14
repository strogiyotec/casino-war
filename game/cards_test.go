package game

import (
	"testing"
)

func TestNewDesk(t *testing.T) {
	desk := NewDesk()
	if len(desk.cards) != 52 {
		t.Fatal("Desk size must be 52")
	}
}

//Discard two cards
func TestDesk_Next(t *testing.T) {
	desk := NewDesk()
	desk.Next()
	desk.Next()
	if len(desk.cards) != 50 {
		t.Fatal("Desk size must be 50")
	}
}
