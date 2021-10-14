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
	chips, err := CheepsPrompt(os.Stdin, r, config)
	if err != nil {
		t.Fatal(err)
	}
	if chips.Amount != 42 {
		t.Fatalf("Wrong amount, expected 42 got %d", chips.Amount)
	}

}
