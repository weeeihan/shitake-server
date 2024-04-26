package ws

import (
	"fmt"
	"log"
	"slices"
	"time"
)

func (room *Room) gameState(msgReq *MessageReq, p *Player, hub *Hub) {

	var msg *Message
	switch msgReq.Action {

	case LEAVE:
		// player leaves room
		msg = p.leaveRoom(room)

	case READY:
		// Player ready
		msg = p.ready(room)

	case START:
		// Player presses start
		msg = startGame(room)

	case UNREADY:
		// Player gets unready
		msg = p.unready()

	case PLAY:
		// play cards
		msg = room.play(msgReq, p, hub)
		// room.Players = players

	case PROCESS:
		// reset timer
		msg = room.processCards(p, hub)

	case ROW:
		msg = room.rows(msgReq.Row, msgReq.Card, p)
		hub.Broadcast <- createMsg(room.ID, CHOOSE_CARD, deckTostring(room.Deck))

	// case ROW:
	// 	sel := msgReq.Row
	// 	card := msgReq.Card
	// 	remark = fmt.Sprintf("%v selected row: %v!", p.ID, sel)
	// 	// EAT POINTS
	// 	row := game.Deck[sel]
	// 	p.Score += getScore(row)
	// 	game.Scores[p.ID] = p.Score
	// 	game.Deck[sel] = []int{card}
	// 	delete(game.Played, card)
	// 	game.State = CALCULATING
	// 	playCards(game)
	// 	if len(p.Hand) == 0 {
	// 		// end game
	// 		state = endGame(game)
	// 	} else {
	// 		state = CHOOSE_CARD
	// 	}
	// 	// updatePlayers(players)

	// case LOBBY:
	// 	state = LOBBY
	// 	room.State = LOBBY
	// 	resetScore(game.Players)
	// 	remark = "Back to lobby!"

	case COUNT:
		// if room.Ticker == nil {
		// 	room.Ticker = time.NewTicker(1 * time.Second)

		// } else {
		// 	room.Ticker.Stop()
		// 	room.Ticker = time.NewTicker(1 * time.Second)
		// 	go func(i int) {
		// 		for ; true; <-room.Ticker.C {
		// 			if i == 0 {
		// 				gameState(&MessageReq{Action: PROCESS}, hub)
		// 				return
		// 			}
		// 			log.Printf("Count: %v", i)
		// 			i--
		// 		}
		// 	}(3)
		// }

	case STOPCOUNT:
		// room.Ticker.Stop()

	case CHAT:
		msg = createMsg(room.ID, CHAT, "blabla")

	case PING:
		msg = createMsg(room.ID, PING, "Pong!")
		// testConn(room)

	}
	// showhands(room)

	hub.Broadcast <- msg

}

// PRE-GAME HANDLERS

func (player *Player) leaveRoom(room *Room) *Message {
	close(player.Message)
	delete(room.Players, player.ID)
	player.Conn.Close()
	return createMsg(room.ID, PLAYER_LEFT, fmt.Sprintf("Player %v left the room!", player.Name))
}

func (player *Player) ready(room *Room) *Message {
	player.Ready = true

	if checkReady(room.Players) {
		if room.State == INIT {

			return createMsg(room.ID, ALREADY, "All ready to start!")
			// Call init game
		} else {
			// Play next round
		}
	}

	return createMsg(room.ID, READY, fmt.Sprintf("Player %v is ready!", player.Name))
}

func (player *Player) unready() *Message {
	player.Ready = false
	return &Message{
		RoomID: player.RoomID,
		State:  UNREADY,
		Remark: fmt.Sprintf("Player %v turns NOT ready!", player.Name),
	}
}

func startGame(room *Room) *Message {
	if checkReady(room.Players) {
		for _, p := range room.Players {
			p.Ready = false
		}
		deck := room.initGame()
		// Handle start game
		return createMsg(room.ID, START, deck)

	}
	return createMsg(room.ID, INIT, "Not all players are ready!")

}

func (room *Room) initGame() string {
	fullDeck := getFullDeck()
	// Populate hands
	// Populate deck
	// Wait for ready
	handLimit := 4
	players := room.Players
	room.State = CHOOSE_CARD
	if len(players) == 10 {
		handLimit = 10
	}
	start := 0
	for _, player := range players {
		dealtHand := fullDeck[start : start+handLimit]
		slices.Sort(dealtHand)
		player.Hand = dealtHand
		room.Hands[player.ID] = dealtHand
		start += handLimit
	}
	var deck [][]int
	for i := start; i < start+4; i++ {
		row := []int{fullDeck[i]}
		deck = append(deck, row)
	}
	room.Deck = deck
	return deckTostring(deck)
}

