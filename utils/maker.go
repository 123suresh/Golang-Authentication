package utils

import (
	"time"

	"example.com/dynamicWordpressBuilding/internal/model"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*model.Payload, error)
}
