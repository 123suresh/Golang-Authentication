package utils

import "time"

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
}