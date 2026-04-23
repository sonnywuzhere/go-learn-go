package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("\nWelcome to the Euchre game: Terminal edition!")

	var player1Name string = "Bob"
	var player2Name string = "Alice"
	var player3Name string = "Drew"
	var player4Name string = "Joon"

	// comment out for easier testing
	fmt.Print("Player 1, enter your name: ")
	fmt.Scan(&player1Name)
	fmt.Print("Player 2, enter your name: ")
	fmt.Scan(&player2Name)
	fmt.Print("Player 3, enter your name: ")
	fmt.Scan(&player3Name)
	fmt.Print("Player 4, enter your name: ")
	fmt.Scan(&player4Name)
	fmt.Println()

	fmt.Println("Here are our players and teams: ")

	fmt.Println("Player 1 (Dealer)", player1Name)
	fmt.Println("Player 2", player2Name)
	fmt.Println("Player 3", player3Name)
	fmt.Println("Player 4", player4Name)

	fmt.Println("Team 0: ", player1Name, "and", player3Name)
	fmt.Println("Team 1: ", player2Name, "and", player4Name)

	var team1TotalPoints int
	var team2TotalPoints int 
	
	// a full game is to play until one team gets 10 points 
	for team1TotalPoints < 10 && team2TotalPoints < 10 {

		var deck = createDeck()
		var shuffledDeck = shuffleDeck(deck)

		// TODO: vary the dealer (player1) so it's not always the same
		var player1Hand, player2Hand, player3Hand, player4Hand, blind = dealCards(shuffledDeck)

		// name, team, bidder, tricksTaken, playerHand
		// player1 is the dealer
		var player1 = Player{player1Name, 0, false, 0, player1Hand}
		var player2 = Player{player2Name, 1, false, 0, player2Hand}
		var player3 = Player{player3Name, 0, false, 0, player3Hand}
		var player4 = Player{player4Name, 1, false, 0, player4Hand}

		var players = [4]Player{player1, player2, player3, player4}

		fmt.Println("\nThe start of a new round!")

		// the result of bidding is to select trump suit
		// the first card of the blind is "turned up" to decide on bidding
		var selectedTrump, playerWhoBid = bidding(player1, player2, player3, player4, blind[0])

		fmt.Println("\n", playerWhoBid.name, "bid!")

		// update player who bid data
		for j, player := range players {
			if player.name == playerWhoBid.name {
				players[j] = playerWhoBid
			}
		}

		// have player1 (dealer) replace card they picked up 
		fmt.Println("\n", player1.name, "what card would you like to replace for the", &blind[0], "?")
		fmt.Println("1, 2, 3...")
		fmt.Println(player1.playerHand)
		var replaceCard int
		fmt.Scan(&replaceCard)
		for i := 0; i < 4; i++ {
			if i == replaceCard {
				player1.playerHand[i - 1] = blind[0]
			}
		}
		fmt.Println("\n", player1.name, "your hand is now", player1.playerHand)

		// cards updated to reflect trump card's values
		players[0].playerHand, players[1].playerHand, players[2].playerHand, players[3].playerHand = updateCardTrumpValues(selectedTrump, player1Hand, player2Hand, player3Hand, player4Hand)

		// player to the left of the dealer leads first trick
		indexOfWhoLeads := 1
	
		// loop for one "round" (playing all cards) - 5
		for x := 0; x < 5; x++ {
			var cardsOnTable [4]Card
			var playersWhoPlayed [4]Player
			// loop for one trick (each player plays a card) - 4
			for i := 0; i < 4; i++ {
				
				fmt.Println("\n", players[indexOfWhoLeads].name, "'s turn")
				var playedCard Card
				playedCard, players[indexOfWhoLeads].playerHand = playCard(players[indexOfWhoLeads])
		
				cardsOnTable[i] = playedCard
				playersWhoPlayed[i] = players[indexOfWhoLeads]
		
				fmt.Println("\n", players[indexOfWhoLeads].name, "played a ", playedCard)
				fmt.Println("\n", players[indexOfWhoLeads].name, "hand", players[indexOfWhoLeads].playerHand)
		
				fmt.Println("\nCards on table: ")
				for z := range i + 1 {
					fmt.Printf("%s played by %s", cardsOnTable[z].String(), playersWhoPlayed[z].name)
					fmt.Println()
				}
				
				fmt.Println()
				fmt.Println(selectedTrump, "is trump")

				indexOfWhoLeads = (indexOfWhoLeads + 1) % 4
			}
			
			// determine who took the trick
			whoTookTrick := takeTrick(cardsOnTable, playersWhoPlayed, selectedTrump)
	
			for y, player := range players {
				if player.name == whoTookTrick.name {
					players[y] = whoTookTrick
					// the person who took the trick becomes the new leader 
					indexOfWhoLeads = y
				}
			}
	
		}

		// end of round (all cards have been played)
		t1, t2 := endOfRound(players)
		team1TotalPoints += t1 
		team2TotalPoints += t2

		fmt.Println("\nHere are the results of the end of the first round")
		fmt.Println("\nTeam 0:", team1TotalPoints)
		fmt.Println("Team1:", team2TotalPoints)
	}


	var whoWon = endOfGame(team1TotalPoints, team2TotalPoints)
	fmt.Println(whoWon)
	fmt.Println("Thanks for playing the Euchre terminal!")

}

