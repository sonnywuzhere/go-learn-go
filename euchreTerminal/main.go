package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the Euchre game!")
	
	var deck = createDeck()
	var shuffledDeck = shuffleDeck(deck)

	// TODO: vary the dealer (player1) so it's not always the same
	var player1Hand, player2Hand, player3Hand, player4Hand, blind = dealCards(shuffledDeck)

	// name, team, bidder, tricksTaken, playerHand
	// player1 is the dealer
	var player1 = Player{"Sonny Joon", 1, false, 0, player1Hand}
	var player2 = Player{"Alice", 2, false, 0, player2Hand}
	var player3 = Player{"Nancy Drew", 1, false, 0, player3Hand}
	var player4 = Player{"Bob", 2, false, 0, player4Hand}

	var players = [4]Player{player1, player2, player3, player4}

	// the result of bidding is to select trump suit
	// the first card of the blind is "turned up" to decide on bidding
	var selectedTrump, playerWhoBid = bidding(player1, player2, player3, player4, blind[0])

	// update player who bid (bidder = true)
	for j, player := range players {
		if player.name == playerWhoBid {
			players[j].bidder = true
		}
	}

	// have player1 (dealer) replace card they picked up 
	fmt.Println(player1.name, "what card would you like to replace?")
	fmt.Println("0, 1, 2, 3, 4")
	fmt.Println(player1.playerHand)
	var replaceCard int
	fmt.Scan(&replaceCard)
	for i := 0; i < 4; i++ {
		if i == replaceCard {
			player1.playerHand[i] = blind[0]
		}
	}
	fmt.Println(player1.name, "your hand is now", player1.playerHand)

	// cards updated to reflect trump card's values
	player1.playerHand, player2.playerHand, player3.playerHand, player4.playerHand = updateCardTrumpValues(selectedTrump, player1Hand, player2Hand, player3Hand, player4Hand)

	fmt.Println(player1.playerHand, player2.playerHand, player3.playerHand, player4.playerHand)

	// 

}

// Card 
// rank, suit, value, isTrump (which changes value of cards for Jacks)
type Card struct {
	rank string 
	suit string 
	value int 
	isTrump bool
}

// "pretty print" Card info
func (c *Card) String() string {
	// https://www.geeksforgeeks.org/go-language/fmt-sprintf-function-in-golang-with-examples/
	// The Sprintf() function in Go is a way to format strings when printing 
	// You use %s for a string, %d for an integer, %f for a float 
	return fmt.Sprintf("%s of %s", c.rank, c.suit)
}

// Player 
// name, team, hasBid, tricksTaken, playerHand (map of Cards)
type Player struct {
	name string 
	team int 
	bidder bool 
	tricksTaken int
	playerHand [5]Card
}

// input is trump and all player hands 
// returns the updated player hands 
func updateCardTrumpValues(trump string, player1Hand [5]Card, player2Hand [5]Card, player3Hand [5]Card, player4Hand [5]Card) ([5]Card, [5]Card, [5]Card, [5]Card) {
	playerHands := [4][5]Card{player1Hand, player2Hand, player3Hand, player4Hand}

	complementarySuit := ""

	switch trump {
	case "hearts":
		complementarySuit = "diamonds"
	case "diamonds":
		complementarySuit = "hearts"
	case "spades":
		complementarySuit = "clubs"
	case "clubs":
		complementarySuit = "spades"
	}

	fmt.Println("trump is", trump)
	for i, hand := range playerHands {
		for j, card := range hand {
			if card.suit == trump {
				playerHands[i][j].isTrump = true 
				// Right bower
				if card.rank == "Jack" {
					playerHands[i][j].value += 11
				// all other trump values 
				} else {
					playerHands[i][j].value += 6
				}
			// Left bower 
			} else if card.suit == complementarySuit && card.rank == "Jack" {
				playerHands[i][j].isTrump = true 
				playerHands[i][j].value += 10
			}
		}
	}

	return playerHands[0], playerHands[1], playerHands[2], playerHands[3]
}

func bidding(player1 Player, player2 Player, player3 Player, player4 Player, bidCard Card) (string, string) {
	var players = [4]Player{player2, player3, player4, player1}
	// player to the left of the dealer goes first 
	// dealer bids last

	selectedTrump := ""
	bidOrPass := ""

	fmt.Println("The top card is: ", &bidCard)
	
	// bidding continues until someone wins the bid
	// or until it goes around 2 times 
	// if it goes around 2 times, you "stick the dealer"

	// bid first on the top facing card 
	for j, player := range players {
		fmt.Println(player.name, "your turn to bid or pass")
		fmt.Scan(&bidOrPass)

		if bidOrPass == "bid" {
			// this is for the scenario of the dealer's partner 
			// telling the dealer to pick it up
			if j == 2 {
				fmt.Println("Pick it up!")
			}
			selectedTrump = bidCard.suit
			return selectedTrump, player.name
		}
	}

	// now players bid on other suits
	for j, player := range players {
		// if no one else bids and it gets to the last player
		// then it's "stick the dealer"
		if j == 3 {
			fmt.Println("Stick the dealer!")
			fmt.Println(player.name, "what do you bid?")
			fmt.Scan(&selectedTrump)
			player.bidder = true
			return selectedTrump, player.name
		} else {
			fmt.Println(j, player.name, "your turn to bid or pass!")
			fmt.Scan(&bidOrPass)
	
			if bidOrPass == "bid" {
				fmt.Println("What do you bid?")
				fmt.Scan(&selectedTrump)
				player.bidder = true
				return selectedTrump, player.name
			} 
		} 
	} 

	// code should never be able to reach here, 
	// but was getting an error about needing a return 
	return selectedTrump, ""
}

