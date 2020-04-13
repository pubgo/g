package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/pubgo/g/json"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type (
	// ErrorHandleFunc error handling function
	ErrorHandleFunc func(*gin.Context, error)
	// Config defines the config for Session middleware
	Config struct {
		// error handling when starting the session
		ErrorHandleFunc ErrorHandleFunc
		// keys stored in the context
		TokenKey string
		// defines a function to skip middleware.Returning true skips processing
		// the middleware.
		Skipper func(*gin.Context) bool
	}
)

var (
	// DefaultConfig is the default middleware config.
	DefaultConfig = Config{
		ErrorHandleFunc: func(ctx *gin.Context, err error) {
			ctx.AbortWithError(500, err)
		},
		TokenKey: "github.com/go-oauth2/gin-server/access-token",
		Skipper: func(_ *gin.Context) bool {
			return false
		},
	}
)

// HandleTokenVerify Verify the access token of the middleware
func HandleTokenVerify(config ...Config) gin.HandlerFunc {
	cfg := DefaultConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.ErrorHandleFunc == nil {
		cfg.ErrorHandleFunc = DefaultConfig.ErrorHandleFunc
	}

	tokenKey := cfg.TokenKey
	if tokenKey == "" {
		tokenKey = DefaultConfig.TokenKey
	}

	return func(c *gin.Context) {
		if cfg.Skipper != nil && cfg.Skipper(c) {
			c.Next()
			return
		}
		ti, err := gServer.ValidationBearerToken(c.Request)
		if err != nil {
			cfg.ErrorHandleFunc(c, err)
			return
		}

		c.Set(tokenKey, ti)
		c.Next()
	}
}

// TokenContainer stores all relevant token information
type TokenContainer struct {
	Token     *oauth2.Token
	Scopes    map[string]interface{} // LDAP record vom Benutzer (cn, ..
	GrantType string                 // password, ??
	Realm     string                 // services, employees
}

// AccessCheckFunction is a function that checks if a given token grants
// access.
type AccessCheckFunction func(tc *TokenContainer, ctx *gin.Context) bool

func extractToken(r *http.Request) (*oauth2.Token, error) {
	hdr := r.Header.Get("Authorization")
	if hdr == "" {
		return nil, errors.New("No authorization header")
	}

	th := strings.Split(hdr, " ")
	if len(th) != 2 {
		return nil, errors.New("Incomplete authorization header")
	}

	return &oauth2.Token{AccessToken: th[1], TokenType: th[0]}, nil
}

