package claims

import "github.com/golang-jwt/jwt/v5"

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
