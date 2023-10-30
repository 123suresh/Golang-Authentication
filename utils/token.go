package utils

import (
	"errors"
	"time"

	"example.com/dynamicWordpressBuilding/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTMaker struct {
	secretKey string
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func NewPayload(username string, duration time.Duration) (*model.Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &model.Payload{
		ID:        tokenID,
		Email:     username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	tokenPayload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenPayload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

//if you have to use CreateToken function then first make Interface We have make as maker.go
//then initialize interface like below and you will be able to use
func NewTokenMaker() Maker {
	return &JWTMaker{}
}