func RequestAuthInfo(t *oauth2.Token) ([]byte, error) {
	var uv = make(url.Values)
	// uv.Set("realm", o.Realm)
	uv.Set("access_token", t.AccessToken)
	infoURL := AuthInfoURL + "?" + uv.Encode()
	client := &http.Client{Transport: &Transport}
	req, err := http.NewRequest("GET", infoURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func ParseTokenContainer(t *oauth2.Token, data map[string]interface{}) (*TokenContainer, error) {
	tdata := make(map[string]interface{})

	ttype := data["token_type"].(string)
	gtype := data["grant_type"].(string)

	realm := data["realm"].(string)
	exp := data["expires_in"].(float64)
	tok := data["access_token"].(string)
	if ttype != t.TokenType {
		return nil, errors.New("Token type mismatch")
	}
	if tok != t.AccessToken {
		return nil, errors.New("Mismatch between verify request and answer")
	}

	scopes := data["scope"].([]interface{})
	for _, scope := range scopes {
		sscope := scope.(string)
		sval, ok := data[sscope]
		if ok {
			tdata[sscope] = sval
		}
	}

	return &TokenContainer{
		Token: &oauth2.Token{
			AccessToken: tok,
			TokenType:   ttype,
			Expiry:      time.Now().Add(time.Duration(exp) * time.Second),
		},
		Scopes:    tdata,
		Realm:     realm,
		GrantType: gtype,
	}, nil
}

func GetTokenContainer(token *oauth2.Token) (*TokenContainer, error) {
	body, err := RequestAuthInfo(token)
	if err != nil {
		glog.Errorf("[Gin-OAuth] RequestAuthInfo failed caused by: %s", err)
		return nil, err
	}
	// extract AuthInfo
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		glog.Errorf("[Gin-OAuth] JSON.Unmarshal failed caused by: %s", err)
		return nil, err
	}
	if _, ok := data["error_description"]; ok {
		var s string
		s = data["error_description"].(string)
		glog.Errorf("[Gin-OAuth] RequestAuthInfo returned an error: %s", s)
		return nil, errors.New(s)
	}
	return ParseTokenContainer(token, data)
}

func getTokenContainer(ctx *gin.Context) (*TokenContainer, bool) {
	var oauthToken *oauth2.Token
	var tc *TokenContainer
	var err error

	if oauthToken, err = extractToken(ctx.Request); err != nil {
		glog.Errorf("[Gin-OAuth] Can not extract oauth2.Token, caused by: %s", err)
		return nil, false
	}
	if !oauthToken.Valid() {
		glog.Infof("[Gin-OAuth] Invalid Token - nil or expired")
		return nil, false
	}

	if tc, err = GetTokenContainer(oauthToken); err != nil {
		glog.Errorf("[Gin-OAuth] Can not extract TokenContainer, caused by: %s", err)
		return nil, false
	}

	return tc, true
}

//
// TokenContainer
//
// Validates that the AccessToken within TokenContainer is not expired and not empty.
func (t *TokenContainer) Valid() bool {
	if t.Token == nil {
		return false
	}
	return t.Token.Valid()
}

// Router middleware that can be used to get an authenticated and authorized service for the whole router group.
// Example:
//
//      var endpoints oauth2.Endpoint = oauth2.Endpoint{
//	        AuthURL:  "https://token.oauth2.corp.com/access_token",
//	        TokenURL: "https://oauth2.corp.com/corp/oauth2/tokeninfo",
//      }
//      var acl []ginoauth2.AccessTuple = []ginoauth2.AccessTuple{{"employee", 1070, "sszuecs"}, {"employee", 1114, "njuettner"}}
//      router := gin.Default()
//	private := router.Group("")
//	private.Use(ginoauth2.Auth(ginoauth2.UidCheck, ginoauth2.endpoints))
//	private.GET("/api/private", func(c *gin.Context) {
//		c.JSON(200, gin.H{"message": "Hello from private"})
//	})
//
func Auth(accessCheckFunction AccessCheckFunction, endpoints oauth2.Endpoint) gin.HandlerFunc {
	return AuthChain(endpoints, accessCheckFunction)
}

// AuthChain is a router middleware that can be used to get an authenticated
// and authorized service for the whole router group. Similar to Auth, but
// takes a chain of AccessCheckFunctions and only fails if all of them fails.
// Example:
//
//      var endpoints oauth2.Endpoint = oauth2.Endpoint{
//	        AuthURL:  "https://token.oauth2.corp.com/access_token",
//	        TokenURL: "https://oauth2.corp.com/corp/oauth2/tokeninfo",
//      }
//      var acl []ginoauth2.AccessTuple = []ginoauth2.AccessTuple{{"employee", 1070, "sszuecs"}, {"employee", 1114, "njuettner"}}
//      router := gin.Default()
//	    private := router.Group("")
//      checkChain := []AccessCheckFunction{
//          ginoauth2.UidCheck,
//          ginoauth2.GroupCheck,
//      }
//      private.Use(ginoauth2.AuthChain(checkChain, ginoauth2.endpoints))
//      private.GET("/api/private", func(c *gin.Context) {
//          c.JSON(200, gin.H{"message": "Hello from private"})
//      })
//
func AuthChain(endpoints oauth2.Endpoint, accessCheckFunctions ...AccessCheckFunction) gin.HandlerFunc {
	// init
	AuthInfoURL = endpoints.TokenURL
	// middleware
	return func(ctx *gin.Context) {
		t := time.Now()
		varianceControl := make(chan bool, 1)

		go func() {
			tokenContainer, ok := getTokenContainer(ctx)
			if !ok {
				// set LOCATION header to auth endpoint such that the user can easily get a new access-token
				ctx.Writer.Header().Set("Location", endpoints.AuthURL)
				ctx.AbortWithError(http.StatusUnauthorized, errors.New("No token in context"))
				varianceControl <- false
				return
			}

			if !tokenContainer.Valid() {
				// set LOCATION header to auth endpoint such that the user can easily get a new access-token
				ctx.Writer.Header().Set("Location", endpoints.AuthURL)
				ctx.AbortWithError(http.StatusUnauthorized, errors.New("Invalid Token"))
				varianceControl <- false
				return
			}

			for i, fn := range accessCheckFunctions {
				if fn(tokenContainer, ctx) {
					varianceControl <- true
					break
				}

				if len(accessCheckFunctions)-1 == i {
					ctx.AbortWithError(http.StatusForbidden, errors.New("Access to the Resource is forbidden"))
					varianceControl <- false
					return
				}
			}
		}()

		select {
		case ok := <-varianceControl:
			if !ok {
				glog.V(2).Infof("[Gin-OAuth] %12v %s access not allowed", time.Since(t), ctx.Request.URL.Path)
				return
			}
		case <-time.After(VarianceTimer):
			ctx.AbortWithError(http.StatusGatewayTimeout, errors.New("Authorization check overtime"))
			glog.V(2).Infof("[Gin-OAuth] %12v %s overtime", time.Since(t), ctx.Request.URL.Path)
			return
		}

		glog.V(2).Infof("[Gin-OAuth] %12v %s access allowed", time.Since(t), ctx.Request.URL.Path)
	}
}