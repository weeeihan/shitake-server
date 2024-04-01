package internal

import (
	"context"
	"database/sql"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	CreateRoom(ctx context.Context, room *Room) (*Room, error)
	// GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error)
	CreateRoom(c context.Context, req *CreateRoomReq) (*CreateRoomRes, error)
	// Login(c context.Context) (*LoginUserRes, error)
}

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
