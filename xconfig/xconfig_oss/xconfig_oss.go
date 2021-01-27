package xconfig_oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pubgo/x/pkg"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xdi"
	"github.com/pubgo/x/xenv"
	"github.com/pubgo/xerror"
	"regexp"
)

type Oss struct {
	_oss map[string]*oss.Client
}

// GetOss
// get oss instance with name
func (t *Oss) GetOss(name ...string) (c *oss.Client) {
	_name := xconfig.DefaultName
	if len(name) > 0 {
		_name = name[0]
	}

	c = t._oss[_name]
	xerror.PanicT(pkg.IsNone(c), "oss instance %s is nil", _name)
	return
}

// GetBucket
// get bucket
// name: bucket [oss name,bucket name]
func (t *Oss) GetBucket(name ...string) (c *oss.Bucket, err error) {
	defer xerror.RespErr(&err)

	_name := ""
	if len(name) > 0 {
		_name = name[len(name)-1]
	}
	c = xerror.PanicErr(t.GetOss(name[:len(name)-1]...).Bucket(_name)).(*oss.Bucket)
	return
}

func init() {
	xdi.InitProvide(func(cfg *xconfig.Config) *Oss {
		defer xerror.Assert()

		// 加载配置
		_cfg := cfg.Storage.Oss
		xerror.PanicT(len(_cfg.Cfg) == 0, "oss config count is 0")
		_oss := make(map[string]*oss.Client, len(_cfg.Cfg))

		for _, cfg := range _cfg.Cfg {
			var opt = new(oss.Config)
			if cfg.URL != "" {
				opt = _ParseOssUrl(cfg.URL)
			} else {
				opt.AccessKeyID = cfg.AccessKeyID
				opt.AccessKeySecret = cfg.AccessKeySecret
				opt.Endpoint = cfg.Endpoint
			}

			opt.IsDebug = xenv.IsDebug()
			if cfg.RetryTimes > 0 {
				opt.RetryTimes = uint(cfg.RetryTimes)
			}

			if cfg.Timeout > 0 {
				opt.Timeout = uint(cfg.Timeout)
			}

			var ocOpt []oss.ClientOption
			_c := xerror.PanicErr(oss.New(opt.Endpoint, opt.AccessKeyID, opt.AccessKeySecret, ocOpt...)).(*oss.Client)

			// check if oss is connected
			xerror.PanicErr(_c.ListBuckets())
			_oss[cfg.Name] = _c
		}
		_oss[xconfig.DefaultName] = _oss[_cfg.Default]
		return &Oss{_oss: _oss}
	})
}

func _ParseOssUrl(url string) *oss.Config {
	c := &oss.Config{}
	dt := regexp.MustCompile(`oss://(?P<username>.*):(?P<password>.*)@(?P<host>.*)`).FindStringSubmatch(url)
	if len(dt) == 0 {
		panic(fmt.Sprintf("url %s parse error", url))
	}
	c.AccessKeyID = dt[1]
	c.AccessKeySecret = dt[2]
	c.Endpoint = dt[3]
	return c
}

var exts = [...]string{
	"gif", "jpg", "jpeg", "bmp", "png", "ico", "psd",
	"mp3", "wma", "wav", "amr",
	"rm", "rmvb", "wmv", "avi", "mpg", "mpeg", "mp4", "mov", "flv", "swf", "mkv", "ogg", "ogv", "webm", "mid",
	"txt", "doc", "docx", "xls", "xlsx", "ppt", "pptx", "pps", "pdf", "chm", "md", "json", "sql",
	"rar", "zip", "7z", "tar", "gz", "bz2", "cab", "iso", "tar.gz", "mmap", "xmind", "md", "xml",
}

func IsExtLimit(ext string) bool {
	for _, s := range exts {
		if s == ext {
			return false
		}
	}
	return true
}
