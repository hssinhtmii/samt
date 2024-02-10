package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	Id          uuid.UUID `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	IssuedAt    time.Time `json:"issuedat"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

func NewPayload(phonenumber string, duration time.Duration) (*Payload, error) {
	TokenId, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	Payload := &Payload{
		Id:          TokenId,
		PhoneNumber: phonenumber,
		IssuedAt:    time.Now(),
		ExpiredAt:   time.Now().Add(duration),
	}
	return Payload, nil
}
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.New("توکن منقضی شده است")
	}
	return nil
}
