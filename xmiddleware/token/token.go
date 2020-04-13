package token

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"strings"
)

//Token 简单token验证
func Token(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RequestURI, "/docs") {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			auths := strings.SplitN(auth, " ", 2)
			if len(auths) != 2 {
				return
			}
			authMethod := auths[0]
			authB64 := auths[1]
			switch authMethod {
			case "Basic":
				authstr, err := base64.StdEncoding.DecodeString(authB64)
				if err != nil {
					io.WriteString(w, "Unauthorized!\n")
					return
				}
				userPwd := strings.SplitN(string(authstr), ":", 2)
				if len(userPwd) != 2 {
					io.WriteString(w, "Unauthorized!\n")
					return
				}
				username := userPwd[0]
				password := userPwd[1]
				if username == "goodrain" && password == "goodrain-api-test" {
					next.ServeHTTP(w, r)
					return
				}
			default:
				io.WriteString(w, "Unauthorized!\n")
				return
			}
			w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := os.Getenv("TOKEN")
		t := r.Header.Get("Authorization")
		if tt := strings.Split(t, " "); len(tt) == 2 {
			if tt[1] == token {
				next.ServeHTTP(w, r)
				return
			}
		}
		util.CloseRequest(r)
		w.WriteHeader(http.StatusUnauthorized)
	}
	return http.HandlerFunc(fn)
}

//FullToken token api校验
func FullToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RequestURI, "/docs") {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			auths := strings.SplitN(auth, " ", 2)
			if len(auths) != 2 {
				return
			}
			authMethod := auths[0]
			authB64 := auths[1]
			switch authMethod {
			case "Basic":
				authstr, err := base64.StdEncoding.DecodeString(authB64)
				if err != nil {
					io.WriteString(w, "Unauthorized!\n")
					return
				}
				userPwd := strings.SplitN(string(authstr), ":", 2)
				if len(userPwd) != 2 {
					io.WriteString(w, "Unauthorized!\n")
					return
				}
				username := userPwd[0]
				password := userPwd[1]
				if username == "goodrain" && password == "goodrain-api-test" {
					next.ServeHTTP(w, r)
					return
				}
			default:
				io.WriteString(w, "Unauthorized!\n")
				return
			}
			w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//logrus.Debugf("request uri is %s", r.RequestURI)
		t := r.Header.Get("Authorization")
		if tt := strings.Split(t, " "); len(tt) == 2 {
			if handler.GetTokenIdenHandler().CheckToken(tt[1], r.RequestURI) {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
	return http.HandlerFunc(fn)
}

func CheckToken() bool {

}
