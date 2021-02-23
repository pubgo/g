package regexutil

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 验证用户名合法性，用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头。
func IsValidateUsername(username string) error {
	if len(username) == 0 {
		return errors.New("请输入用户名")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

// 验证是否是合法的邮箱
func IsValidateEmail(email string) (err error) {
	if len(email) == 0 {
		err = errors.New("邮箱格式不符合规范")
		return
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		err = errors.New("邮箱格式不符合规范")
	}
	return
}

//检测参数
func CheckParam(params ...string) bool {
	for _, value := range params {
		if len(value) == 0 {
			return false
		}
	}
	return true
}

//判断是否是手机号
func IsPhone(mobileNum string) bool {
	tmp := `^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`
	reg := regexp.MustCompile(tmp)
	return reg.MatchString(mobileNum)
}

//判断用户是否是邮件用户
func IsMail(username string) (isMail bool) {
	isMail = false
	if strings.Contains(username, "@") {
		isMail = true //是邮箱
	}
	return
}

//判断是否在测试环境下使用
func IsRunTesting() bool {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
		return strings.HasPrefix(os.Args[1], "-test")
	}
	return false
}

//判断是否是18或15位身份证
func IsIdCard(cardNo string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	if m, _ := regexp.MatchString(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`, cardNo); !m {
		return false
	}
	return true
}

// 0 is considered as a non valid port
func IsValidPort(port int) bool {
	return port > 0 && port < 65535
}
