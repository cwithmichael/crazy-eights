package player

import (
	"errors"
	"fmt"

	"github.com/cwithmichael/crazy_eights/internal/card"
)

// Player represents a player in a Crazy 8s game
type Player struct {
	id   int
	hand []card.Card
}

func (player *Player) String() string {
	return fmt.Sprintf("Player %d | Hand: %v", player.id, player.hand)
}

// NewPlayer creates a new Player instance with a given ID
func NewPlayer(id int) *Player {
	return &Player{id: id, hand: make([]card.Card, 0)}
}

// Hand returns the Player's current hand
func (player *Player) Hand() []card.Card {
	return player.hand
}

// ID returns the Player's ID
func (player *Player) ID() int {
	return player.id
}

// AddToHand adds a Card (or Cards) to a Player's hand
func (player *Player) AddToHand(cards ...card.Card) {
	player.hand = append(player.hand, cards...)
}

// DiscardFromHand removes a Card specified by cardIndex
// from the Player's hand
func (player *Player) DiscardFromHand(cardIndex int) error {
	if len(player.hand) < 1 {
		return errors.New("Can't discard from empty hand")
	}
	player.hand = append(player.hand[:cardIndex], player.hand[cardIndex+1:]...)
	return nil
}
