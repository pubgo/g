package xerror

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

var _ error = (*XErr)(nil)
var replace = strings.ReplaceAll

// NewXErr xerr
func NewXErr(name string) XErr {
	name = replace(replace(strings.Title(name), " ", ""), "-", "")
	return XErr{name: name}
}

// XErr struct
type XErr struct {
	name string
}

func (t XErr) Error() string {
	return t.name
}

// Err err
func (t XErr) Err(format string, args ...interface{}) error {
	return fmt.Errorf(t.name+": "+format, args...)
}

// Msg xerr
func (t XErr) Msg(format string, args ...interface{}) string {
	return fmt.Sprintf(t.name+": "+format, args...)
}

// Wrap xerr
func (t XErr) Wrap(err error) error {
	if err == nil {
		return nil
	}
	return errors.New(t.name + ": " + err.Error())
}

// New xerr
func (t XErr) New(format string) XErr {
	t.name = t.name + ":" + replace(replace(strings.Title(format), " ", ""), "-", "")
	return t
}

type _Err1 struct {
	M      map[string]interface{} `json:"m,omitempty"`
	Err    string                 `json:"err,omitempty"`
	Caller []string               `json:"caller,omitempty"`
	Sub    *_Err1                 `json:"sub,omitempty"`
}

func (t *_Err1) String() string {
	if t == nil {
		return ""
	}

	_dt, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	return string(_dt)

}

var _ IErr = (*_Err)(nil)

// _Err errors
type _Err struct {
	err    error
	sub    IErr
	caller []string
	m      map[string]interface{}
}

func (t *_Err) P() {
	fmt.Println(t.Error())
}

func (t *_Err) next() IErr {
	if t.sub == nil {
		return nil
	}

	return t.sub
}

func (t *_Err) node(e IErr) {
	if t.sub == nil {
		t.sub = e
		return
	}
	t.sub.node(e)
}

func (t *_Err) SetErr(err interface{}, args ...interface{}) {
	switch _err := err.(type) {
	case XErr:
		t.err = _err
	case error:
		t.err = _err
	case string:
		t.err = fmt.Errorf(_err, args...)
	default:
		log.Fatal(ErrUnknownType.Err("type %#v", _err).Error())
	}
}

// Err errors
func (t *_Err) Err() error {
	return t.err
}

// StackTrace
// errors
func (t *_Err) StackTrace() *_Err1 {
	if t.isNil() {
		return nil
	}

	defer t.reset()

	return t._err()
}

// Error
func (t *_Err) Error() string {
	if t.isNil() {
		return ""
	}

	var buf = &strings.Builder{}
	defer buf.Reset()

	var _errs = t._err1()
	_filterFile := []string{
		"env_test.go",
		"testing/testing.go",
		"src/runtime",
		"testing/testing.go",
		"src/reflect",
	}

	_filter := func(k string) bool {
		for _, _k := range _filterFile {
			if strings.Contains(k, _k) {
				return true
			}

			if skipErrorFile {
				if strings.Contains(k, "github.com/pubgo/g") && strings.Contains(k, "/xerror") {
					return true
				}
			}
		}
		return false
	}

	buf.WriteString("========================================================================================================================\n")
	for i := len(_errs) - 1; i > -1; i-- {
		if len(_errs[i].Caller) < 1 {
			continue
		}

		buf.WriteString(fmt.Sprintf("[%s]: %s %s\n", colorize("Debug", colorYellow), time.Now().Format("2006/01/02 - 15:04:05"), _errs[i].Caller[0]))
		//if _errs[i].Msg != "" {
		//	buf.WriteString(fmt.Sprintf("[ %s ]: %s\n", colorize("Msg", colorGreen), _errs[i].Msg))
		//}

		if _errs[i].Err != "" {
			buf.WriteString(fmt.Sprintf("[ %s ]: %s\n", colorize("Err", colorRed), _errs[i].Err))
		}

		if _errs[i].M != nil || len(_errs[i].M) != 0 {
			buf.WriteString(fmt.Sprintf("[  %s  ]: %#v\n", colorize("M", colorMagenta), _errs[i].M))
		}

		for _, k := range _errs[i].Caller[1:] {
			if _filter(k) {
				continue
			}

			buf.WriteString(time.Now().Format("[Debug] 2006/01/02 - 15:04:05 "))
			buf.WriteString(fmt.Sprintln(k))
		}
	}
	buf.WriteString("========================================================================================================================")
	return buf.String()
}

// Caller errors
func (t *_Err) Caller(caller ...string) {
	if !t.isNil() && len(caller) > 0 {
		t.caller = append(t.caller, caller...)
	}
}

func (t *_Err) Is(err error) bool {
	if t.err == nil {
		return false
	}

	return t.err == err
}

func (t *_Err) As(err error) bool {
	if t.err == nil {
		return false
	}

	return strings.HasPrefix(t.err.Error(), err.Error())
}

// M errors
func (t *_Err) M(k string, v interface{}) {
	if !t.isNil() {
		if t.m == nil {
			t.m = make(map[string]interface{})
		}
		t.m[k] = v
	}
}

func (t *_Err) isNil() bool {
	return t == nil || _isNone(t)
}

func (t *_Err) reset() {
	t.m = nil
	t.err = nil
	t.caller = t.caller[:0]
	t.sub = nil
}

func (t *_Err) _err1() []*_Err1 {
	if t.sub == nil {
		return []*_Err1{{
			M:      t.m,
			Err:    _err(t.err),
			Caller: t.caller,
		}}
	}

	return append(t.sub._err1(), &_Err1{
		M:      t.m,
		Err:    _err(t.err),
		Caller: t.caller,
		Sub:    t.sub._err(),
	})
}

func _err(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}

func (t *_Err) _err() *_Err1 {
	if t.sub == nil {
		return &_Err1{
			M:      t.m,
			Err:    _err(t.err),
			Caller: t.caller,
		}
	}

	return &_Err1{
		M:      t.m,
		Err:    _err(t.err),
		Caller: t.caller,
		Sub:    t.sub._err(),
	}
}
