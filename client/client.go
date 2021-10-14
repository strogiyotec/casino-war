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

func ContinueBetPrompt(writer io.WriteCloser, reader io.ReadCloser) (bool, error) {
	prompt := promptui.Select{
		Label:  "It's a tie do what do you want to do ? ",
		Items:  []string{"Surrender", "Go to War"},
		Stdin:  reader,
		Stdout: writer,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	return result == "Go to War", nil
}

//prompt to get user's bet
//it takes reader and writer for testing purposes
func ChipsPrompt(writer io.WriteCloser, reader io.ReadCloser, config UserConfig) (*Chips, error) {
	//validate input
	validate := func(input string) error {
		val, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			return errors.New("invalid chips Amount")
		}
		if uint32(val) > config.TotalChips {
			return errors.New("you don't have that much chips")
		}
		if val%2 != 0 {
			return errors.New("you have to bet even number of chips")
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
