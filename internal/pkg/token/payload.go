package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	Sub      uuid.UUID `json:"sub"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}
