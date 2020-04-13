package auth

import (
	"github.com/go-session/session"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/server"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/linkedin"
)

// SetTokenType token type
func SetTokenType(tokenType string) {
	gServer.Config.TokenType = tokenType
}

// SetAllowGetAccessRequest to allow GET requests for the token
func SetAllowGetAccessRequest(allow bool) {
	gServer.Config.AllowGetAccessRequest = allow
}

// SetAllowedResponseType allow the authorization types
func SetAllowedResponseType(types ...oauth2.ResponseType) {
	gServer.Config.AllowedResponseTypes = types
}

// SetAllowedGrantType allow the grant types
func SetAllowedGrantType(types ...oauth2.GrantType) {
	gServer.Config.AllowedGrantTypes = types
}

// SetClientInfoHandler get client info from request
func SetClientInfoHandler(handler server.ClientInfoHandler) {
	gServer.ClientInfoHandler = handler
}

// SetClientAuthorizedHandler check the client allows to use this authorization grant type
func SetClientAuthorizedHandler(handler server.ClientAuthorizedHandler) {
	gServer.ClientAuthorizedHandler = handler
}

// SetClientScopeHandler check the client allows to use scope
func SetClientScopeHandler(handler server.ClientScopeHandler) {
	gServer.ClientScopeHandler = handler
}

// SetUserAuthorizationHandler get user id from request authorization
func SetUserAuthorizationHandler(handler server.UserAuthorizationHandler) {
	gServer.UserAuthorizationHandler = handler
}

// SetPasswordAuthorizationHandler get user id from username and password
func SetPasswordAuthorizationHandler(handler server.PasswordAuthorizationHandler) {
	gServer.PasswordAuthorizationHandler = handler
}

// SetRefreshingScopeHandler check the scope of the refreshing token
func SetRefreshingScopeHandler(handler server.RefreshingScopeHandler) {
	gServer.RefreshingScopeHandler = handler
}

// SetResponseErrorHandler response error handling
func SetResponseErrorHandler(handler server.ResponseErrorHandler) {
	gServer.ResponseErrorHandler = handler
}

// SetInternalErrorHandler internal error handling
func SetInternalErrorHandler(handler server.InternalErrorHandler) {
	gServer.InternalErrorHandler = handler
}

// SetExtensionFieldsHandler in response to the access token with the extension of the field
func SetExtensionFieldsHandler(handler server.ExtensionFieldsHandler) {
	gServer.ExtensionFieldsHandler = handler
}

// SetAccessTokenExpHandler set expiration date for the access token
func SetAccessTokenExpHandler(handler server.AccessTokenExpHandler) {
	gServer.AccessTokenExpHandler = handler
}

// SetAuthorizeScopeHandler set scope for the access token
func SetAuthorizeScopeHandler(handler server.AuthorizeScopeHandler) {
	gServer.AuthorizeScopeHandler = handler
}


// Google returns a new Google OAuth 2.0 backend endpoint.
func Google(conf *oauth2.Config) macaron.Handler {
	conf.Endpoint = google.Endpoint
	return NewOAuth2Provider(conf)
}

// Github returns a new Github OAuth 2.0 backend endpoint.
func Github(conf *oauth2.Config) macaron.Handler {
	conf.Endpoint = github.Endpoint
	return NewOAuth2Provider(conf)
}

// Facebook returns a new Facebook OAuth 2.0 backend endpoint.
func Facebook(conf *oauth2.Config) macaron.Handler {
	conf.Endpoint = facebook.Endpoint
	return NewOAuth2Provider(conf)
}

// LinkedIn returns a new LinkedIn OAuth 2.0 backend endpoint.
func LinkedIn(conf *oauth2.Config) macaron.Handler {
	conf.Endpoint = linkedin.Endpoint
	return NewOAuth2Provider(conf)
}

// NewOAuth2Provider returns a generic OAuth 2.0 backend endpoint.
func NewOAuth2Provider(conf *oauth2.Config) macaron.Handler {
	return func(s session.Store, ctx *macaron.Context) {
		if ctx.Req.Method == "GET" {
			switch ctx.Req.URL.Path {
			case PathLogin:
				login(conf, ctx, s)
			case PathLogout:
				logout(ctx, s)
			case PathCallback:
				handleOAuth2Callback(conf, ctx, s)
			}
		}
		tk := unmarshallToken(s)
		if tk != nil {
			// check if the access token is expired
			if tk.Expired() && tk.Refresh() == "" {
				s.Delete(KEY_TOKEN)
				tk = nil
			}
		}
		// Inject tokens.
		ctx.MapTo(tk, (*Tokens)(nil))
	}
}