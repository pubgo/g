package jwt

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	// InternalUserID is the auth.id used for internal requests
	InternalUserID string = "internal-sc-user"
	secretEnv             = "JWT_SECRET"
	AlgorithmHS256        = "HS256"
)

var (
	ErrInvalidSigningMethod        = errors.New("invalid signing method type")
	ErrTokenVerification           = errors.New("AUTH: JWT token could not be verified")
	TokenExpired            error  = errors.New("Token is expired")
	TokenNotValidYet        error  = errors.New("Token not active yet")
	TokenMalformed          error  = errors.New("That's not even a token")
	TokenInvalid            error  = errors.New("Couldn't handle this token:")
	SignKey                 string = "qmPlus"
	// ErrMissingSecretKey indicates Secret key is required
	ErrMissingSecretKey = errors.New("secret key is required")

	// ErrForbidden when HTTP status 403 is given
	ErrForbidden = errors.New("you don't have permission to access this resource")

	// ErrMissingAuthenticatorFunc indicates Authenticator is required
	ErrMissingAuthenticatorFunc = errors.New("ginJWTMiddleware.Authenticator func is undefined")

	// ErrMissingLoginValues indicates a user tried to authenticate without username or password
	ErrMissingLoginValues = errors.New("missing Username or Password")

	// ErrFailedAuthentication indicates authentication failed, could be faulty username or password
	ErrFailedAuthentication = errors.New("incorrect Username or Password")

	// ErrFailedTokenCreation indicates JWT Token failed to create, reason unknown
	ErrFailedTokenCreation = errors.New("failed to create JWT Token")

	// ErrExpiredToken indicates JWT token has expired. Can't refresh.
	ErrExpiredToken = errors.New("token is expired")

	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = errors.New("auth header is empty")

	// ErrMissingExpField missing exp field in token
	ErrMissingExpField = errors.New("missing exp field")

	// ErrWrongFormatOfExp field must be float64 format
	ErrWrongFormatOfExp = errors.New("exp must be float64 format")

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = errors.New("auth header is invalid")

	// ErrEmptyQueryToken can be thrown if authing with URL Query, the query token variable is empty
	ErrEmptyQueryToken = errors.New("query token is empty")

	// ErrEmptyCookieToken can be thrown if authing with a cookie, the token cokie is empty
	ErrEmptyCookieToken = errors.New("cookie token is empty")

	// ErrEmptyParamToken can be thrown if authing with parameter in path, the parameter in path is empty
	ErrEmptyParamToken = errors.New("parameter token is empty")

	// ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512, RS256, RS384 or RS512
	ErrInvalidSigningAlgorithm = errors.New("invalid signing algorithm")

	// ErrNoPrivKeyFile indicates that the given private key is unreadable
	ErrNoPrivKeyFile = errors.New("private key file unreadable")

	// ErrNoPubKeyFile indicates that the given public key is unreadable
	ErrNoPubKeyFile = errors.New("public key file unreadable")

	// ErrInvalidPrivKey indicates that the given private key is invalid
	ErrInvalidPrivKey = errors.New("private key invalid")

	// ErrInvalidPubKey indicates the the given public key is invalid
	ErrInvalidPubKey = errors.New("public key invalid")

	// IdentityKey default identity key
	IdentityKey  = "identity"
	rrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
)

//var _ IToken = (*tokenImpl)(nil)

type tokenImpl struct {
}

func (t *tokenImpl) Refresh(token string) (interface{}, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	_token, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return "secret", nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := _token.Claims.(*TokenClaims); ok && _token.Valid {
		jwt.TimeFunc = time.Now
		claims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return t.Create(claims, "")
	}
	return "", TokenExpired
}

// TokenClaims holds the JWT token claims
type TokenClaims = struct {
	Role string
	jwt.StandardClaims
}

// GetRole returns the role present in the token claims
func GetRole(c TokenClaims) (string, error) {
	//return "", errors.New("role is not present in the token claims")
	return c.Role, nil
}

func (t *tokenImpl) Create(claims *TokenClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	return token.SignedString([]byte(secret))
}

func (t *tokenImpl) Parse(token string, secret string) (TokenClaims, error) {
	// Parse and verify jwt access token
	_token, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse error")
		}
		return []byte("00000000"), nil

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return secret, nil
		} else {
			return nil, fmt.Errorf("expect token signed with HMAC but got %v", t.Header["alg"])
		}

		//if jwt.GetSigningMethod(mw.SigningAlgorithm) != token.Method {
		//	return nil, errors.New("Invalid signing algorithm")
		//}

		// Don't forget to validate the alg is what you expect:
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, ErrInvalidSigningMethod
		}
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return TokenClaims{}, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return TokenClaims{}, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return TokenClaims{}, TokenNotValidYet
			} else {
				return TokenClaims{}, TokenInvalid
			}
		}
	}

	claims, ok := _token.Claims.(TokenClaims)
	if !ok || !_token.Valid {
	}

	return claims, ErrTokenVerification
}

func (t *tokenImpl) IsInternal(token string) error {
	claims, err := t.Parse(token, "secret")
	if err != nil {
		return err
	}

	if claims.Id == InternalUserID {
		return nil
	}
	return errors.New("token has not been created internally")
}

func (t *tokenImpl) CreateClaims() TokenClaims {
	//req := TokenClaims{}
	//req["email"] = "kooksee@163.com"
	//req["password"] = "123456"
	//req["username"] = "kooksee"
	//req["group"] = "kooksee"
	//req["role"] = "master"
	//req["id"] = uuid.NewV1().String()
	//req["exp"] = time.Now().Add(24 * time.Hour * 7).Unix() // 过期时间 一周
	//req["uid"] = 0
	//req["sub"] = 0                               //Subject
	//req["iss"] = "gin-blog"                      //Issuer 签名的发行者
	//req["nbf"] = int64(time.Now().Unix() - 1000) // 签名生效时间
	//req["iat"] = time.Now().Unix()               //IssuedAt
	// Phone string `json:"phone"`
	/**
	UUID        uuid.UUID
	ID          uint
	NickName    string
	AuthorityId string
	*/
	return TokenClaims{}
}

func (t *tokenImpl) Validate(token string) (interface{}, error) {
	if len(token) == 0 {
		return nil, fmt.Errorf("token length is zero")
	}
	return t.Parse(token, "")
}
