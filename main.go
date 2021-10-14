package main

import (
	"casinowar/client"
	"fmt"
	"os"
)

func main() {
	config := client.UserConfig{TotalChips: 50}
	prompt, err := client.CheepsPrompt(os.Stdin, os.Stdout, config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prompt)
}
