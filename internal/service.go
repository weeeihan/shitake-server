package internal

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "secret"
)

type service struct {
	Repository
	timeout time.Duration
}

type MyJWTClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u := &User{
		Name: req.Name,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:   strconv.Itoa(int(r.ID)),
		Name: r.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(r.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &CreateUserRes{}, err
	}

	return &CreateUserRes{accessToken: ss, Name: r.Name, ID: strconv.Itoa(int(r.ID))}, nil
}

func (s *service) CreateRoom(c context.Context, req *CreateRoomReq) (*CreateRoomRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	r := &Room{
		Code:    1234,
		Players: []string{req.Player},
	}

	r, err := s.Repository.CreateRoom(ctx, r)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return &CreateRoomRes{}, err
	}

	return &CreateRoomRes{ID: strconv.Itoa(int(r.ID)), Code: r.Code, Players: r.Players}, nil
}
