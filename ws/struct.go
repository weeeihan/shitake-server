package ws

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
)

// enum for lobby

// For exporting to client
var gamestates = map[string]int{
	"NEW_PLAYER_JOINED": NEW_PLAYER_JOINED,
	"REGISTERED":        REGISTERED,
	"PLAYER_LEFT":       PLAYER_LEFT,
	"LEAVE":             LEAVE,
	"READY":             READY,
	"UNREADY":           UNREADY,
	"ALREADY":           ALREADY,
	"START":             START,
	"PLAY":              PLAY,
	"PROCESS":           PROCESS,
	"ROW":               ROW,
	"INIT":              INIT,
	"CHOOSE_CARD":       CHOOSE_CARD,
	"CHOOSE_ROW":        CHOOSE_ROW,
	"CALCULATING":       CALCULATING,
	"ROUND_END":         ROUND_END,
	"GAME_END":          GAME_END,
	"LOBBY":             LOBBY,
	"PING":              PING,
	"COUNT":             COUNT,
	"ROW_SELECTED":      ROW_SELECTED,
	"NEXT_PLAY":         NEXT_PLAY,
}

var (
	NEW_PLAYER_JOINED int = 1
	REGISTERED        int = 2
	PLAYER_LEFT       int = 3
	LEAVE             int = 4

	// enum for player Message
	READY   int = 5
	UNREADY int = 6
	ALREADY int = 7
	START   int = 8
	PLAY    int = 9
	PROCESS int = 10
	ROW     int = 11

	// enum for game state
	INIT         int = 12
	CHOOSE_CARD  int = 13
	CHOOSE_ROW   int = 14
	NEXT_PLAY    int = 25
	ROW_SELECTED int = 15
	CALCULATING  int = 16
	ROUND_END    int = 17
	GAME_END     int = 18
	LOBBY        int = 19
	COUNT        int = 20
	STOPCOUNT    int = 21
	RESET        int = 22
	PING         int = 23

	CHAT         int = 23
	DISCONNECTED int = 24

	GETCARD int = 100
)

type Player struct {
	Conn         *websocket.Conn
	Message      chan *Message
	ID           string       `json:"id"`
	RoomID       string       `json:"roomId"`
	Name         string       `json:"name"`
	Hand         []int        `json:"hand"`
	HP           int          `json:"hp"`
	Play         int          `json:"play"`
	Ready        bool         `json:"ready"`
	DamageReport DamageReport `json:"damageReport"`
}

// {"action": "", "card": num, "row": num}

type MessageReq struct {
	Action int `json:"action"`
	Card   int `json:"card"`
	Row    int `json:"row"`
}

type DamageReport struct {
	Mushrooms      []int `json:"mushrooms"`
	Damage         int   `json:"damageTaken"`
	RoundMushrooms []int `json:"roundMushrooms"`
	RoundDamage    int   `json:"roundDamage"`
}

type Message struct {
	// ID       string `json:"id"`
	// Selected int    `json:"selected"`
	RoomID string `json:"roomId"`
	State  int    `json:"state"`
	Remark string `json:"remark"`
}

/*
roomStates : LOBBY, CHOOSE_CARD, CHOOSE_ROW, ROUND_END, PAUSE
*/

type Room struct {
	ID      string  `json:"id"`
	Deck    [][]int `json:"deck"`
	Players map[string]*Player
	State   int `json:"state"`

	// map[card]Playedbywhom
	Played map[int]string `json:"played"`
	Select map[string]int `json:"select"`

	// map[playerID]Hands
	Hands map[string][]int `json:"hands"`
	//map[playerID]Scores

	//Row chooser
	Chooser string `json:"chooser"`

	HPs   map[string]int `json:"hps"`
	Pause bool           `json:"pause"`
	Ready int            `json:"ready"`

	// For animation
	Moves  [][]string `json:"moves"`
	Ticker *time.Ticker
}

type RoomRes struct {
	ID      string           `json:"id"`
	State   int              `json:"state"`
	Deck    [][]int          `json:"deck"`
	Players []*PlayerDisplay `json:"players"`
	Played  map[string]int   `json:"played"`
	Chooser string           `json:"chooser"`
	Moves   [][]string       `json:"moves"`
}

// type GameRes struct {
// 	ID      string           `json:"id"`
// 	Deck    [][]int          `json:"deck"`
// 	State   int              `json:"state"`
// 	Players []*PlayerDisplay `json:"players"`
// }

