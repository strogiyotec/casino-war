package server

import (
	"casinowar/client"
	"casinowar/game"
	"fmt"
	"io"
)

type Server struct {
	chips client.Chips
	desk  game.Desk
}

//reader that reads a single card
type Reader func() game.Card

//Play game
//takes a reader for testing
func (server Server) Play(cardReader Reader, output io.Writer) error {
	userCard := cardReader()
	dealerCard := cardReader()
	compare := game.Compare(userCard, dealerCard)
	if compare > 1 {
		//wiki doesn't say how much dealer bets so I assume it's the same amount as user
		_, err := output.Write([]byte(fmt.Sprintf("user won %d chips", server.chips.Amount*2)))
		return err
	} else if compare < 1 {
		_, err := output.Write([]byte(fmt.Sprintf("dealer won %d chips", server.chips.Amount*2)))
		return err
	} else {
		//tie
		return server.tie(cardReader, output)
	}
}

//TODO: Finish tie
func (server Server) tie(cardReader Reader, output io.Writer) error {
	//discard three cards
	for i := 0; i < 3; i++ {
		card := cardReader()
		_, err := output.Write([]byte(fmt.Sprintf("Card %s  was discarded", card.String())))
		if err != nil {
			return err
		}
	}
	userCard := cardReader()
	dealerCard := cardReader()
	compare := game.Compare(userCard, dealerCard)
	//Wiki is saying "must double his stake" I assume it meant user has to double original bet
	if compare > 1 {
		_, err := output.Write([]byte(fmt.Sprintf("user won %d chips", server.chips.Amount*2)))
		return err
	} else if compare < 1 {
		_, err := output.Write([]byte(fmt.Sprintf("dealer won %d chips", server.chips.Amount*2)))
		return err
	} else {
		//tie
		return server.tie(cardReader, output)
	}
}
