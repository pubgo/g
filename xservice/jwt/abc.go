package jwt

type IToken interface {
	Create(claims TokenClaims, secret string) (string, error)
	Parse(token string, secret string) (TokenClaims, error)
	IsInternal(token string) error
	CreateClaims() TokenClaims
	Validate(token string) (interface{}, error)
	Refresh(token string) (interface{}, error)
}

