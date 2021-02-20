package sendcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xconfig/xconfig_email/abc"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pubgo/x/pkg"
	"github.com/pubgo/xerror"
)

var _email map[string]*email

// InitEmail init email
func InitEmail() (err error) {
	defer xerror.RespErr(&err)

	// 加载配置
	_cfg := xconfig.Default().Email.SendCloud
	xerror.PanicT(len(_cfg.Cfg) == 0, "email config count is 0")
	_email = make(map[string]*email, len(_cfg.Cfg))

	for _, cfg := range _cfg.Cfg {
		var opt = new(email)
		if cfg.URL == "" {
			cfg.URL = "http://api.sendcloud.net/apiv2/mail/send"
		}
		opt.url = cfg.URL
		opt.apiKey = cfg.APIKey
		opt.apiUser = cfg.APIUser
		opt.from = cfg.From
		opt.fromName = cfg.FromName
		_email[cfg.Name] = opt
	}
	_email[xconfig.DefaultName] = _email[_cfg.Default]
	return
}

// GetEmail get redis instance with name
func GetEmail(name ...string) (c abc.IMessage) {
	_name := xconfig.DefaultName
	if len(name) > 0 {
		_name = name[0]
	}
	c = _email[_name]
	xerror.PanicT(pkg.IsNone(c), "email instance %s is nil", _name)
	return
}

type email struct {
	url      string
	apiUser  string
	apiKey   string
	from     string
	fromName string
}

func (e *email) Send(tos []string, templateTitle, templateContent string, templateParams []string) (err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(len(tos) != len(templateParams), "params tos and templateParams length not equal")

	for i, to := range tos {
		params := url.Values{
			"apiUser":  {e.apiUser},
			"apiKey":   {e.apiKey},
			"from":     {e.from},
			"fromName": {e.fromName},
			"to":       {to},
			"subject":  {templateTitle},
			"html":     {fmt.Sprintf(templateContent, templateParams[i])},
		}

		req, err := http.NewRequest("POST", e.url, bytes.NewBufferString(params.Encode()))
		xerror.Panic(err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp := xerror.PanicErr(client.Do(req)).(*http.Response)
		bodyContent := xerror.PanicErr(ioutil.ReadAll(resp.Body)).([]byte)
		xerror.Panic(resp.Body.Close())

		var emailRes emailResult
		xerror.Panic(json.Unmarshal(bodyContent, &emailRes))
		xerror.PanicT(emailRes.StatusCode != 200, emailRes.Message)
	}
	return
}
