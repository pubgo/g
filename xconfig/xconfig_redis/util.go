package xconfig_redis

import (
	"crypto/tls"
	"errors"
	"github.com/go-redis/redis/v7"
	"github.com/pubgo/x/xerror"
	"math/big"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func _ParseRedisUrl(uri string) (o *redis.Options, err error) {
	defer xerror.RespErr(&err)

	o = &redis.Options{Network: "tcp"}
	u := xerror.PanicErr(url.Parse(uri)).(*url.URL)
	xerror.PanicT(u.Scheme != "redis" && u.Scheme != "rediss", "invalid redis URL scheme: %s", u.Scheme)
	xerror.PanicT(len(u.Query()) > 0, "no options supported")

	if u.User != nil {
		if p, ok := u.User.Password(); ok {
			o.Password = p
		}
	}

	h, p, err := net.SplitHostPort(u.Host)
	if err != nil {
		h = u.Host
	}
	if h == "" {
		h = "localhost"
	}
	if p == "" {
		p = "6379"
	}
	o.Addr = net.JoinHostPort(h, p)

	f := strings.FieldsFunc(u.Path, func(r rune) bool {
		return r == '/'
	})

	switch len(f) {
	case 0:
		o.DB = 0
	case 1:
		if o.DB, err = strconv.Atoi(f[0]); err != nil {
			xerror.PanicM(err, "invalid redis database number: %q", f[0])
		}
	default:
		xerror.Panic(errors.New("invalid redis URL path: " + u.Path))
	}

	if u.Scheme == "rediss" {
		o.TLSConfig = &tls.Config{ServerName: h}
	}

	dt := regexp.MustCompile(`redis://:(?P<password>.*)@(?P<host>.*)/(?P<db>.*)`).FindStringSubmatch(uri)
	xerror.PanicT(len(dt) == 0, "url %s parse error", uri)
	if o.Password == "" {
		o.Password = dt[1]
	}

	db, _ := big.NewInt(0).SetString(dt[3], 10)
	o.DB = int(db.Int64())
	return
}
