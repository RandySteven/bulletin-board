package securities

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("supersecret")

type JWTClaim struct {
	UserID uint64
	RoleID uint64
	jwt.RegisteredClaims
}
