package claims

import "github.com/golang-jwt/jwt"

type RefreshTokenClaims struct {
	UserID    int       `json:"user_id"`
	CusKey    string    `json:"cus_key"`
	TokenType TokenType `json:"token_type"`
	jwt.StandardClaims
}
