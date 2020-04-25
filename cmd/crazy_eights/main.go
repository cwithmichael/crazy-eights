package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/cwithmichael/crazy_eights/internal/card"
	"github.com/cwithmichael/crazy_eights/pkg/game"
	"github.com/cwithmichael/crazy_eights/pkg/player"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Crazy 8s")
	numberOfPlayers, cardsPerPlayer := 2, 7
	crazy := game.NewGame(numberOfPlayers)
	crazy.DealCards(cardsPerPlayer)
	player1, player2 := 0, 1

	for {
		// Player 1's turn
		realPlayerTurn(crazy, player1)
		// Player 1 chooses a card
		cardIndex, _ := promptUser(r)
		// Check if card played was valid
		playedCard := checkCard(crazy, r, player1, cardIndex)
		// Check if card played was an 8
		if playedCard.Rank() == card.Eight {
			fmt.Println("1: Spades 2: Hearts 3: Diamonds 4: Clubs")
			fmt.Print("Enter # of Suit you want to play: ")
			suit, err := getDesiredIndex(r)
			for err != nil {
				fmt.Println("Invalid choice")
				suit, err = getDesiredIndex(r)
			}
			crazy.HandleEight(suit)
		}

		// Check to see if Player 1's hand is empty
		if checkIfWinner(crazy.Players[player1]) {
			break
		}

		// CPU Turn
		cpuTurn(crazy, player2)
		// CPU plays the first valid card
		for i := 0; i < len(crazy.Players[player2].Hand()); i++ {
			if crazy.ValidPlay(player2, i) {
				playedCard, err := crazy.PlayCard(player2, i)
				if err != nil {
					fmt.Fprintf(os.Stderr, "something went terribly wrong: %v\n", err)
					os.Exit(1)
				}
				fmt.Printf("CPU played %v\n", playedCard)
				if playedCard.Rank() == card.Eight {
					newSuit := rand.Intn(4)
					crazy.HandleEight(newSuit)
					fmt.Printf("CPU switched the suit to %v\n", card.Suit(newSuit))
				}
				break
			}
		}
		// Check to see if the cpu's hand is empty
		if checkIfWinner(crazy.Players[player2]) {
			break
		}

	}
}

// realPlayerTurn makes sure a player has an eligible card to play
func realPlayerTurn(crazy *game.CrazyEights, playerID int) {
	eligibleTurn := crazy.EligibleTurn(playerID)
	for eligibleTurn == false {
		fmt.Println("Drawing a card because you don't have any playable cards")
		_, err := crazy.DrawCard(playerID)
		//TODO: Handle draw deck being empty
		if err != nil {
			fmt.Println("The draw deck is exhausted")
			os.Exit(0)
		}
		eligibleTurn = crazy.EligibleTurn(playerID)
	}
	topCard, _ := crazy.TopOfDiscardPile()
	fmt.Printf("\nTop of pile %v\n", topCard)
	showHand(crazy.Players[playerID])
}

func cpuTurn(crazy *game.CrazyEights, playerID int) {
	eligibleTurn := crazy.EligibleTurn(playerID)
	for eligibleTurn == false {
		fmt.Println("CPU is drawing a card")
		_, err := crazy.DrawCard(playerID)
		//TODO: Handle draw deck being empty
		if err != nil {
			fmt.Println("The draw deck is exhausted")
			os.Exit(0)
		}
		eligibleTurn = crazy.EligibleTurn(playerID)
	}
}

// checkCard checks to see if the card played is a valid play
func checkCard(crazy *game.CrazyEights, r *bufio.Reader, playerID, cardIndex int) card.Card {
	validPlay := crazy.ValidPlay(playerID, cardIndex)
	for validPlay == false {
		showHand(crazy.Players[playerID])
		topCard, _ := crazy.TopOfDiscardPile()
		fmt.Println("Top of pile", topCard)
		cardIndex, _ = promptUser(r)
		validPlay = crazy.ValidPlay(playerID, cardIndex)
	}
	playedCard, err := crazy.PlayCard(playerID, cardIndex)
	if err != nil {
		fmt.Fprintf(os.Stderr, "something went terribly wrong: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("You played %v\n", playedCard)
	return playedCard

}

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
	fmt.Printf("\nEnter # of card you want to play: ")
	cardIndex, err := getDesiredIndex(r)
	return cardIndex, err
}

func getDesiredIndex(r *bufio.Reader) (int, error) {
	input, _, err := r.ReadLine()
	if err != nil {
		return -1, err
	}
	index, err := strconv.Atoi(string(input))
	if err != nil {
		return -1, err
	}
	return index - 1, nil
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
