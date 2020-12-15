package qrcode

import (
	"github.com/pubgo/x/pkg/encoding/hashutil"
	"github.com/pubgo/x/pkg/fileutil"
	"github.com/pubgo/x/xerror"
	"image/jpeg"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

// NewQrCode initialize instance
func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) QrCode {
	return QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

// GetQrCodeFileName get qr file name
func GetQrCodeFileName(value string) string {
	return hashutil.MD5(value)
}

// GetQrCodeExt get qr file ext
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

// Encode generate QR code
func (q *QrCode) Encode(path string) (name string, err error) {
	defer xerror.RespErr(&err)

	name = GetQrCodeFileName(q.URL) + q.GetQrCodeExt()

	if fileutil.CheckNotExist(path+name) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		xerror.Panic(err)

		code, err = barcode.Scale(code, q.Width, q.Height)
		xerror.Panic(err)

		f, err := fileutil.MustOpen(name, path)
		xerror.Panic(err)
		defer f.Close()

		xerror.Panic(jpeg.Encode(f, code, nil))
	}

	return
}
