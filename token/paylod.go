package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token is expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID        `json:"id"`
	Username  string           `json:"username"`
	Subject   string           `json:"subject"`
	Audience  jwt.ClaimStrings `json:"aud,omitempty"`
	IssuedAt  time.Time        `json:"issued_at"`
	ExpiredAt time.Time        `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredAt), nil
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Now()), nil
}

func (payload *Payload) GetIssuer() (string, error) {
	return payload.Username, nil
}

func (payload *Payload) GetSubject() (string, error) {
	return payload.Subject, nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return payload.Audience, nil
}
