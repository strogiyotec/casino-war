package game

import (
	"errors"
	"fmt"
	"math/rand"
)

//card names
const (
	TWO   string = "two"
	THREE string = "three"
	FOUR  string = "four"
	FIVE  string = "five"
	SIX   string = "six"
	SEVEN string = "seven"
	EIGHT string = "eight"
	NINE  string = "nine"
	TEN   string = "ten"
	JACK  string = "jack"
	QUEEN string = "queen"
	KING  string = "king"
	ACE   string = "ace"
)

//colours
const (
	CLUBS    string = "clubs"
	DIAMONDS string = "diamonds"
	HEARTS   string = "hearts"
	SPADES   string = "spades"
)

type Card struct {
	color string //color of this card
	rank  Rank   //what is the rank of this card
}

func (c Card) String() string {
	return fmt.Sprintf("%s %s", c.rank.name, c.color)
}

//Rank representation of a card
type Rank struct {
	name  string //name of of a card Ex:aces
	order uint   //the order of this card, aces have highest
}

//the desk with cards
type Desk struct {
	cards []Card
}

//compare two cards
func Compare(first, second Card) int {
	if first.rank.order < second.rank.order {
		return 0
	}
	if first.rank.order > second.rank.order {
		return 1
	}
	return 0
}

func NewDesk() Desk {
	var cards []Card
	//twos
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: TWO, order: uint(2)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: TWO, order: uint(2)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: TWO, order: uint(2)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: TWO, order: uint(2)}})
	//threes
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: THREE, order: uint(3)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: THREE, order: uint(3)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: THREE, order: uint(3)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: THREE, order: uint(3)}})
	//fours
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: FOUR, order: uint(4)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: FOUR, order: uint(4)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: FOUR, order: uint(4)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: FOUR, order: uint(4)}})
	//fives
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: FIVE, order: uint(5)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: FIVE, order: uint(5)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: FIVE, order: uint(5)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: FIVE, order: uint(5)}})
	//six
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: SIX, order: uint(6)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: SIX, order: uint(6)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: SIX, order: uint(6)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: SIX, order: uint(6)}})
	//seven
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: SEVEN, order: uint(7)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: SEVEN, order: uint(7)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: SEVEN, order: uint(7)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: SEVEN, order: uint(7)}})
	//eight
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: EIGHT, order: uint(8)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: EIGHT, order: uint(8)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: EIGHT, order: uint(8)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: EIGHT, order: uint(8)}})
	//nine
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: NINE, order: uint(9)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: NINE, order: uint(9)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: NINE, order: uint(9)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: NINE, order: uint(9)}})
	//ten
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: TEN, order: uint(10)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: TEN, order: uint(10)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: TEN, order: uint(10)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: TEN, order: uint(10)}})
	//jack
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: JACK, order: uint(11)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: JACK, order: uint(11)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: JACK, order: uint(11)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: JACK, order: uint(11)}})
	//queen
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: QUEEN, order: uint(12)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: QUEEN, order: uint(12)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: QUEEN, order: uint(12)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: QUEEN, order: uint(12)}})
	//king
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: KING, order: uint(13)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: KING, order: uint(13)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: KING, order: uint(13)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: KING, order: uint(13)}})
	//ace
	cards = append(cards, Card{color: CLUBS, rank: Rank{name: ACE, order: uint(14)}})
	cards = append(cards, Card{color: DIAMONDS, rank: Rank{name: ACE, order: uint(14)}})
	cards = append(cards, Card{color: HEARTS, rank: Rank{name: ACE, order: uint(14)}})
	cards = append(cards, Card{color: SPADES, rank: Rank{name: ACE, order: uint(14)}})
	//Let's shuffle cards to remove order
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return Desk{cards: cards}
}

func (desk *Desk) Next() (Card, error) {
	if len(desk.cards) == 0 {
		return Card{}, errors.New("no more cards")
	}
	lastCard := desk.cards[len(desk.cards)-1]
	desk.cards = desk.cards[:len(desk.cards)-1]
	return lastCard, nil
}
