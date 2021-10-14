package server

import (
	"casinowar/client"
	"casinowar/game"
	"fmt"
	"io"
)

const (
	BURN = 3
)

type Server struct {
	chips client.Chips
	desk  game.Desk
}

func NewServer(chips client.Chips, desk game.Desk) Server {
	return Server{chips: chips, desk: desk}
}

//reader that reads a single card
type Reader func() game.Card

//function that will prompt a user to choose weather he/she wants to "go to war"
type ContinueBet func() bool

//Play game
//takes a reader for testing
func (server Server) Play(cardReader Reader, output io.Writer, toBet ContinueBet) error {
	userCard := cardReader()
	dealerCard := cardReader()
	_, err := output.Write(
		[]byte(
			fmt.Sprintf(
				"Dealer got %s, User got %s\n",
				dealerCard.String(),
				userCard.String(),
			),
		),
	)
	if err != nil {
		return nil
	}
	compare := game.Compare(userCard, dealerCard)
	if compare == 1 {
		//wiki doesn't say how much dealer bets so I assume it's the same amount as user
		_, err := output.Write([]byte(fmt.Sprintf("user won %d chips\n", server.chips.Amount*2)))
		return err
	} else if compare == -1 {
		_, err := output.Write([]byte(fmt.Sprintf("dealer won %d chips\n", server.chips.Amount*2)))
		return err
	} else {
		//tie
		if toBet() {
			return server.tie(cardReader, output)
		} else {
			_, err := output.Write([]byte(fmt.Sprintf("it's a tie user decided to surrender and got half of the bet %d\n", server.chips.Amount/2)))
			return err
		}
	}
}

func (server Server) tie(cardReader Reader, output io.Writer) error {
	err := server.burn(output, cardReader)
	if err != nil {
		return nil
	}
	userCard := cardReader()
	err = server.burn(output, cardReader)
	if err != nil {
		return nil
	}
	dealerCard := cardReader()
	compare := game.Compare(userCard, dealerCard)
	//Wiki is saying "must double his stake" I assume it meant user has to double original bet
	if compare > 1 {
		_, err := output.Write([]byte(fmt.Sprintf("user won %d chips\n", server.chips.Amount*2)))
		return err
	} else if compare < 1 {
		_, err := output.Write([]byte(fmt.Sprintf("dealer won %d chips\n", server.chips.Amount*2)))
		return err
	} else {
		_, err := output.Write([]byte(fmt.Sprintf("it's a tie, user get's his/her money back %d\n", server.chips.Amount)))
		return err
	}
}

//burn cards in case of a tie
func (server Server) burn(output io.Writer, cardReader Reader) error {
	//discard three cards
	for i := 0; i < BURN; i++ {
		card := cardReader()
		_, err := output.Write([]byte(fmt.Sprintf("Card %s was discarded\n", card.String())))
		if err != nil {
			return err
		}
	}
	return nil
}
