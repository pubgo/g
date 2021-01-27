package xmlrpc

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/pubgo/x/xerror"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Array []interface{}
type Map map[string]interface{}

var xmlSpecial = map[byte]string{
	'<':  "&lt;",
	'>':  "&gt;",
	'"':  "&quot;",
	'\'': "&apos;",
	'&':  "&amp;",
}

func xmlEscape(s string) string {
	var b bytes.Buffer
	for i := 0; i < len(s); i++ {
		c := s[i]
		if s, ok := xmlSpecial[c]; ok {
			b.WriteString(s)
		} else {
			b.WriteByte(c)
		}
	}
	return b.String()
}

type valueNode struct {
	Type string `xml:"attr"`
	Body string `xml:"chardata"`
}

func next(p *xml.Decoder) (_ xml.Name, _ interface{}, err error) {
	defer xerror.RespErr(&err)

	se, e := nextStart(p)
	xerror.Panic(e)

	var nv interface{}
	switch se.Name.Local {
	case "string":
		var s string
		xerror.Panic(p.DecodeElement(&s, &se))
		return xml.Name{}, s, nil
	case "boolean":
		var s string
		xerror.Panic(p.DecodeElement(&s, &se))
		s = strings.TrimSpace(s)
		var b bool
		switch s {
		case "true", "1":
			b = true
		case "false", "0":
			b = false
		default:
			e = errors.New("invalid boolean value")
		}
		return xml.Name{}, b, e
	case "int", "i1", "i2", "i4", "i8":
		var s string
		xerror.Panic(p.DecodeElement(&s, &se))
		return xml.Name{}, xerror.PanicErr(strconv.Atoi(strings.TrimSpace(s))), e
	case "double":
		var s string
		xerror.Panic(p.DecodeElement(&s, &se))
		return xml.Name{}, xerror.PanicErr(strconv.ParseFloat(strings.TrimSpace(s), 64)).(float64), e
	case "dateTime.iso8601":
		var s string
		xerror.Panic(p.DecodeElement(&s, &se))
		t, e := time.Parse("20060102T15:04:05", s)
		if e != nil {
			t, e = time.Parse("2006-01-02T15:04:05-07:00", s)
			if e != nil {
				t, e = time.Parse("2006-01-02T15:04:05", s)
			}
		}
		return xml.Name{}, t, xerror.Wrap(e, "time parsing error")
	case "base64":
		var s string
		xerror.Panic(p.DecodeElement(&s, &se))
		return xml.Name{}, xerror.PanicErr(base64.StdEncoding.DecodeString(s)), nil
	case "member":
		xerror.PanicErr(nextStart(p))
		return next(p)
	case "value":
		xerror.PanicErr(nextStart(p))
		return next(p)
	case "name":
		xerror.PanicErr(nextStart(p))
		return next(p)
	case "struct":
		st := Map{}

		se, e = nextStart(p)
		for e == nil && se.Name.Local == "member" {
			// name
			se, e = nextStart(p)
			xerror.Panic(e)
			xerror.PanicT(se.Name.Local != "name", "invalid response")

			var name string
			xerror.Panic(p.DecodeElement(&name, &se))

			se, e = nextStart(p)
			xerror.Panic(e)

			// value
			_, value, e := next(p)
			xerror.Panic(e)
			xerror.PanicT(se.Name.Local != "value", "invalid response")

			st[name] = value

			se, e = nextStart(p)
			xerror.Panic(e)
		}
		return xml.Name{}, st, nil
	case "array":
		var ar Array
		xerror.PanicErr(nextStart(p))
		xerror.PanicErr(nextStart(p))
		// data
		// top of value
		for {
			_, value, e := next(p)
			if e != nil {
				break
			}
			ar = append(ar, value)

			if reflect.ValueOf(value).Kind() != reflect.Map {
				xerror.PanicErr(nextStart(p))
			}
		}
		return xml.Name{}, ar, nil
	case "nil":
		return xml.Name{}, nil, nil
	}

	if e = p.DecodeElement(nv, &se); e != nil {
		return xml.Name{}, nil, e
	}
	return se.Name, nv, e
}
func nextStart(p *xml.Decoder) (_ xml.StartElement, err error) {
	defer xerror.RespErr(&err)

	for {
		t, e := p.Token()
		if e != nil && e == io.EOF {
			break
		} else {
			xerror.Panic(e)
		}

		switch t := t.(type) {
		case xml.StartElement:
			return t, nil
		}
	}
	return
}

