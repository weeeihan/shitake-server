package ws

import (
	"fmt"
	"log"
	"runtime"
	"slices"
	"time"
)

func (room *Room) gameState(msgReq *MessageReq, p *Player, hub *Hub) {

	var msg *Message
	log.Print(msgReq)
	switch msgReq.Action {

	case LEAVE:
		// player leaves room
		msg = p.leaveRoom(room)
		if len(room.Players) == 0 {
			delete(hub.Rooms, room.ID)
		}

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
		// msg = createMsg(room.ID, PROCESS, "START PROCESSING CARDS")
		msg = room.processCards(p, hub)

	case NEXT_PLAY:
		room.State = CHOOSE_CARD
		msg = createMsg(room.ID, CHOOSE_CARD, "CHOOSE CARD")
		p.Message <- msg
		return

	case ROW:
		msg = room.rows(msgReq.Row, msgReq.Card, p, hub)

	case ROUND_END:
		p.End = 1
		msg = createMsg(room.ID, ROUND_END, "Player ended round!")

	case GAME_END:
		p.End = 2
		msg = createMsg(room.ID, GAME_END, "Player ended game!")

	case COUNT:

	case STOPCOUNT:

	case CHAT:
		msg = createMsg(room.ID, CHAT, "blabla")

	case PING:
		msg = createMsg(room.ID, PING, fmt.Sprintf("Number of goroutines: %v", runtime.NumGoroutine()))
		// testConn(room)

	}
	// showhands(room)
	// log.Print(room)
	hub.Broadcast <- msg

}

// PRE-GAME HANDLERS

func (player *Player) leaveRoom(room *Room) *Message {
	if room.State == GAME_END {
		// cancel all ready
		for _, p := range room.Players {
			p.Ready = false
			p.End = 0
			p.DamageReport.Damage = 0
			p.DamageReport.Mushrooms = 0
		}
		room.State = INIT
		return createMsg(room.ID, LOBBY, "Back to lobby!")
	}
	// close(player.Message)
	delete(room.Players, player.ID)
	player.Conn.Close()
	// Delete room if nobody in it

	return createMsg(room.ID, PLAYER_LEFT, fmt.Sprintf("Player %v left the room!", player.Name))
}

