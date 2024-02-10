package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const KeyLen = 32

var (
	errExpired = errors.New("توکن منقضی شده است")
	errInvalid = errors.New("توکن اشتباه است")
)

type JWTMaker struct {
	SecretKey string
}

func NewJWTMaker(secretkey string) (Make_Token, error) {

	if len(secretkey) < KeyLen {
		return nil, fmt.Errorf("key size is invalid... must be %d", KeyLen)
	}
	return &JWTMaker{secretkey}, nil
}

func (maker *JWTMaker) CreateToken(PhoneNumber string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(PhoneNumber, duration)

	if err != nil {
		return "", payload, err
	}

	JwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := JwtToken.SignedString([]byte(maker.SecretKey))
	return token, payload, err
}
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errInvalid
		}
		return []byte(maker.SecretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		valerr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(valerr.Inner, errExpired) {
			return nil, errExpired
		}
		return nil, errInvalid
	}
	Payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, errInvalid
	}

	return Payload, nil
}
