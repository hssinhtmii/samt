package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

type PasetoMaker struct {
	paseto      *paseto.V2
	SymetricKey []byte
}

func NewPasetoMaker(Symetrickey string) (Make_Token, error) {
	if len(Symetrickey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("اندازه کلید نا معتبر است. کلید باید %d باشد", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:      paseto.NewV2(),
		SymetricKey: []byte(Symetrickey),
	}
	return maker, nil
}

func (makepaseto *PasetoMaker) CreateToken(PhoneNumber string, duration time.Duration) (string, *Payload, error) {
	Payload, err := NewPayload(PhoneNumber, duration)

	if err != nil {
		return "",Payload, nil
	}
	token, err := makepaseto.paseto.Encrypt(makepaseto.SymetricKey, Payload, nil)
	return token, Payload, err
}
func (makepaseto *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := makepaseto.paseto.Decrypt(token, makepaseto.SymetricKey, payload, nil)

	if err != nil {
		return nil, errInvalid
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
