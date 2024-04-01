package ws

import (
	"fmt"
	"log"
	"slices"
)

func gameState(msgReq *MessageReq, hub *Hub, game *Game, p *Player) *Message {
	players := game.Players
	room := hub.Rooms[p.RoomID]
	log.Printf("PLAYERS: %v", players)
	var remark string
	var state int
	switch msgReq.Action {

	case READY:
		p.Ready = true
		remark = fmt.Sprintf("Player %v turns ready!", p.ID)
		if checkReady(players) {
			remark = "Start game!"
			state = CHOOSE_CARD
			game.State = CHOOSE_CARD
			room.State = PLAY
			for _, p := range players {
				p.Ready = false
			}
			initGame(hub.Games[p.RoomID])
		}

	case UNREADY:
		remark = fmt.Sprintf("Player %v turns NOT ready!", p.ID)
		p.Ready = false

	case PLAY:
		// play cards
		sel := msgReq.Card
		remark = fmt.Sprintf("%v played card.", p.Name)
		game.Select[p.ID] = sel
		if len(game.Select) == len(players) {
			for ID, sel := range game.Select {
				game.Played[sel] = ID
			}
			// Remove played card from players' hands
			removePlayed(game.Played, players, game)
			sortedCards := getSortedCards(game.Played)
			if isSmallest(sortedCards[0], game) {
				game.State = CHOOSE_ROW
				state = CHOOSE_ROW
				remark = game.Played[sortedCards[0]]
			} else {
				game.State = CHOOSE_CARD
				state = CHOOSE_CARD
				// playCards
				remark = "Played cards!"
				playCards(game)
				if len(p.Hand) == 0 {
					// end game
					state = endGame(game)
				}
			}

		}
		// room.Players = players

	case UNPLAY:
		delete(game.Played, msgReq.Card)
		remark = fmt.Sprintf("Player %v recalled card!", p.ID)
		// unplay cards

	case ROW:
		sel := msgReq.Row
		card := msgReq.Card
		remark = fmt.Sprintf("%v selected row: %v!", p.ID, sel)
		// EAT POINTS
		row := game.Deck[sel]
		p.Score += getScore(row)
		game.Scores[p.ID] = p.Score
		game.Deck[sel] = []int{card}
		delete(game.Played, card)
		game.State = CALCULATING
		playCards(game)
		if len(p.Hand) == 0 {
			// end game
			endGame(game)
		}
		state = CHOOSE_CARD
		// updatePlayers(players)

	case LOBBY:
		state = LOBBY
		room.State = LOBBY
		remark = "Back to lobby!"

	}

	msg := &Message{
		// ID:     p.ID,
		GameID: p.RoomID,
		State:  state,
		Remark: remark,
	}
	return msg

}

func initGame(game *Game) {
	fullDeck := getFullDeck()
	// Populate hands
	// Populate deck
	// Wait for ready
	handLimit := 4
	players := game.Players
	if len(players) == 10 {
		handLimit = 10
	}
	start := 0
	for _, player := range players {
		dealtHand := fullDeck[start : start+handLimit]
		slices.Sort(dealtHand)
		player.Hand = dealtHand
		game.Hands[player.ID] = dealtHand
		start += handLimit
	}
	var deck [][]int
	for i := start; i < start+4; i++ {
		row := []int{fullDeck[i]}
		deck = append(deck, row)
	}
	game.Deck = deck
}

func playCards(game *Game) {
	//Check for nearest
	players := game.Players
	playedCard := game.Played
	sortedCards := getSortedCards(playedCard)
	deck := game.Deck
	// Calculation
	for _, card := range sortedCards {
		nearestPos := getNearest(card, deck)
		if len(deck[nearestPos]) == 5 {
			// BUSTED
			players[playedCard[card]].Score += getScore(deck[nearestPos])
			deck[nearestPos] = []int{}
		}
		deck[nearestPos] = append(deck[nearestPos], card)
	}
	game.Played = make(map[int]string)
	game.Select = make(map[string]int)

	game.State = CHOOSE_CARD

}

func endGame(game *Game) int {
	// Check if any player crosse 66 points
	// -> If no, init new round
	// -> If yes, go back to lobby

	var exceed bool
	for _, p := range game.Players {
		if p.Score >= 66 {
			exceed = true
			break
		}
	}

	if exceed {
		// WE GOT A WINNER!
		game.State = INIT
		resetScore(game.Players)
		return GAME_END
		// go back to lobby
	} else {
		game.State = ROUND_END
		// start new round
		return ROUND_END
	}

}