// Card 
// rank, suit, value, isTrump (which changes value of cards for Jacks)
type Card struct {
	rank string 
	suit string 
	value int 
	isTrump bool
}

// constructor with default values 
func NewCard() *Card {
	return &Card{
		rank: "",
		suit: "",
		value: 0,
		isTrump: false,
	}
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

func (p *Player) String() string {
	return fmt.Sprintf("%s", p.name)
}


// playCard() 
// user enters the index of card they want to play
// card is removed from players hand 
// retuns card chosen
func playCard(player Player) (Card, [5]Card) {
	var cardPlayed Card
	var cardChosen int
	updatedHand := []Card{}
	
	fmt.Println(player.name, "What card would you like to play?")
	fmt.Println(player.playerHand)
	fmt.Println("1, 2, 3...")
	fmt.Scan(&cardChosen)
	
	cardPlayed = player.playerHand[cardChosen - 1]

	for i := 0; i < len(player.playerHand); i++ {
		if player.playerHand[i] != cardPlayed {
			updatedHand = append(updatedHand, player.playerHand[i])
		}
	}

	// add null card placeholders 
	cardPlaceholder := Card{"", "", 0, false}
	for i := len(updatedHand); i < 5; i++ {
		updatedHand = append(updatedHand, cardPlaceholder)
	}

	return cardPlayed, [5]Card(updatedHand)
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

	fmt.Println("\n", trump, "is trump")
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

func bidding(player1 Player, player2 Player, player3 Player, player4 Player, bidCard Card) (string, Player) {
	var players = [4]Player{player2, player3, player4, player1}
	// player to the left of the dealer goes first 
	// dealer bids last

	selectedTrump := ""
	bidOrPass := ""

	fmt.Println("\nThe top card is: ", &bidCard)
	
	// bidding continues until someone wins the bid
	// or until it goes around 2 times 
	// if it goes around 2 times, you "stick the dealer"

	// bid first on the top facing card 
	for j, player := range players {
		fmt.Println("\n", player.name, ": your turn to bid or pass")
		fmt.Println("Here are your cards", player.playerHand)
		fmt.Scan(&bidOrPass)

		if bidOrPass == "bid" {
			// this is for the scenario of the dealer's partner 
			// telling the dealer to pick it up
			if j == 1 {
				fmt.Println("Pick it up!")
			}
			player.bidder = true
			selectedTrump = bidCard.suit
			return selectedTrump, player
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
			return selectedTrump, player
		} else {
			fmt.Println(j, player.name, "your turn to bid or pass!")
			fmt.Scan(&bidOrPass)
	
			if bidOrPass == "bid" {
				fmt.Println("What do you bid?")
				fmt.Scan(&selectedTrump)
				player.bidder = true
				return selectedTrump, player
			} 
		} 
	} 

	// code should never be able to reach here, 
	// but was getting an error about needing a return 
	return selectedTrump, player1
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


// takeTrick() 
// determines who takes the trick 
// if trump was played, the highest trump card takes the trick 
// if no trump was played, the highest matching suit of the led card takes the trick
// adds to tricksTaken 
func takeTrick(cards [4]Card, players [4]Player, trumpSuit string) (Player) {

	var playerWhoTookTrick Player

	leadingCardSuit := cards[0].suit

	var filteredCards []Card
	var filteredPlayers []Player

	for i, card := range cards {
		if card.suit == leadingCardSuit || card.suit == trumpSuit {
			filteredCards = append(filteredCards, cards[i])
			filteredPlayers = append(filteredPlayers, players[i])
		}
	}

	maxValue, index := filteredCards[0].value, 0

	for i, card := range filteredCards {
		if card.value > maxValue {
			maxValue, index = filteredCards[i].value, i
		}
	}

	fmt.Println()
	fmt.Println(filteredPlayers[index].name, "took the trick!")
	playerWhoTookTrick = filteredPlayers[index]
	playerWhoTookTrick.tricksTaken++

	return playerWhoTookTrick
} 

// endOfRound() 
// once all cards have been played, see which team took the most tricks 
// if the team that bid took all the tricks, they get 2 points
// if the team that bid didn't take all the tricks, they only get 1 point ("held to one")
// otherwise, the team that didn't bid gets a point (the other team got "euched") 
func endOfRound(players [4]Player) (int, int) {
	var team1Points int 
	var team2Points int 
	var teamWhoBid int 
	var teamWhoWon int
	var teamWhoWonFullPoints bool

	var pointsForTeam1 int 
	var pointsForTeam2 int 

	for i, player := range players {
		if player.bidder {
			teamWhoBid = player.team % 2
		}
		if (i % 2) == 0 {
			team1Points += player.tricksTaken
		} else {
			team2Points += player.tricksTaken
		}
	}

	fmt.Println("team 0 got", team1Points, "points")
	fmt.Println("team 1 got", team2Points, "points")
	
	fmt.Println("Team", teamWhoBid, "bid")

	if team1Points > team2Points {
		teamWhoWon = 0
		if team1Points == 5 {
			teamWhoWonFullPoints = true 
		}
	} else if team2Points > team1Points {
		teamWhoWon = 1
		if team2Points == 5 {
			teamWhoWonFullPoints = true 
		}
	} 

	// see if the team who bid ended up winning
	if teamWhoBid == teamWhoWon {
		// check to see if they got all the points 
		if teamWhoWonFullPoints {
			fmt.Println("You got 2 points!")
			if teamWhoWon == 0 {
				pointsForTeam1 = 2
			} else {
				pointsForTeam2 = 2
			}
		} else {
			fmt.Println(teamWhoWon, "got held to 1 point!")
			fmt.Println("held 'em to 1!")
			if teamWhoWon == 0 {
				pointsForTeam1 = 1 
			} else {
				pointsForTeam2 = 1
			}
		}
	// the team who bid did not win 
	} else {
		fmt.Println(teamWhoBid, "got euched!")
		fmt.Println("The other team gets a point")
		// the team opposite to the team who bid gets points 
		if teamWhoBid == 0 {
			pointsForTeam2 = 1
		} else {
			pointsForTeam1 = 1
		}
	}

	return pointsForTeam1, pointsForTeam2
}

func endOfGame(team1Points int, team2Points int) (string) {
	var whoWon string 

	if team1Points > team2Points {
		whoWon = "Team 0 won!"
	} else {
		whoWon = "Team 1 won!"
	}
	return whoWon 
}

// Milestones: 
// Level 0: player functionality in terminal 
// Level 1: frontend built with Wails & Go code integrated (but still local functionality)
// Level 2: multi-player functionality (Scocket.io?) 
// Level 3: Deployed code 
