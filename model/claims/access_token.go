package claims

import "github.com/golang-jwt/jwt"

type AccessTokenClaims struct {
	UserID    int       `json:"user_id"`
	TokenType TokenType `json:"token_type"`
	jwt.StandardClaims
}
