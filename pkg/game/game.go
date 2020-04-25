package game

import (
	"errors"
	"github.com/cwithmichael/crazy_eights/internal/card"
	"github.com/cwithmichael/crazy_eights/internal/deck"
	"github.com/cwithmichael/crazy_eights/pkg/player"
)

// CrazyEights represents a game of Crazy 8s
type CrazyEights struct {
	Players     []*player.Player
	DrawPile    []card.Card
	DiscardPile []card.Card
}

// NewGame creates a new game of Crazy 8s
// numberOfPlayers is the deisred number of Players for the game
func NewGame(numberOfPlayers int) *CrazyEights {
	standardDeck := deck.NewDeck()
	shuffledCards := standardDeck.ShuffleCards()
	players := make([]*player.Player, numberOfPlayers)
	for i := 0; i < numberOfPlayers; i++ {
		players[i] = player.NewPlayer(i)
	}

	return &CrazyEights{Players: players, DrawPile: shuffledCards, DiscardPile: make([]card.Card, 0)}
}

// DealCards deals the Cards to the Players in the Game
// cardDist is the number of Cards to distribute to each Player
func (c8 *CrazyEights) DealCards(cardDist int) {
	for i := 0; i < len(c8.Players); i++ {
		c8.Players[i].AddToHand(c8.DrawPile[cardDist*i : cardDist*(i+1)]...)
	}
	idx := cardDist * len(c8.Players)
	if idx+1 < len(c8.DrawPile) {
		c8.DrawPile = append(c8.DrawPile[:idx], c8.DrawPile[idx+1:]...)
	}
	c8.addToDiscardPile(c8.DrawPile[len(c8.DrawPile)-1])
	c8.DrawPile = c8.DrawPile[:len(c8.DrawPile)-1]
}

// DrawCard draws a Card from the top of the DrawPile
// playerID is the ID of the Player that is drawing the Card
func (c8 *CrazyEights) DrawCard(playerID int) (card.Card, error) {
	topOfPile, err := c8.TopOfDrawPile()
	if err != nil {
		return card.Card{}, err
	}
	c8.Players[playerID].AddToHand(topOfPile)
	c8.DrawPile = c8.DrawPile[:len(c8.DrawPile)-1]
	return topOfPile, nil
}

// PlayCard removes a Card from a Player's hand
// playerID is the id of the Player that is playing the Card
// cardIndex is the index of the Card
func (c8 *CrazyEights) PlayCard(playerID, cardIndex int) (card.Card, error) {
	hand := c8.Players[playerID].Hand()
	if cardIndex >= len(hand) {
		return card.Card{}, errors.New("Invalid card index")
	}
	discardedCard := hand[cardIndex]
	err := c8.Players[playerID].DiscardFromHand(cardIndex)
	if err != nil {
		return card.Card{}, err
	}
	c8.addToDiscardPile(discardedCard)
	return discardedCard, nil
}

func (c8 *CrazyEights) addToDiscardPile(discardedCard card.Card) {
	c8.DiscardPile = append(c8.DiscardPile, discardedCard)
}

// HandleEight changes the suit of the 8 to match the player's
// preferred choice of suit
func (c8 *CrazyEights) HandleEight(desiredSuit int) error {
	if desiredSuit < 0 || desiredSuit > 4 {
		return errors.New("Unknown Suit")
	}
	if len(c8.DiscardPile) > 0 {
		c8.DiscardPile = c8.DiscardPile[:len(c8.DiscardPile)-1]
	}
	c8.DiscardPile = append(c8.DiscardPile, card.NewCard(card.Suit(desiredSuit), card.Eight))
	return nil
}

// TopOfDiscardPile shows the top card of the DiscardPile
func (c8 *CrazyEights) TopOfDiscardPile() (card.Card, error) {
	if len(c8.DiscardPile) > 0 {
		return c8.DiscardPile[len(c8.DiscardPile)-1], nil
	}
	return card.Card{}, errors.New("Empty card list")
}

// TopOfDrawPile shows the top card of the DrawPile
func (c8 *CrazyEights) TopOfDrawPile() (card.Card, error) {
	if len(c8.DrawPile) > 0 {
		return c8.DrawPile[len(c8.DrawPile)-1], nil
	}
	return card.Card{}, errors.New("Empty card list")
}

// EligibleTurn checks to see if a Player can play
// with any of the Cards in their hand
// playerID is the ID of the Player
func (c8 *CrazyEights) EligibleTurn(playerID int) bool {
	topCard, _ := c8.TopOfDiscardPile()
	for _, v := range c8.Players[playerID].Hand() {
		if v.Rank() == card.Eight ||
			v.Suit() == topCard.Suit() ||
			v.Rank() == topCard.Rank() {
			return true
		}
	}
	return false
}

// ValidPlay checks to see if the Card played is a valid option
// playerID is the ID of the Player playing the Card
// cardIndex is the index of the played Card from the Player's hand
func (c8 *CrazyEights) ValidPlay(playerID int, cardIndex int) bool {
	topCard, _ := c8.TopOfDiscardPile()
	return c8.Players[playerID].Hand()[cardIndex].Rank() == card.Eight ||
		c8.Players[playerID].Hand()[cardIndex].Rank() == topCard.Rank() ||
		c8.Players[playerID].Hand()[cardIndex].Suit() == topCard.Suit()
}
