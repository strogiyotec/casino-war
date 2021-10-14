package main

import (
	"casinowar/client"
	"casinowar/game"
	"casinowar/server"
	"fmt"
	"os"
)

func main() {
	config := client.UserConfig{TotalChips: 50}
	chips, err := client.ChipsPrompt(os.Stdin, os.Stdout, config)
	if err != nil {
		panic(err)
	}
	desk := game.NewDesk()
	gameServer := server.NewServer(*chips, desk)
	err = gameServer.Play(
		func() game.Card {
			next, err := desk.Next()
			if err != nil {
				panic(err)
			}
			return next
		},
		os.Stdout,
		func() bool {
			toContinue, err := client.ContinueBetPrompt(os.Stdin, os.Stdout)
			if err != nil {
				panic(err)
			}
			return toContinue
		},
	)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Exit game")
	}

}
