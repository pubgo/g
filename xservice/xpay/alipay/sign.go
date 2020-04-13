package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func sign(m url.Values) string {
	//对url.values进行排序
	sign := ""
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if m.Get(k) != "" {
			if i == 0 {
				sign = k + "=" + m.Get(k)
			} else {
				sign = sign + "&" + k + "=" + m.Get(k)
			}
		}
	}

	//对排序后的数据进行rsa2加密，获得sign
	b, _ := rsaEncrypt([]byte(sign))
	return base64.StdEncoding.EncodeToString(b)
}

func rsaEncrypt(origData []byte) ([]byte, error) {
	//key := private_key
	//log.Println(key)
	block2, _ := pem.Decode([]byte(main.pemPrivateKey)) //PiravteKeyData为私钥文件的字节数组
	if block2 == nil {
		fmt.Println("block空")
		return nil, nil
	}
	//priv即私钥对象,block2.Bytes是私钥的字节流
	priv, err := x509.ParsePKCS1PrivateKey(block2.Bytes)
	if err != nil {
		fmt.Println("无法还原私钥")
		return nil, nil
	}
	p := priv
	h2 := sha256.New()
	h2.Write(origData)
	hashed := h2.Sum(nil)
	signature2, err := rsa.SignPKCS1v15(rand.Reader, p,
		crypto.SHA256, hashed) //签名
	return signature2, err
}

func VerifySign(data url.Values, key []byte) (ok bool, err error) {
	return verifySign(data, key)
}

func verifySign(data url.Values, key []byte) (ok bool, err error) {
	sign := data.Get("sign")
	signType := data.Get("sign_type")

	var keys = make([]string, 0, 0)
	for key, value := range data {
		if key == "sign" || key == "sign_type" {
			continue
		}
		if len(value) > 0 {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(data.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var s = strings.Join(pList, "&")

	return verifyData([]byte(s), signType, sign, key)
}

func verifyData(data []byte, signType, sign string, key []byte) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	if signType == "RSA" {
		err = VerifyPKCS1v15(data, signBytes, key, crypto.SHA1)
	} else {
		err = VerifyPKCS1v15(data, signBytes, key, crypto.SHA256)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func VerifyPKCS1v15(src, sig, key []byte, hash crypto.Hash) error {
	pub, err := ParsePKCS1PublicKey(key)
	if err != nil {
		return err
	}
	return VerifyPKCS1v15WithKey(src, sig, pub, hash)
}

func VerifyPKCS1v15WithKey(src, sig []byte, key *rsa.PublicKey, hash crypto.Hash) error {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)
	return rsa.VerifyPKCS1v15(key, hash, hashed, sig)
}

func ParsePKCS1PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("private key error")
	}

	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, err
}

func ParsePKCS1PublicKey(data []byte) (key *rsa.PublicKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key error")
	}

	return key, err
}

func NewRequest(method, url string, params url.Values) (*http.Request, error) {
	var m = strings.ToUpper(method)
	var body io.Reader
	if m == "GET" || m == "HEAD" {
		if len(params) > 0 {
			if strings.Contains(url, "?") {
				url = url + "&" + params.Encode()
			} else {
				url = url + "?" + params.Encode()
			}
		}
	} else {
		body = strings.NewReader(params.Encode())
	}
	return http.NewRequest(m, url, body)
}

//
//func NotifyVerify(partnerId, notifyId string) bool {
//	var param = url.Values{}
//	param.Add("service", "notify_verify")
//	param.Add("partner", partnerId)
//	param.Add("notify_id", notifyId)
//	req, err := NewRequest("GET", notify_domain, param)
//	if err != nil {
//		return false
//	}
//
//	rep, err := http.Client.Do(req)
//	if err != nil {
//		return false
//	}
//	defer rep.Body.Close()
//
//	data, err := ioutil.ReadAll(rep.Body)
//	if err != nil {
//		return false
//	}
//	if string(data) == "true" {
//		return true
//	}
//	return false
//}
