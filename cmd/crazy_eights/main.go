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

	var playedCard card.Card
	var err error
	for {
		// Checks to see if Player 1 has an eligible card in her hand
		// If she does, then this function does nothing
		// If she doesn't, then this function will force the player
		// to draw from the draw pile until they have a playable card
		// TODO: refactor this function to take out card drawing responsibility
		checkForPlayableCard(crazy, player1)

		// Present the user with the top of the pile and their hand
		topCard, _ := crazy.TopOfDiscardPile()
		fmt.Printf("\nTop of pile %v\n", topCard)
		showHand(crazy.Players[player1])

		for {
			// Player 1 chooses a card
			cardIndex, _ := promptUser(r)
			// Check if card played was valid
			playedCard, err = checkCard(crazy, r, player1, cardIndex)
			if err == nil {
				break
			}
			fmt.Println("Invalid card # - Please try again")
		}
		// Check if card played was an 8
		if playedCard.Rank() == card.Eight {
			fmt.Println("1: Spades 2: Hearts 3: Diamonds 4: Clubs")
			fmt.Print("Enter # of Suit you want to switch to: ")
			suit, err := getDesiredIndex(r)
			for {
				if err != nil || suit < 0 || suit > 4 {
					fmt.Println("Invalid choice - Please try again")
					fmt.Print("Enter # of Suit you want to switch to: ")
					suit, err = getDesiredIndex(r)
				} else {
					break
				}
			}
			crazy.HandleEight(suit)
		}

		// Check to see if Player 1's hand is empty
		if checkIfWinner(crazy.Players[player1]) {
			break
		}

		// CPU Turn
		cpuTurn(crazy, player2)
		// CPU plays the first valid card it finds in its hand
		for i := range crazy.Players[player2].Hand() {
			validPlay, _ := crazy.ValidPlay(player2, i)
			if validPlay {
				playedCard, err := crazy.PlayCard(player2, i)
				if err != nil {
					fmt.Fprintf(os.Stderr, "something went terribly wrong: %v\n", err)
					os.Exit(1)
				}
				fmt.Printf("CPU played %v\n", playedCard)
				// CPU switches to a random suit
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

// checkForPlayableCard makes sure a player has an eligible card to play
func checkForPlayableCard(crazy *game.CrazyEights, playerID int) {
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
func checkCard(crazy *game.CrazyEights, r *bufio.Reader, playerID, cardIndex int) (card.Card, error) {
	validPlay, err := crazy.ValidPlay(playerID, cardIndex)
	if err != nil {
		return card.Card{}, err
	}
	for validPlay == false {
		showHand(crazy.Players[playerID])
		topCard, _ := crazy.TopOfDiscardPile()
		fmt.Println("Top of pile", topCard)
		cardIndex, _ = promptUser(r)
		validPlay, err = crazy.ValidPlay(playerID, cardIndex)
		if err != nil {
			return card.Card{}, err
		}
	}
	playedCard, err := crazy.PlayCard(playerID, cardIndex)
	if err != nil {
		fmt.Fprintf(os.Stderr, "something went terribly wrong: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("You played %v\n", playedCard)
	return playedCard, nil

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
