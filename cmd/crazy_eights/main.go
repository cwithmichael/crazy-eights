package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/cwithmichael/crazy_eights/pkg/game"
	"github.com/cwithmichael/crazy_eights/pkg/player"
)

func showHand(p *player.Player) {
	fmt.Printf("\nCurrent hand: \n")
	for k, v := range p.Hand() {
		if k >= 3 && k%3 == 0 {
			fmt.Println()
		}
		fmt.Printf("(%d %v) ", k+1, v)
	}
	fmt.Println()
}

func promptUser(r *bufio.Reader) (int, error) {
	fmt.Print("Enter # of card you want to play: ")
	cardIndex, err := getDesiredCardIndex(r)
	return cardIndex, err
}

func getDesiredCardIndex(r *bufio.Reader) (int, error) {
	input, _, err := r.ReadLine()
	cardIndex, err := strconv.Atoi(string(input))
	if err != nil {
		return -1, err
	}
	return cardIndex - 1, nil
}

func checkIfWinner(p *player.Player) bool {
	if len(p.Hand()) == 0 {
		switch p.ID() {
		case 0:
			fmt.Println("You won!!!!!")
		case 1:
			fmt.Println("The CPU doesn't have any cards left.")
			fmt.Println("You Lost :(")

		}
		return true
	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Crazy 8s")
	crazy := game.NewGame(2)
	crazy.DealCards(7)
	player1, player2 := 0, 1

	for {
		// Player 1 Turn
		eligibleTurn := crazy.EligibleTurn(player1)
		for eligibleTurn == false {
			fmt.Println("Drawing a card because you don't have any playable cards")
			_, err := crazy.DrawCard(player1)
			//TODO: Handle draw deck being empty
			if err != nil {
				fmt.Println("The draw deck is exhuasted")
				os.Exit(0)
			}
			eligibleTurn = crazy.EligibleTurn(player1)
		}
		topCard, _ := crazy.TopOfDiscardPile()
		fmt.Printf("\nTop of pile %v\n", topCard)
		showHand(crazy.Players[player1])
		cardIndex, _ := promptUser(r)

		// Check if card played was valid
		validPlay := crazy.ValidPlay(player1, cardIndex)
		for validPlay == false {
			showHand(crazy.Players[player1])
			topCard, _ = crazy.TopOfDiscardPile()
			fmt.Println("Top of pile", topCard)
			cardIndex, _ = promptUser(r)
			validPlay = crazy.ValidPlay(player1, cardIndex)
		}
		playedCard, err := crazy.PlayCard(player1, cardIndex)
		if err != nil {
			fmt.Errorf("something went terribly wrong: %v", err)
		}
		fmt.Printf("You played %v\n", playedCard)
		if checkIfWinner(crazy.Players[player1]) {
			break
		}

		// CPU Turn
		eligibleTurn = crazy.EligibleTurn(player2)
		for eligibleTurn == false {
			fmt.Println("CPU is drawing a card")
			_, err := crazy.DrawCard(player2)
			//TODO: Handle draw deck being empty
			if err != nil {
				fmt.Println("The draw deck is exhuasted")
				os.Exit(0)
			}
			eligibleTurn = crazy.EligibleTurn(player2)
		}
		// CPU plays the first valid card
		for i := 0; i < len(crazy.Players[player2].Hand()); i++ {
			if crazy.ValidPlay(player2, i) {
				playedCard, err := crazy.PlayCard(player2, i)
				if err != nil {
					fmt.Errorf("something went terribly wrong: %v", err)
				}
				fmt.Printf("CPU played %v\n", playedCard)
				break
			}
		}
		if checkIfWinner(crazy.Players[player2]) {
			break
		}

	}
}
