package ws

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func newPlayer(id string, name string, roomId string) *Player {
	return &Player{
		RoomID:  roomId,
		ID:      id,
		Name:    name,
		Hand:    []int{},
		Play:    -1,
		HP:      100,
		Ready:   false,
		Message: make(chan *Message, 10),
		End:     0,
		DamageReport: &DamageReport{
			Mushrooms:      0,
			Damage:         0,
			RoundMushrooms: 0,
			RoundDamage:    0,
			MushroomTypes:  []int{},
		},
	}
}

func newPlayerRes(player *Player) PlayerRes {
	return PlayerRes{
		ID:           player.ID,
		Name:         player.Name,
		Play:         player.Play,
		Hand:         player.Hand,
		HP:           player.HP,
		Ready:        player.Ready,
		End:          player.End,
		DamageReport: *player.DamageReport,
	}

}

func checkReady(players map[string]*Player) bool {
	for _, p := range players {
		if !p.Ready {
			return false
		}
	}
	return true
}

func getRandID(rooms map[string]*Room) string {
	var fuse int

Again:
	fuse++
	if fuse == 100 {
		log.Print("Fuse broke")
	}
	newId := ""
	for i := 0; i < 4; i++ {
		randN := (rand.Intn(10))
		newId += strconv.Itoa(randN)
	}
	for id := range rooms {
		if newId == id {
			goto Again
		}
	}
	return newId
}

func getPlayerID(roomID string) string {
	c := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	newId := ""
	for i := 0; i < 32; i++ {
		randN := (rand.Intn(62))
		newId += string(c[randN])
	}
	return newId + roomID
}

func arrIntToString(arr []int) string {
	var res string
	for _, x := range arr {
		res += strconv.Itoa(x)
		res += ", "
	}
	return res
}

func deckTostring(arr [][]int) string {
	if len(arr) == 0 {
		return "[]"
	}
	var rows []string
	for _, x := range arr {
		var vals string
		for _, y := range x {
			vals += strconv.Itoa(y) + ", "
		}
		rows = append(rows, vals)
	}
	return fmt.Sprintf("Row 1: [%v] Row 2: [%v] Row 3: [%v] Row 4: [%v]", rows[0], rows[1], rows[2], rows[3])
}

func damage(row []int, mushrooms map[int]Mushroom) int {
	var dmg int
	for _, x := range row {

		// Default all the undefined mushrooms to shiitake,
		// Ignore invalid input first, will fix later.
		mush, ok := mushrooms[x]
		if !ok {
			mush = mushrooms[0]
		}
		dmg += mush.Damage
	}
	return dmg
}

