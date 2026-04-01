package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the Euchre game!")
	
	var deck = createDeck()
	var shuffledDeck = shuffleDeck(deck)

	
	var player1Hand, player2Hand, player3Hand, player4Hand, blind = dealCards(shuffledDeck)

	var player1 = Player{"Sonny Joon", 1, false, 0, player1Hand}
	var player2 = Player{"Alice", 2, false, 0, player2Hand}
	var player3 = Player{"Nancy Drew", 1, false, 0, player3Hand}
	var player4 = Player{"Bob", 2, false, 0, player4Hand}

	// the result of bidding is to select trump suit
	var selectedTrump = bidding(player1, player2, player3, player4)

	fmt.Println(selectedTrump)

	fmt.Println("The blind: ", blind)

}

// Card 
// rank, suit, value 
// isTrump (which changes value of cards for Jacks)
type Card struct {
	rank string 
	suit string 
	value int 
	isTrump bool
}

// "pretty print" Card info


// Player 
// name, team, hasBid, tricksTaken 
// playerHand --> map {rankSuit: value}? struct?
type Player struct {
	name string 
	team int 
	hasBid bool 
	tricksTaken int
	playerHand [5]Card
}

func bidding(player1 Player, player2 Player, player3 Player, player4 Player) (string) {
	selectedTrump := ""
	
	var suitsRanking = map[string]int{"hearts": 4, "diamonds": 3, "clubs": 2, "spades": 1}
	
	var currentBid int = 0
	// bidding continues until someone wins the bid
	// or until it goes around 2 times 

	var players [4]Player 
	players[0] = player1
	players[1] = player2
	players[2] = player3
	players[3] = player4

	for i := 0; i < 2; i++ {
		for i, player := range players {
			fmt.Println(i, player.name, "your turn to bid or pass!")
			
			fmt.Scan(&currentBid)

			if currentBid 
		}
	}

	fmt.Println(currentBid)

	return selectedTrump
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
	ranks[2] = "J"
	ranks[3] = "Q"
	ranks[4] = "K"
	ranks[5] = "A"

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

// "pretty print" Player info


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

// Game Logic 
// createDeck()
// shuffleDeck() 
// deal 5 cards to each player, with 4 cards left in the blind 


// keep track of whose turn it is 
// keep track of who leads (player 1 starts)
// loop through 5 times (to play all cards in hand)
// create "table" array
// leader plays card - card added to "table"
// All other players play a card 
// takeTrick() 
// update who leads (whichever player took the trick)

// endOfRound() 


// Milestones: 
// Level 0: player functionality in terminal 
// Level 1: frontend built with Wails & Go code integrated (but still local functionality)
// Level 2: multi-player functionality (Scocket.io?) 
// Level 3: Deployed code 
