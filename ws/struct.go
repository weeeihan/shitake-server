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
	Conn    *websocket.Conn
	Message chan *Message
	ID      string `json:"id"`
	RoomID  string `json:"roomId"`
	Name    string `json:"name"`
	Hand    []int  `json:"hand"`
	Score   int    `json:"score"`
	Play    int    `json:"play"`
	Ready   bool   `json:"ready"`
}

// {"action": "", "card": num, "row": num}

type MessageReq struct {
	Action int `json:"action"`
	Card   int `json:"card"`
	Row    int `json:"row"`
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

	Scores map[string]int `json:"scores"`
	Pause  bool           `json:"pause"`
	Ready  int            `json:"ready"`
	Ticker *time.Ticker
}

type RoomRes struct {
	ID      string           `json:"id"`
	State   int              `json:"state"`
	Deck    [][]int          `json:"deck"`
	Players []*PlayerDisplay `json:"players"`
	Played  map[string]int   `json:"played"`
	Chooser string           `json:"chooser"`
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
	ID    string `json:"id"`
	Name  string `json:"name"`
	Hand  []int  `json:"hand"`
	Score int    `json:"score"`
	Ready bool   `json:"ready"`
}

type MyJWTClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type PlayerDisplay struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Ready bool   `json:"ready"`
}
