package sendcloud

import (
	"bytes"
	"encoding/json"
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/pkg/encoding/hashutil"
	"github.com/pubgo/g/xconfig"
	"github.com/pubgo/g/xconfig/xconfig_sms/abc"
	"github.com/pubgo/g/xerror"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

var _sms map[string]*sms

// InitSms init email
func InitSms() (err error) {
	defer xerror.RespErr(&err)

	// 加载配置
	_cfg := xconfig.Default().Sms.SendCloud
	xerror.PanicT(len(_cfg.Cfg) == 0, "sms config count is 0")
	_sms = make(map[string]*sms, len(_cfg.Cfg))

	for _, cfg := range _cfg.Cfg {
		var opt = new(sms)
		if cfg.URL == "" {
			cfg.URL = "http://www.sendcloud.net/smsapi/send"
		}
		opt.url = cfg.URL
		opt.apiKey = cfg.APIKey
		opt.apiUser = cfg.APIUser
		opt.from = cfg.From
		opt.fromName = cfg.FromName

		opt.msgType = cfg.MsgType
		opt.templateId = cfg.TemplateID

		_sms[cfg.Name] = opt
	}
	_sms[xconfig.DefaultName] = _sms[_cfg.Default]
	return
}

// GetRedis get redis instance with name
func GetSms(name ...string) (c abc.IMessage) {
	_name := xconfig.DefaultName
	if len(name) > 0 {
		_name = name[0]
	}
	c = _sms[_name]
	xerror.PanicT(pkg.IsNone(c), "email instance %s is nil", _name)
	return
}

type sms struct {
	url      string
	apiUser  string
	apiKey   string
	from     string
	fromName string

	msgType    string
	templateId string
}

func (e *sms) Send(tos []string, templateTitle, templateContent string, templateParams []string) (err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(len(tos) != len(templateParams), "params tos and templateParams length not equal")

	for i, to := range tos {
		postData := map[string]string{
			"msgType":    "0",
			"smsUser":    e.apiUser,
			"templateId": e.templateId,
			"phone":      to,
			"vars":       "{%code%:\"" + templateParams[i] + "\"}",
		}
		var keys []string
		for k := range postData {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var paramStr string
		for _, k := range keys {
			str := k + "=" + postData[k] + "&"
			paramStr += str
		}

		paramStr = e.apiKey + "&" + paramStr + e.apiKey
		signature := hashutil.MD5(paramStr)
		signature = strings.ToUpper(signature)
		postData["signature"] = signature

		postValues := url.Values{}
		for postKey, PostValue := range postData {
			postValues.Set(postKey, PostValue)
		}

		req, err := http.NewRequest("POST", e.url, bytes.NewBufferString(postValues.Encode()))
		xerror.Panic(err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp := xerror.PanicErr(client.Do(req)).(*http.Response)
		bodyContent := xerror.PanicErr(ioutil.ReadAll(resp.Body)).([]byte)
		xerror.Panic(resp.Body.Close())

		var smsRes smsResult
		xerror.Panic(json.Unmarshal(bodyContent, &smsRes))
		xerror.PanicT(smsRes.StatusCode != 200, smsRes.Message)
	}
	return
}
