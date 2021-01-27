package session_redis_store

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/pubgo/x/xerror"
)

// SessionSerializer provides an interface hook for alternative serializers
type SessionSerializer interface {
	Deserialize(d []byte, ss *sessions.Session) error
	Serialize(ss *sessions.Session) ([]byte, error)
}

// JSONSerializer encode the session map to JSON.
type JSONSerializer struct{}

// Serialize to JSON. Will err if there are unmarshalable key values
func (s JSONSerializer) Serialize(ss *sessions.Session) (dt []byte, err error) {
	defer xerror.RespErr(&err)

	m := make(map[string]interface{}, len(ss.Values))
	for k, v := range ss.Values {
		ks, ok := k.(string)
		xerror.PanicT(!ok, "Non-string key value, cannot serialize session to JSON: %v", k)
		m[ks] = v
	}
	return json.Marshal(m)
}

// Deserialize back to map[string]interface{}
func (s JSONSerializer) Deserialize(d []byte, ss *sessions.Session) (err error) {
	defer xerror.RespErr(&err)

	m := make(map[string]interface{})
	xerror.PanicM(json.Unmarshal(d, &m), "redistore.JSONSerializer.deserialize()")

	for k, v := range m {
		ss.Values[k] = v
	}
	return
}

// GobSerializer uses gob package to encode the session map
type GobSerializer struct{}

// Serialize using gob
func (s GobSerializer) Serialize(ss *sessions.Session) (dt []byte, err error) {
	defer xerror.RespErr(&err)

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	xerror.PanicM(enc.Encode(ss.Values), "GobSerializer Encode Error")
	dt = buf.Bytes()

	return
}

// Deserialize back to map[interface{}]interface{}
func (s GobSerializer) Deserialize(d []byte, ss *sessions.Session) error {
	return gob.NewDecoder(bytes.NewBuffer(d)).Decode(&ss.Values)
}
