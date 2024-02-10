package token

import "time"

type Make_Token interface {
	CreateToken(PhoneNumber string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}