func (player *Player) ready(room *Room) *Message {
	log.Print("Get ready?")
	// log.Println("YO")
	player.Ready = true

	if checkReady(room.Players) {
		if room.State == INIT {

			return createMsg(room.ID, ALREADY, "All ready to start!")
			// Call init game
		}

		if room.State == ROUND_END || room.State == GAME_END {
			startGame(room)
			return createMsg(room.ID, START, "Start game")
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
	if room.State == INIT || room.State == ROUND_END || room.State == GAME_END {
		if checkReady(room.Players) {
			for _, p := range room.Players {
				p.DamageReport.RoundDamage = 0
				p.DamageReport.RoundMushrooms = 0
				if room.State == GAME_END || room.State == INIT {
					p.HP = 100
				}
				p.End = 0
				p.Ready = false
			}
			room.State = CHOOSE_CARD

			deck := room.initGame()
			// Handle start game
			return createMsg(room.ID, START, deck)

		}
	}
	return createMsg(room.ID, INIT, "Not all players are ready!")

}

func (room *Room) initGame() string {
	fullDeck := getFullDeck()
	// Populate hands
	// Populate deck
	// Wait for ready
	handLimit := 3
	players := room.Players
	if len(players) == 10 {
		handLimit = 10
	}
	start := 0
	for _, player := range players {
		dealtHand := fullDeck[start : start+handLimit]
		slices.Sort(dealtHand)
		player.Hand = dealtHand
		// player.Hand = []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
		room.Hands[player.ID] = dealtHand
		start += handLimit
	}
	var deck [][]int
	for i := start; i < start+4; i++ {
		row := []int{fullDeck[i]}
		deck = append(deck, row)
	}
	// TESTING
	room.Deck = deck
	// room.Deck = [][]int{{5, 6}, {2, 7}, {3, 8}, {4}}
	return deckTostring(deck)
}

// GAME HANDLERS

func (room *Room) play(msgReq *MessageReq, p *Player, hub *Hub) *Message {
	sel := msgReq.Card
	// remark = fmt.Sprintf("%v played card.", p.Name)
	room.Select[p.ID] = sel
	p.Play = sel
	p.Ready = true
	if len(room.Select) == len(room.Players) {
		// room.gameState(&MessageReq{Action: PROCESS}, p, hub)

		// Start counting
		if room.Ticker == nil {
			room.Ticker = time.NewTicker(1 * time.Second)
			go room.timer(room.ID, 3, PLAY, p, hub)
		}
		return createMsg(room.ID, CALCULATING, "Processing...")
	}

	return createMsg(room.ID, CHOOSE_CARD, fmt.Sprintf("%v played card.", p.Name))
}

func (room *Room) processCards(p *Player, hub *Hub) *Message {
	//reset ticker
	room.Ticker = nil
	room.State = PROCESS

	// reset moves
	room.Moves = [][]string{}
	// reset played cards
	room.Played = make(map[int]string)
	for ID, sel := range room.Select {
		room.Played[sel] = ID
		room.Players[ID].Play = -1
	}
	// Remove played card from players' hands
	removePlayed(room.Played, room.Players, room)
	// hub.Broadcast <- createMsg(room.ID, PROCESS, showPlayed(room))
	sortedCards := getSortedCards(room.Played)
	// hub.Broadcast <- createMsg(room.ID, PROCESS, "PROCESSING CARD")
	if isSmallest(sortedCards[0], room) {
		room.State = CHOOSE_ROW
		room.Chooser = room.Played[sortedCards[0]]
		// Return
		return createMsg(room.ID, CHOOSE_ROW, "CHOOSE ROW")
	} else {
		// Play cards
		// game.State = CHOOSE_CARD
		// state = CHOOSE_CARD
		// playCards
		// remark = "Played cards!"
		room.playCards(hub)
		if len(p.Hand) == 0 {
			// end game
			return createMsg(room.ID, endGame(room), "ENDED!")
		}
	}
	return createMsg(room.ID, PROCESS, deckTostring(room.Deck))
}

func (room *Room) rows(sel int, card int, p *Player, hub *Hub) *Message {
	if room.State == CHOOSE_ROW {
		room.Chooser = ""
		// 	remark = fmt.Sprintf("%v selected row: %v!", p.ID, sel)
		// 	// EAT POINTS
		row := room.Deck[sel]
		p.eat(row, room)
		room.HPs[p.ID] = p.HP
		room.Deck[sel] = []int{}
		room.State = CALCULATING
		room.playCards(hub)
		if len(p.Hand) == 0 {
			// end game
			return createMsg(room.ID, endGame(room), "ENDED!")
		}

	}
	return createMsg(room.ID, PROCESS, fmt.Sprintf("%v selected row: %v!", p.Name, sel))
	// 	// updatePlayers(players)
}

func (room *Room) playCards(hub *Hub) {
	//Check for nearest
	players := room.Players
	deck := room.Deck
	var moves [][]string

	// Calculation
	for _, card := range getSortedCards(room.Played) {

		nearestPos := getNearest(card, deck)
		id := room.Played[card]
		moves = append(moves, []string{room.Players[id].Name, str(card), str(nearestPos), str(len(deck[nearestPos]))})
		if len(deck[nearestPos]) == 5 {
			// BUSTED
			players[room.Played[card]].eat(deck[nearestPos], room)
			deck[nearestPos] = []int{}
		}

		deck[nearestPos] = append(deck[nearestPos], card)
		// hub.Broadcast <- createMsg(room.ID, PLAY, "SNAPSHOT")
	}

	room.Moves = moves
	// Reset selections
	room.Select = make(map[string]int)

	// room.State = CHOOSE_CARD

}

func (player *Player) eat(row []int, room *Room) {
	// update damage report
	damage := damage(row, room.Mushrooms)
	player.HP -= damage
	dr := player.DamageReport
	dr.Damage += damage
	dr.Mushrooms++
	dr.RoundMushrooms++
	dr.RoundDamage += damage
	dr.MushroomTypes = addMush(dr.MushroomTypes, row, room.Mushrooms)
}

func endGame(room *Room) int {
	// Check for casualties
	// -> If no, init new round
	// -> If yes, go back to lobby

	var end bool
	for _, p := range room.Players {
		if p.HP <= 0 {
			end = true
			break
		}
	}

	if end {
		// WE GOT A WINNER!
		room.State = GAME_END
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
