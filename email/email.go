package email

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"regexp"
	"strings"
	"time"

	"github.com/pubgo/xerror"
)

const forceDisconnectAfter = time.Second * 5

var (
	ErrBadFormat        = errors.New("invalid format")
	ErrUnresolvableHost = errors.New("unresolvable host")

	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

var _PersonalEmail = map[string]bool{
	"163.com":        true,
	"126.com":        true,
	"yeah.net":       true,
	"foxmail.com":    true,
	"188.com":        true,
	"sina.com":       true,
	"sina.cn":        true,
	"vip.sina.com":   true,
	"21cn.com":       true,
	"sohu.com":       true,
	"qq.com":         true,
	"gmail.com":      true,
	"outlook.com":    true,
	"hotmail.com":    true,
	"189.cn":         true,
	"139.com":        true,
	"iCloud.com":     true,
	"mydomain.com":   true,
	"hushmail.com":   true,
	"eyou.com":       true,
	"ymail.cn":       true,
	"mail200.com.tw": true,
	"pie.com.tw":     true,
	"kiss99.com":     true,
	"home.com.tw":    true,
	"zj.com":         true,
	"aol.com":        true,
	"nextmail.ru":    true,
	"dezigner.ru":    true,
	"email.su":       true,
	"epage.ru":       true,
	"hu2.ru":         true,
	"mail2k.ru":      true,
	"nxt.ru":         true,
	"programist.ru":  true,
	"student.su":     true,
	"xaker.ru":       true,
	"wo.cn":          true,
	"aliyun.com":     true,
	"tianya.cn":      true,
	"hainan.net":     true,
	"mail.ru":        true,
	"inbox.ru":       true,
	"list.ru":        true,
	"bk.ru":          true,
	"goo.jp":         true,
	"tom.com":        true,
	"netease.com":    true,
	"2980.com":       true,
	"kuikoo.com":     true,
	"fastmail.fm":    true,
	"china.com":      true,
	"vip.tom.net":    true,
	"example.com":    true,
	"icloud.com":     true,
}

// GetName get name of email
func GetName(email string) string {
	slices := strings.SplitN(email, "@", 2)
	if len(slices) > 0 {
		return slices[0]
	}
	return email
}

// IsCompanyEmail checks if email is a company email
func IsCompanyEmail(email string) (bool, string, error) {
	status, domain := VerifyEmailFormat(email)

	if !status {
		return false, "", errors.New("wrong email format")
	}

	emailStatus := IsKnowHost(domain)

	return emailStatus, domain, nil

}

// VerifyEmailFormat
func VerifyEmailFormat(email string) (bool, string) {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	status := reg.MatchString(email)
	domains := strings.Split(email, "@")
	if len(domains) != 2 {
		return status, ""
	}

	return status, domains[1]
}

func LookMX(domain string) ([]string, error) {
	ns, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}

	var mx []string

	for _, n := range ns {

		mx = append(mx, n.Host)
	}

	return mx, nil
}

func IsKnowHost(host string) bool {
	if v, ok := _PersonalEmail[host]; ok {
		if v {
			return false
		}
	}
	return true
}

func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

func ValidateHost(email string) (err error) {
	defer xerror.RespErr(&err)

	_, host := split(email)
	mx, err := net.LookupMX(host)
	xerror.Panic(err, "unresolvable host")

	client := xerror.PanicErr(DialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), forceDisconnectAfter)).(*smtp.Client)
	defer client.Close()

	err = client.Hello("checkmail.me")
	if err != nil {
		return err
	}
	err = client.Mail("lansome-cowboy@gmail.com")
	if err != nil {
		return err
	}
	err = client.Rcpt(email)
	if err != nil {
		return err
	}
	return nil
}

// DialTimeout returns a new Client connected to an SMTP server at addr.
// The addr must include a port, as in "mail.example.com:smtp".
func DialTimeout(addr string, timeout time.Duration) (c *smtp.Client, err error) {
	defer xerror.RespErr(&err)

	conn := xerror.PanicErr(net.DialTimeout("tcp", addr, timeout)).(net.Conn)
	t := time.AfterFunc(timeout, func() { conn.Close() })
	defer t.Stop()

	host, _, err := net.SplitHostPort(addr)
	xerror.Panic(err)

	c, err = smtp.NewClient(conn, host)
	xerror.Panic(err, "smtp client connect failed")
	return
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}
