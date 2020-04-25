package player

import (
	"testing"

	"github.com/cwithmichael/crazy_eights/internal/card"
)

func TestPlayer_AddToHand(t *testing.T) {
	testCards := []card.Card{card.NewCard(card.Hearts, card.Five), card.NewCard(card.Spades, card.Ten)}
	type fields struct {
		id   int
		hand []card.Card
	}
	type args struct {
		cards []card.Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"AddMultipleCardsToHand", fields{id: 0, hand: make([]card.Card, 0)}, args{cards: testCards}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := &Player{
				id:   tt.fields.id,
				hand: tt.fields.hand,
			}
			player.AddToHand(tt.args.cards...)
			if len(player.hand) != len(tt.args.cards) {
				t.Errorf("CrazyEights.AddToHand() got = %d, want %d", len(player.hand), len(tt.args.cards))
			}
		})
	}
}

func TestPlayer_DiscardFromHand(t *testing.T) {
	testHand := []card.Card{card.NewCard(card.Clubs, card.Four), card.NewCard(card.Diamonds, card.Ace)}
	type fields struct {
		id   int
		hand []card.Card
	}
	type args struct {
		cardIndex int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"DiscardFromEmptyHand", fields{id: 0, hand: make([]card.Card, 0)}, args{cardIndex: 1}, true},
		{"DiscardFromHand", fields{id: 0, hand: testHand}, args{cardIndex: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := &Player{
				id:   tt.fields.id,
				hand: tt.fields.hand,
			}
			if err := player.DiscardFromHand(tt.args.cardIndex); (err != nil) != tt.wantErr {
				t.Errorf("Player.DiscardFromHand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
