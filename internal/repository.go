package internal

import (
	"context"
)

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	query := "INSERT INTO users(name) VALUES ($1) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Name).Scan(&lastInsertId)
	if err != nil {
		return &User{}, err
	}
	user.ID = int64(lastInsertId)
	return user, nil
}

func (r *repository) CreateRoom(ctx context.Context, room *Room) (*Room, error) {
	var lastInsertId int
	// generate room code

	query := "INSERT INTO rooms(code, players) VALUES ($1, $2) returning id"
	err := r.db.QueryRowContext(ctx, query, room.Code, room.Players).Scan(&lastInsertId)
	if err != nil {
		return &Room{}, err
	}
	room.ID = int64(lastInsertId)
	return room, nil
}
