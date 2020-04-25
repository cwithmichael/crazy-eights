package game

import (
	"reflect"
	"testing"

	"github.com/cwithmichael/crazy_eights/internal/card"
	"github.com/cwithmichael/crazy_eights/pkg/player"
)

func TestCrazyEights_DealCards(t *testing.T) {
	testPlayers := []*player.Player{player.NewPlayer(0), player.NewPlayer(1)}
	testDrawPile := []card.Card{
		card.NewCard(card.Clubs, card.Five),
		card.NewCard(card.Diamonds, card.Four),
		card.NewCard(card.Hearts, card.Jack),
		card.NewCard(card.Spades, card.Nine),
	}
	testDiscardPile := make([]card.Card, 0)
	type fields struct {
		Players     []*player.Player
		DrawPile    []card.Card
		DiscardPile []card.Card
	}
	type args struct {
		cardDist int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"DealCards",
			fields{Players: testPlayers, DrawPile: testDrawPile, DiscardPile: testDiscardPile},
			args{cardDist: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c8 := &CrazyEights{
				Players:     tt.fields.Players,
				DrawPile:    tt.fields.DrawPile,
				DiscardPile: tt.fields.DiscardPile,
			}
			c8.DealCards(tt.args.cardDist)
		})
	}
}

func TestCrazyEights_DrawCard(t *testing.T) {
	type fields struct {
		Players     []*player.Player
		DrawPile    []card.Card
		DiscardPile []card.Card
	}
	type args struct {
		playerID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    card.Card
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c8 := &CrazyEights{
				Players:     tt.fields.Players,
				DrawPile:    tt.fields.DrawPile,
				DiscardPile: tt.fields.DiscardPile,
			}
			got, err := c8.DrawCard(tt.args.playerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrazyEights.DrawCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CrazyEights.DrawCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrazyEights_EligibleTurn(t *testing.T) {
	type fields struct {
		Players     []*player.Player
		DrawPile    []card.Card
		DiscardPile []card.Card
	}
	type args struct {
		playerID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c8 := &CrazyEights{
				Players:     tt.fields.Players,
				DrawPile:    tt.fields.DrawPile,
				DiscardPile: tt.fields.DiscardPile,
			}
			if got := c8.EligibleTurn(tt.args.playerID); got != tt.want {
				t.Errorf("CrazyEights.EligibleTurn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrazyEights_PlayCard(t *testing.T) {
	type fields struct {
		Players     []*player.Player
		DrawPile    []card.Card
		DiscardPile []card.Card
	}
	type args struct {
		playerID  int
		cardIndex int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    card.Card
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c8 := &CrazyEights{
				Players:     tt.fields.Players,
				DrawPile:    tt.fields.DrawPile,
				DiscardPile: tt.fields.DiscardPile,
			}
			got, err := c8.PlayCard(tt.args.playerID, tt.args.cardIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrazyEights.PlayCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CrazyEights.PlayCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrazyEights_ValidPlay(t *testing.T) {
	type fields struct {
		Players     []*player.Player
		DrawPile    []card.Card
		DiscardPile []card.Card
	}
	type args struct {
		playerID  int
		cardIndex int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c8 := &CrazyEights{
				Players:     tt.fields.Players,
				DrawPile:    tt.fields.DrawPile,
				DiscardPile: tt.fields.DiscardPile,
			}
			got, err := c8.ValidPlay(tt.args.playerID, tt.args.cardIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrazyEights.ValidPlay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CrazyEights.ValidPlay() = %v, want %v", got, tt.want)
			}
		})
	}
}