func toXml(v interface{}, typ bool) (s string) {
	if v == nil {
		return "<nil/>"
	}
	r := reflect.ValueOf(v)
	t := r.Type()
	k := t.Kind()

	if b, ok := v.([]byte); ok {
		return "<base64>" + base64.StdEncoding.EncodeToString(b) + "</base64>"
	}

	switch k {
	case reflect.Invalid:
		xerror.Panic("unsupported type")
	case reflect.Bool:
		var b string
		if v.(bool) {
			b = "1"
		} else {
			b = "0"
		}
		return fmt.Sprintf("<boolean>%s</boolean>", b)
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if typ {
			return fmt.Sprintf("<int>%v</int>", v)
		}
		return fmt.Sprintf("%v", v)
	case reflect.Uintptr:
		xerror.Panic("unsupported type")
	case reflect.Float32, reflect.Float64:
		if typ {
			return fmt.Sprintf("<double>%v</double>", v)
		}
		return fmt.Sprintf("%v", v)
	case reflect.Complex64, reflect.Complex128:
		xerror.Panic("unsupported type")
	case reflect.Array:
		s = "<array><data>"
		for n := 0; n < r.Len(); n++ {
			s += "<value>"
			s += toXml(r.Index(n).Interface(), typ)
			s += "</value>"
		}
		s += "</data></array>"
		return s
	case reflect.Chan:
		xerror.Panic("unsupported type")
	case reflect.Func:
		xerror.Panic("unsupported type")
	case reflect.Interface:
		return toXml(r.Elem(), typ)
	case reflect.Map:
		s = "<struct>"
		for _, key := range r.MapKeys() {
			s += "<member>"
			s += "<name>" + xmlEscape(key.Interface().(string)) + "</name>"
			s += "<value>" + toXml(r.MapIndex(key).Interface(), typ) + "</value>"
			s += "</member>"
		}
		s += "</struct>"
		return s
	case reflect.Ptr:
		xerror.Panic("unsupported type")
	case reflect.Slice:
		s = "<array><data>"
		for n := 0; n < r.Len(); n++ {
			s += "<value>"
			s += toXml(r.Index(n).Interface(), typ)
			s += "</value>"
		}
		s += "</data></array>"
		return s
	case reflect.String:
		if typ {
			return fmt.Sprintf("<string>%v</string>", xmlEscape(v.(string)))
		}
		return xmlEscape(v.(string))
	case reflect.Struct:
		s = "<struct>"
		for n := 0; n < r.NumField(); n++ {
			s += "<member>"
			s += "<name>"

			_s := t.Field(n).Tag.Get("xml")
			if _s == "" {
				_s = strings.ToLower(t.Field(n).Name[:1]) + t.Field(n).Name[1:]
			}
			s += _s
			s += "</name>"
			s += "<value>" + toXml(r.FieldByIndex([]int{n}).Interface(), true) + "</value>"
			s += "</member>"
		}
		s += "</struct>"
		return s
	case reflect.UnsafePointer:
		return toXml(r.Elem(), typ)
	}
	return
}

// Client is client of XMLRPC
type Client struct {
	HttpClient *http.Client
	url        string
}

// NewClient create new Client
func NewClient(url string) *Client {
	return &Client{
		HttpClient: &http.Client{Transport: http.DefaultTransport, Timeout: 10 * time.Second},
		url:        url,
	}
}

func makeRequest(name string, args ...interface{}) *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.WriteString(`<?xml version="1.0"?><methodCall>`)
	buf.WriteString("<methodName>" + xmlEscape(name) + "</methodName>")
	buf.WriteString("<params>")
	for _, arg := range args {
		buf.WriteString("<param><value>")
		buf.WriteString(toXml(arg, true))
		buf.WriteString("</value></param>")
	}
	buf.WriteString("</params></methodCall>")
	return buf
}

func call(client *http.Client, url, name string, args ...interface{}) (v interface{}, e error) {
	defer xerror.RespErr(&e)
	r := xerror.PanicErr(client.Post(url, "text/xml", makeRequest(name, args...))).(*http.Response)

	// Since we do not always read the entire body, discard the rest, which
	// allows the http transport to reuse the connection.
	defer io.Copy(ioutil.Discard, r.Body)
	defer r.Body.Close()

	xerror.PanicT(r.StatusCode/100 != 2, http.StatusText(http.StatusBadRequest))

	p := xml.NewDecoder(r.Body)
	se := xerror.PanicErr(nextStart(p)).(xml.StartElement) // methodResponse
	xerror.PanicT(se.Name.Local != "methodResponse", "invalid response: missing methodResponse")

	se = xerror.PanicErr(nextStart(p)).(xml.StartElement)
	xerror.PanicT(se.Name.Local != "params", "invalid response: missing params")

	se = xerror.PanicErr(nextStart(p)).(xml.StartElement)
	xerror.PanicT(se.Name.Local != "param", "invalid response: missing param")

	se = xerror.PanicErr(nextStart(p)).(xml.StartElement)
	xerror.PanicT(se.Name.Local != "value", "invalid response: missing value")

	_, v, e = next(p)
	return v, e
}

// Call call remote procedures function name with args
func (c *Client) Call(name string, args ...interface{}) (v interface{}, e error) {
	return call(c.HttpClient, c.url, name, args...)
}

// Global httpClient allows us to pool/reuse connections and not wastefully
// re-create transports for each request.
var httpClient = &http.Client{Transport: http.DefaultTransport, Timeout: 10 * time.Second}

// Call call remote procedures function name with args
func Call(url, name string, args ...interface{}) (v interface{}, e error) {
	return call(httpClient, url, name, args...)
}
