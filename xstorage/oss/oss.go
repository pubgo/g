package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xconfig/xconfig_oss"
	"github.com/pubgo/xerror"
	"io"
	"strings"
)

// New initialize
// name, bucket
func New(name ...string) Oss {
	_cfg := xconfig.Default().Storage.Oss
	_name := _cfg.Default
	_endpoint := _cfg.Endpoint

	if len(name) > 2 {
		_name = name[0]
	}

	for _, cfg := range _cfg.Cfg {
		if cfg.Name == _name {
			_endpoint = cfg.Endpoint
			break
		}
	}

	xerror.PanicT(_endpoint == "", "name or endpoint is empty")
	return Oss{
		name:     _name,
		bucket:   name[len(name)-1],
		endpoint: _endpoint,
		client:   xerror.PanicErr(xconfig_oss.GetBucket(name...)).(*oss.Bucket),
	}
}

type Oss struct {
	name     string
	bucket   string
	endpoint string
	client   *oss.Bucket
}

func (s Oss) PutObject(filename string, r io.Reader) error {
	return s.client.PutObject(filename, r)
}

func (s Oss) GetStoragePath(files ...string) string {
	return "https://" + s.bucket + "." + s.endpoint + "/" + strings.Join(files, "/")
}

func (s Oss) GetObject(filename string) (io.ReadCloser, error) {
	return s.client.GetObject(filename)
}
