package card

import "fmt"

// Suit represents the suit of a standard card
type Suit int

// Rank represents the rank of a standard card
type Rank int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Card is a representation of a standard card
type Card struct {
	rank Rank
	suit Suit
}

// NewCard creates a new Card with a given Suit and Rank
func NewCard(suit Suit, rank Rank) Card {
	return Card{suit: suit, rank: rank}
}

func (card Card) String() string {
	return fmt.Sprintf("[%v of %v]", card.rank, card.suit)
}

// Rank returns the Rank of a card
func (card Card) Rank() Rank {
	return card.rank
}

// Suit returns the Suit of a card
func (card Card) Suit() Suit {
	return card.suit
}

func (suit Suit) String() string {
	suits := [...]string{
		"Spades",
		"Hearts",
		"Diamonds",
		"Clubs",
	}

	if suit < 0 || suit > 4 {
		return "Unknown"
	}

	return suits[suit]
}

func (rank Rank) String() string {
	ranks := [...]string{
		"Ace",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"J",
		"Q",
		"K",
	}

	if rank < 0 || rank > 13 {
		return "Unknown"
	}

	return ranks[rank]
}
