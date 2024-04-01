package ws

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
)

// enum for lobby
var (
	NEW_PLAYER_JOINED int = 0
	PLAYER_LEFT       int = 1

	// enum for player Message
	READY   int = 2
	UNREADY int = 3
	START   int = 4
	PLAY    int = 5
	UNPLAY  int = 6
	ROW     int = 7

	// enum for game state
	INIT        int = 8
	CHOOSE_CARD int = 9
	CHOOSE_ROW  int = 10
	CALCULATING int = 11
	ROUND_END   int = 12
	GAME_END    int = 13
	LOBBY       int = 14
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
	GameID string `json:"roomId"`
	State  int    `json:"state"`
	Remark string `json:"remark"`
}

/*
roomStates : LOBBY, CHOOSE_CARD, CHOOSE_ROW, ROUND_END, PAUSE
*/

type Room struct {
	ID      string `json:"id"`
	Players map[string]*Player
	State   int `json:"state"`
}

type RoomRes struct {
	ID      string           `json:"id"`
	State   int              `json:"state"`
	Players []*PlayerDisplay `json:"players"`
}

type GameRes struct {
	ID      string           `json:"id"`
	Deck    [][]int          `json:"deck"`
	State   int              `json:"state"`
	Players []*PlayerDisplay `json:"players"`
}

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
	Games      map[string]*Game
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
