package session_redis_store

import (
	"encoding/base32"
	sssions "github.com/gin-contrib/sessions"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/pubgo/g/xerror"
	"net/http"
	"strings"
)

//func New(name ...string) IStore {
//
//	xerror.Panic(xdi.Invoke(func(cfg xconfig.Config) {
//	}))
//
//	xerror.PanicT(_name == "", "name is empty")
//
//	for _, cfg := range _cfg.Cfg {
//		if _name != cfg.Name {
//			continue
//		}
//
//		_session := cfg.Session
//
//		var keys [][]byte
//		for _, key := range _session.KeyPairs {
//			keys = append(keys, []byte(key))
//		}
//
//		xerror.PanicT(_session.DriverName == "", "redis driver name is empty")
//
//		return &Store{
//			client: xconfig_redis.GetRedis(_session.DriverName),
//			Codecs: securecookie.CodecsFromPairs(keys...),
//			options: &sessions.Options{
//				SameSite: http.SameSite(_session.SameSite),
//				HttpOnly: _session.HTTPOnly,
//				Secure:   _session.Secure,
//				Domain:   _session.Domain,
//				Path:     _session.Path,
//				MaxAge:   _session.MaxAge,
//			},
//			DefaultMaxAge: 60 * 20,
//			maxLength:     4096,
//			keyPrefix:     cfg.Session.KeyPrefix,
//			serializer:    &GobSerializer{},
//		}
//	}
//
//	xerror.Panic(fmt.Errorf("initialize error"))
//
//	return nil
//}

// Store stores sessions in a redis backend.
type Store struct {
	client        _IRedis
	Codecs        []securecookie.Codec
	options       *sessions.Options // default configuration
	DefaultMaxAge int               // default Redis TTL for a MaxAge == 0 session
	maxLength     int
	keyPrefix     string
	serializer    SessionSerializer
}

// Close closes the underlying *redis.Pool
func (s *Store) Close() error {
	return s.client.Close()
}

// Get returns a session for the given name after adding it to the registry.
// See gorilla/sessions FilesystemStore.Get().
func (s *Store) Get(r *http.Request, name string) (*sessions.Session, error) {
	return sessions.GetRegistry(r).Get(s, name)
}

// New returns a session for the given name without adding it to the registry.
// See gorilla/sessions FilesystemStore.New().
func (s *Store) New(r *http.Request, name string) (sess *sessions.Session, err error) {
	defer xerror.RespErr(&err)

	sess = sessions.NewSession(s, name)
	options := *s.options
	sess.Options = &options
	sess.IsNew = true
	if c, errCookie := r.Cookie(name); errCookie == nil {
		xerror.Panic(securecookie.DecodeMulti(name, c.Value, &sess.ID, s.Codecs...))
		xerror.Panic(s.load(sess))
		sess.IsNew = false
	}
	return
}

// Save adds a single session to the response.
func (s *Store) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) (err error) {
	defer xerror.RespErr(&err)

	// Marked for deletion.
	if session.Options.MaxAge <= 0 {
		xerror.Panic(s.delete(session))
		http.SetCookie(w, sessions.NewCookie(session.Name(), "", session.Options))
		return
	}

	// Build an alphanumeric key for the redis store.
	if session.ID == "" {
		session.ID = strings.TrimRight(base32.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32)), "=")
	}

	xerror.Panic(s.save(session))
	http.SetCookie(w,
		sessions.NewCookie(session.Name(),
			xerror.PanicErr(securecookie.EncodeMulti(session.Name(),
				session.ID, s.Codecs...)).(string), session.Options))
	return
}

// Delete removes the session from redis, and sets the cookie to expire.
//
// WARNING: This method should be considered deprecated since it is not exposed via the gorilla/sessions interface.
// Set session.Options.MaxAge = -1 and call Save instead. - July 18th, 2013
func (s *Store) Delete(r *http.Request, w http.ResponseWriter, session *sessions.Session) (err error) {
	defer xerror.RespErr(&err)

	xerror.Panic(s.client.Del(s.keyPrefix + session.ID).Err())

	// Set cookie to expire.
	options := *session.Options
	options.MaxAge = -1
	http.SetCookie(w, sessions.NewCookie(session.Name(), "", &options))

	// Clear session values.
	for k := range session.Values {
		delete(session.Values, k)
	}

	return
}

// save stores the session in redis.
func (s *Store) save(session *sessions.Session) (err error) {
	defer xerror.RespErr(&err)

	b := xerror.PanicErr(s.serializer.Serialize(session)).([]byte)
	xerror.PanicT(s.maxLength != 0 && len(b) > s.maxLength, "SessionStore: the value to store is too big")
	xerror.Panic(s.client.SetNX(s.keyPrefix+session.ID, b, sessionExpire).Err())
	return
}

// load reads the session from redis.
// returns true if there is a sessoin data in DB
func (s *Store) load(sess *sessions.Session) error {
	_str := s.client.Get(s.keyPrefix + sess.ID)
	xerror.Panic(_str.Err())
	return s.serializer.Deserialize([]byte(_str.String()), sess)
}

// delete removes keys from redis if MaxAge<0
func (s *Store) delete(session *sessions.Session) error {
	return s.client.Del(s.keyPrefix + session.ID).Err()
}

func (s *Store) Options(options sssions.Options) {
	s.options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
