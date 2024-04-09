package ws

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/websocket"
)

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

func getScore(row []int) int {
	var score int
	for range row {
		score++
	}
	return score
}

func getFullDeck() []int {
	var deck []int
	for i := 1; i <= 104; i++ {
		deck = append(deck, i)
	}

	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func isSmallest(smallest int, room *Game) bool {
	for _, row := range room.Deck {
		if smallest > row[len(row)-1] {
			return false
		}
	}
	return true
}

func removePlayed(played map[int]string, players map[string]*Player, room *Game) {
	for card, id := range played {
		p := players[id]
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
	for card := range played {
		cards = append(cards, card)
	}
	slices.Sort(cards)
	return cards
}

func getNearest(card int, deck [][]int) int {
	min := 1000
	var pos int
	for i, row := range deck {
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

func resetScore(players map[string]*Player) {
	for _, p := range players {
		p.Score = 0
	}
}

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
			Score: p.Score,
			Ready: p.Ready,
		}
		res = append(res, player)
	}
	return res
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
