package claims

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken           = "refresh"
)
