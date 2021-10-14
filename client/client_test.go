package client

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCheepsPrompt(t *testing.T) {
	config := UserConfig{TotalChips: 50}
	//read 42
	r := ioutil.NopCloser(strings.NewReader("42\n"))
	defer r.Close()
	chips, err := ChipsPrompt(os.Stdin, r, config)
	if err != nil {
		t.Fatal(err)
	}
	if chips.Amount != 42 {
		t.Fatalf("Wrong amount, expected 42 got %d", chips.Amount)
	}

}

//Test that user chose to Surrender
func TestContinueBetPrompt(t *testing.T) {
	//read 42
	r := ioutil.NopCloser(strings.NewReader("Surrender\n"))
	defer r.Close()
	toBet, err := ContinueBetPrompt(os.Stdin, r)
	if err != nil {
		t.Fatal(err)
	}
	if toBet {
		t.Fatal("User didn't want to bet")
	}
}
