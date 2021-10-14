package client

import (
	"errors"
	"github.com/manifoldco/promptui"
	"io"
	"strconv"
)

//user's initial setting
type UserConfig struct {
	TotalChips uint32
}

//how much chips user is betting
type Chips struct {
	Amount uint32
}

//prompt to get user's bet
//it takes reader and writer for testing purposes
func CheepsPrompt(writer io.WriteCloser, reader io.ReadCloser,config UserConfig) (*Chips, error) {
	//validate input
	validate := func(input string) error {
		val, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			return errors.New("invalid chips Amount")
		}
		if uint32(val) > config.TotalChips {
			return errors.New("you don't have that much chips")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Amount of chips to bet",
		Validate: validate,
		Stdin:    reader,
		Stdout:   writer,
	}
	result, err := prompt.Run()

	if err != nil {
		return nil, err
	}
	resAsInt, _ := strconv.ParseUint(result, 10, 32)
	return &Chips{Amount: uint32(resAsInt)}, nil
}