// GAME HANDLERS

func (room *Room) play(msgReq *MessageReq, p *Player, hub *Hub) *Message {
	sel := msgReq.Card
	// remark = fmt.Sprintf("%v played card.", p.Name)
	room.Select[p.ID] = sel
	if len(room.Select) == len(room.Players) {
		// Start counting
		if room.Ticker == nil {
			room.Ticker = time.NewTicker(1 * time.Second)
			go room.timer(room.ID, 5, PLAY, p, hub)
		} else {
			room.Ticker.Stop()
			room.Ticker = time.NewTicker(1 * time.Second)
			go room.timer(room.ID, 5, PLAY, p, hub)
		}
		return createMsg(room.ID, CALCULATING, "Processing...")
	}

	return createMsg(room.ID, CHOOSE_CARD, fmt.Sprintf("%v played card.", p.Name))
}

func (room *Room) processCards(p *Player, hub *Hub) *Message {
	//reset ticker
	room.Ticker = nil
	for ID, sel := range room.Select {
		room.Played[sel] = ID
	}
	// Remove played card from players' hands
	removePlayed(room.Played, room.Players, room)
	hub.Broadcast <- createMsg(room.ID, PROCESS, showPlayed(room))
	sortedCards := getSortedCards(room.Played)
	if isSmallest(sortedCards[0], room) {
		room.State = CHOOSE_ROW
		// Return
		return createMsg(room.ID, CHOOSE_ROW, "Choose Row!")
	} else {
		// Play cards
		// game.State = CHOOSE_CARD
		// state = CHOOSE_CARD
		// playCards
		// remark = "Played cards!"
		room.playCards()
		if len(p.Hand) == 0 {
			// end game
			return createMsg(room.ID, endGame(room), "ENDED!")
		}
	}
	return createMsg(room.ID, CHOOSE_CARD, deckTostring(room.Deck))
}

func (room *Room) rows(sel int, card int, p *Player) *Message {
	log.Println("Doing row stuff!")

	// 	remark = fmt.Sprintf("%v selected row: %v!", p.ID, sel)
	// 	// EAT POINTS
	row := room.Deck[sel]
	p.Score += getScore(row)
	room.Scores[p.ID] = p.Score
	room.Deck[sel] = []int{card}
	delete(room.Played, card)
	room.State = CALCULATING
	room.playCards()
	if len(p.Hand) == 0 {
		// end game
		return createMsg(room.ID, endGame(room), "ENDED!")
	}
	return createMsg(room.ID, CHOOSE_ROW, fmt.Sprintf("%v selected row: %v!", p.Name, sel))
	// 	// updatePlayers(players)
}

func (room *Room) playCards() {
	//Check for nearest
	players := room.Players
	deck := room.Deck
	// Calculation
	for _, card := range getSortedCards(room.Played) {
		nearestPos := getNearest(card, deck)
		if len(deck[nearestPos]) == 5 {
			// BUSTED
			players[room.Played[card]].Score += getScore(deck[nearestPos])
			deck[nearestPos] = []int{}
		}
		deck[nearestPos] = append(deck[nearestPos], card)
	}
	room.Played = make(map[int]string)
	room.Select = make(map[string]int)

}

func endGame(room *Room) int {
	// Check if any player crosse 66 points
	// -> If no, init new round
	// -> If yes, go back to lobby

	var exceed bool
	for _, p := range room.Players {
		if p.Score >= 5 {
			exceed = true
			break
		}
	}

	if exceed {
		// WE GOT A WINNER!
		room.State = INIT
		return GAME_END
		// go back to lobby
	} else {
		room.State = ROUND_END
		// start new round
		return ROUND_END
	}

}

// TESTER HANDLERS
func showhands(room *Room) {
	for _, p := range room.Players {
		p.Message <- createMsg(room.ID, 100, arrIntToString(p.Hand))
	}
}

func testConn(room *Room) {
	log.Println("CONNECTION STATUS:")
	for _, p := range room.Players {
		log.Println(p.Conn)
	}
}
