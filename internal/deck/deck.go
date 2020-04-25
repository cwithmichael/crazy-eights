package deck

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cwithmichael/crazy_eights/internal/card"
)

// StandardDeck represents a standard deck of Cards
type StandardDeck struct {
	cards []card.Card
}

// NewDeck returns a new StandardDeck with the cards
// field initialized with 52 cards
func NewDeck() StandardDeck {
	cards := make([]card.Card, 52)
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			cards[(i*13)+j] = card.NewCard(card.Suit(i), card.Rank(j))
		}
	}
	return StandardDeck{cards: cards}
}

func (deck StandardDeck) String() string {
	return fmt.Sprintf("Cards: %v", deck.cards)
}

// ShuffleCards shuffles the StandardDeck and returns a
// new slice of the shuffled cards (Does not modify the original deck)
func (deck StandardDeck) ShuffleCards() []card.Card {
	rand.Seed(time.Now().UTC().UnixNano())
	shuffledCards := make([]card.Card, len(deck.cards))
	copy(shuffledCards, deck.cards)
	rand.Shuffle(len(deck.cards), func(i, j int) {
		shuffledCards[i], shuffledCards[j] = shuffledCards[j], shuffledCards[i]
	})
	return shuffledCards
}
