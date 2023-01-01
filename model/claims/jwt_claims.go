package claims

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  int    `json:"role"`
}