// createDeck() 
// creates Card deck of cards 
func createDeck() ([24]Card) {
	// hearts, diamonds, clubs, spades 
	var suits [4]string 
	suits[0] = "hearts"
	suits[1] = "diamonds"
	suits[2] = "clubs"
	suits[3] = "spades"

	// 9, 10, J, Q, K, A
	var ranks [6]string 
	ranks[0] = "9"
	ranks[1] = "10"
	ranks[2] = "Jack"
	ranks[3] = "Queen"
	ranks[4] = "King"
	ranks[5] = "Ace"

	// 9, 10, 11, 12, 13, 14
	var values [6]int
	values[0] = 9
	values[1] = 10 
	values[2] = 11 
	values[3] = 12 
	values[4] = 13 
	values[5] = 14 

	var deck [24]Card
	
	var deckLength = 24
	var counter = 0
	for counter < deckLength {
		for i := 0; i < len(suits); i++ {
			for j := 0; j < len(ranks); j++ {
				deck[counter] = Card{ranks[j], suits[i], values[j], false}
				counter++
			}
		}
	}

	return deck
}

// shuffleDeck() 
// shuffles card deck 
func shuffleDeck(deck [24]Card) ([24]Card) {
	// https://www.tutorialspoint.com/article/golang-program-to-shuffle-the-elements-of-an-array
	var shuffledDeck [24]Card = deck
	for i := 0; i < len(shuffledDeck); i++ {
		j := rand.Intn(i + 1)
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	}
	return shuffledDeck
} 

func dealCards(shuffledDeck [24]Card) ([5]Card, [5]Card, [5]Card, [5]Card, [4]Card) {
	var player1Hand [5]Card
	var player2Hand [5]Card 
	var player3Hand [5]Card 
	var player4Hand [5]Card 
	var blind [4]Card 

	player1Hand[0] = shuffledDeck[0]
	player2Hand[0] = shuffledDeck[1]
	player3Hand[0] = shuffledDeck[2]
	player4Hand[0] = shuffledDeck[3]

	player1Hand[1] = shuffledDeck[4]
	player2Hand[1] = shuffledDeck[5]
	player3Hand[1] = shuffledDeck[6]
	player4Hand[1] = shuffledDeck[7]

	player1Hand[2] = shuffledDeck[8]
	player2Hand[2] = shuffledDeck[9]
	player3Hand[2] = shuffledDeck[10]
	player4Hand[2] = shuffledDeck[11]

	player1Hand[3] = shuffledDeck[12]
	player2Hand[3] = shuffledDeck[13]
	player3Hand[3] = shuffledDeck[14]
	player4Hand[3] = shuffledDeck[15]

	player1Hand[4] = shuffledDeck[16]
	player2Hand[4] = shuffledDeck[17]
	player3Hand[4] = shuffledDeck[18]
	player4Hand[4] = shuffledDeck[19]

	blind[0] = shuffledDeck[20]
	blind[1] = shuffledDeck[21]
	blind[2] = shuffledDeck[22]
	blind[3] = shuffledDeck[23]
	
	return player1Hand, player2Hand, player3Hand, player4Hand, blind
}


// playCard() 
// user enters the index of card they want to play
// card is removed from players hand 
// retuns card chosen


// takeTrick() 
// determines who takes the trick 
// if trump was played, the highest trump card takes the trick 
// if no trump was played, the highest matching suit of the led card takes the trick
// adds to tricksTaken 

// endOfRound() 
// once all cards have been played, see which team took the most tricks 
// if the team that bid took all the tricks, they get 2 points
// if the team that bid took more than the other team, they get 1 points ("held to one")
// otherwise, the team that didn't bid gets a point (the other team got "euched") 


// keep track of whose turn it is 
// keep track of who leads (player 1 starts)
// loop through 5 times (to play all cards in hand)
// create "table" array
// leader plays card - card added to "table"
// All other players play a card 
// takeTrick() 
// update who leads (whichever player took the trick)


// Milestones: 
// Level 0: player functionality in terminal 
// Level 1: frontend built with Wails & Go code integrated (but still local functionality)
// Level 2: multi-player functionality (Scocket.io?) 
// Level 3: Deployed code 
