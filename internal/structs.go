package internal

import "time"

type Room struct {
	ID      int64    `json:"id" db:"id"`
	Players []string `json:"players" db:"players"`
	Code    int      `json:"code" db:"code"`
	Deck    []string `json:"deck" db:"deck"`
	Scores  []int    `json:"scores" db:"scores"`
	timer   *time.Ticker
}

type CreateRoomReq struct {
	Player string `json:"player" db:"player"`
}

type CreateRoomRes struct {
	ID      string   `json:"id" db:"id"`
	Code    int      `json:"code" db:"code"`
	Players []string `json:"players" db:"players"`
}

type User struct {
	ID    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Hand  string `json:"hand" db:"hand"`
	Score int    `json:"score" db:"score"`
	Room  string `json:"room" db:"room"`
}

type CreateUserReq struct {
	Name string `json:"name" db:"name"`
}

type CreateUserRes struct {
	accessToken string
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
}

type repository struct {
	db DBTX
}