type Game struct {
	ID      string  `json:"id"`
	Deck    [][]int `json:"deck"`
	Players map[string]*Player
	State   int `json:"state"`
	// map[card]Playedbywhom
	Played map[int]string `json:"played"`

	Select map[string]int `json:"select"`

	// map[playerID]Hands
	Hands map[string][]int `json:"hands"`
	//map[playerID]Scores
	Scores map[string]int `json:"scores"`
	Pause  bool           `json:"pause"`
	Ready  int            `json:"ready"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Player
	Unregister chan *Player
	Broadcast  chan *Message
}

type Handler struct {
	hub *Hub
}

// type GameRes struct {
// 	ID        string           `json:"id"`
// 	Deck      string           `json:"deck"`
// 	GameState string           `json:"gameState"`
// 	Players   []string         `json:"players"`
// 	Hands     map[string][]int `json:"hands"`
// 	Scores    map[string]int   `json:"scores"`
// }

type PlayerRes struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Play         int          `json:"play"`
	Hand         []int        `json:"hand"`
	HP           int          `json:"hp"`
	Ready        bool         `json:"ready"`
	DamageReport DamageReport `json:"damageReport"`
}

type MyJWTClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type PlayerDisplay struct {
	Name  string `json:"name"`
	HP    int    `json:"hp"`
	Ready bool   `json:"ready"`
}

type Mushroom struct {
	Name        string `json:"name"`
	Damage      int    `json:"damage"`
	Description string `json:"desc"`
	Color       string `json:"color"`
}

var mushrooms = map[int]Mushroom{
	1: {
		Name:        "White button",
		Damage:      2,
		Description: "Agaricus bisporus, commonly known as the cultivated mushroom, is a basidiomycete mushroom native to grasslands in Eurasia and North America. It is cultivated in more than 70 countries and is one of the most commonly and widely consumed mushrooms in the world.",
		Color:       "white",
	},
	2: {
		Name:        "Enoki",
		Damage:      2,
		Description: "Flammulina filiformis is a species of edible agaric in the family Physalacriaceae. It is widely cultivated in East Asia, and well known for its role in Japanese and Chinese cuisine.",
		Color:       "white",
	},
	3: {
		Name:        "Morel",
		Damage:      2,
		Description: "Morchella, the true morels, is a genus of edible sac fungi closely related to anatomically simpler cup fungi in the order Pezizales. These distinctive fungi have a honeycomb appearance due to the network of ridges with pits composing their caps.",
		Color:       "black",
	},
	4: {
		Name:        "Shiitake",
		Damage:      2,
		Description: "Lentinula edodes is a species of edible mushroom native to East Asia, which is cultivated and consumed in many Asian countries. It is considered a medicinal mushroom in some forms of traditional medicine.",
		Color:       "brown",
	},
	5: {
		Name:        "Oyster",
		Damage:      2,
		Description: "Pleurotus ostreatus, the oyster mushroom, is a common edible mushroom. It was first cultivated in Germany as a subsistence measure during World War I and is now grown commercially around the world for food.",
		Color:       "white",
	},
	6: {
		Name:        "Porcini",
		Damage:      2,
		Description: "Boletus edulis is a basidiomycete fungus, and the type species of the genus Boletus. Widely distributed in the Northern Hemisphere across Europe, Asia, and North America, it does not occur naturally in the Southern Hemisphere, although it has been introduced to southern Africa, Australia, and New Zealand.",
		Color:       "brown",
	},
	7: {
		Name:        "Chanterelle",
		Damage:      2,
		Description: "Cantharellus cibarius, commonly known as the chanterelle, golden chanterelle or girolle, is a fungus. It is probably the best known species of the genus Cantharellus, if not the entire family of Cantharellaceae.",
		Color:       "yellow",
	},
	8: {
		Name:        "Lion's Mane",
		Damage:      2,
		Description: "Hericium erinaceus is a species of tooth fungus in the family Hericiaceae. It is native to North America, Europe, and Asia. It can be mistaken for other species of Hericium, all popular edibles, which grow across the same range.",
		Color:       "white",
	},
	9: {
		Name:        "Reishi",
		Damage:      2,
		Description: "Ganoderma lucidum is a species of bracket fungus, and the type species of the genus Ganoderma. It lives on deadwood, especially dead trees, and is generally considered to be a saprotroph, rather than a parasite.",
		Color:       "red",
	},
	10: {
		Name:        "Maitake",
		Damage:      2,
		Description: "Grifola frondosa is a polypore mushroom that grows in clusters at the base of trees, particularly oaks. The mushroom is commonly known among English speakers as hen of the woods, ram's head, and sheep's head.",
		Color:       "white",
	},
}