func getFullDeck() []int {
	var deck []int
	for i := 1; i <= 100; i++ {
		deck = append(deck, i)
	}

	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func isSmallest(smallest int, room *Room) bool {
	for _, row := range room.Deck {
		if smallest > row[len(row)-1] {
			return false
		}
	}
	return true
}

func removePlayed(played map[int]string, players map[string]*Player, room *Room) {
	for card, id := range played {
		p := players[id]
		// reset player ready flag
		p.Ready = false
		var newHand []int
		for _, c := range p.Hand {
			if c != card {
				newHand = append(newHand, c)
			}
		}
		p.Hand = newHand
		room.Hands[id] = p.Hand
	}
}

func getSortedCards(played map[int]string) []int {
	var cards []int
	smallest := -1
	sorted := true
	for card := range played {
		if card < smallest && sorted {
			sorted = false
		}
		smallest = card
		cards = append(cards, card)
	}
	if !sorted {
		slices.Sort(cards)
	}
	return cards
}

func showPlayed(room *Room) string {
	sorted := getSortedCards(room.Played)
	var out string
	for _, card := range sorted {
		out += fmt.Sprintf("%v:%v ", room.Players[room.Played[card]].Name, card)

	}
	return out
}

func getNearest(card int, deck [][]int) int {

	min := 1000
	var pos int
	for i, row := range deck {
		if len(row) == 0 {
			//bypass
			return i
		}
		tail := row[len(row)-1]
		if card < tail {
			continue
		}
		if (card - tail) < min {
			min = card - tail
			pos = i
		}
	}
	return pos
}

// func updatePlayers(players map[string]*Player) {
// 	for _, p := range players {
// 		p.Message <- &Message{
// 			STATE: CHOOSE_CARD,
// 			Remark: arrIntToString(p.Hand),
// 		}
// 	}
// }

func isEither(state string, checker ...string) bool {
	for _, x := range checker {
		if state == x {
			return true
		}
	}
	return false
}

func playersArr(players map[string]*Player) []*PlayerDisplay {
	var res []*PlayerDisplay

	for _, p := range players {
		player := &PlayerDisplay{
			Name:  p.Name,
			HP:    p.HP,
			Ready: p.Ready,
		}
		res = append(res, player)
	}
	return res
}

func getPlayed(r *Room) map[string]int {
	played := make(map[string]int)
	for p, id := range r.Played {
		played[r.Players[id].Name] = p
	}
	return played
}

func getChooser(r *Room) string {
	if r.Chooser == "" {
		return ""
	}
	return r.Players[r.Chooser].Name
}

// func generateToken(id string) string {
// 	secretKey := "secret"
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
// 		Name: id,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			Issuer:    id,
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
// 		},
// 	})

// 	ss, err := token.SignedString([]byte(secretKey))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return ss
// }

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (room *Room) timer(roomID string, i int, state int, p *Player, hub *Hub) {
	ticker := room.Ticker
	if state == IDLE {
		ticker = time.NewTicker(1 * time.Second)
	}
	for {
		select {
		case <-ticker.C:
			if i == 0 {
				if state == PLAY {
					log.Println("PLAY!")

					hub.Broadcast <- createMsg(roomID, COUNT, strconv.Itoa(i))
					room.gameState(&MessageReq{Action: PROCESS}, p, hub)
				}

				if state == IDLE {
					delete(hub.Rooms, roomID)
					log.Println("Room deleted")
				}
				return
			}
			if state == PLAY {
				hub.Broadcast <- createMsg(roomID, COUNT, strconv.Itoa(i))
			}
			i--

		case player := <-hub.Register:
			if state == IDLE {
				if player.RoomID == roomID {
					return
				}
			}
		}

	}
}

func createMsg(roomID string, state int, remark string) *Message {
	return &Message{
		RoomID: roomID,
		State:  state,
		Remark: remark,
	}
}

func str(n int) string {
	return strconv.Itoa(n)
}

func getHighestDamage(row []int) int {
	highest := 0
	for _, x := range row {
		if x > highest {
			highest = x
		}
	}
	return highest
}

func addMush(mush []int, add []int, mushrooms map[int]Mushroom) []int {
	if len(mush) == 0 {
		mush = []int{add[0]}
	}
	add = add[1:]

	for _, a := range add {
		for i, m := range mush {
			mushA, okA := mushrooms[a]
			mushM, okM := mushrooms[m]
			if !okA {
				mushA = mushrooms[0]
			}
			if !okM {
				mushM = mushrooms[0]
			}
			if mushA.Name == mushM.Name {
				break
			}
			if mushA.Damage <= mushM.Damage {
				mush = append(mush[:i], append([]int{a}, mush[i:]...)...)
				break
			}
			mush = append(mush, a)
		}
	}
	return mush
}

func (r *Room) CheckConn() {
	// Check everyone's connection
	for _, p := range r.Players {
		if p.Conn != nil {
			return
		}
	}

	// If everyone is disconnected, start idle timer
	// Once the timer is up, delete the room

}

func getMushrooms() map[int]Mushroom {
	mush := make(map[int]Mushroom)

	// Default to shiitake
	mush[0] = mushroomsLib[4]

	// Special mushroom. Will work on randomizing later

	mush[11] = mushroomsLib[5]
	mush[27] = mushroomsLib[8]
	mush[39] = mushroomsLib[9]
	mush[51] = mushroomsLib[2]

	return mush
}
