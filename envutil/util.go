package envutil

import (
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var replace = strings.ReplaceAll

const _quote = "\\\n\r\"!$`"

func doubleQuoteEscape(line string) string {
	var n string
	for _, c := range _quote {
		if c == '\n' {
			n = `\n`
		}
		if c == '\r' {
			n = `\r`
		}
		line = replace(line, string(c), n)
	}
	return line
}

func _EnvKey(key string) string {
	prefix := Cfg.Prefix
	if prefix == "" {
		return key
	}

	key = upper(trim(key))
	kl := len(prefix) - 1
	if strings.HasPrefix(key, prefix[:kl]) {
		if key[kl] != DefaultPrefixSeparator {
			return strings.Replace(key, prefix[:kl], prefix, 1)
		}
		return key
	}

	return prefix + key
}

func _EnvValue(value string) string {
	return doubleQuoteEscape(trim(value))
}

func _HasPrefix(key string) bool {
	return strings.HasPrefix(key, Cfg.Prefix)
}

// copyFromSystem 加载并处理系统环境变量前缀
func copyFromSystem(prefix string) {
	if prefix == "" {
		return
	}

	kl := len(prefix) - 1
	for _, env := range os.Environ() {
		if _envs := strings.SplitN(env, "=", 2); len(_envs) == 2 && _envs[0] != "" {
			key := upper(trim(_envs[0]))
			if strings.HasPrefix(key, prefix[:kl]) && key[kl] != DefaultPrefixSeparator {
				fatal(os.Unsetenv(_envs[0]))
				fatal(os.Setenv(strings.Replace(key, prefix[:kl], prefix, 1), _envs[1]))
			}
		}
	}
}

// LoadFile 加载.env文件并添加前缀
func LoadFile(envFiles ...string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case error:
				err = r.(error)
			default:
				logger.Fatalln(r)
			}
		}
	}()

	if len(envFiles) == 0 {
		envFiles = []string{".env"}
	}

	for _, filename := range envFiles {
		_envPath, err := filepath.EvalSymlinks(filename)
		fatal(err)

		f, err := ioutil.ReadFile(_envPath)
		fatal(err)

		for _, env := range strings.Split(string(f), "\n") {
			env = trim(env)
			if env == "" {
				continue
			}

			if _envs := strings.SplitN(env, "=", 2); len(_envs) == 2 {
				if _envs[0] == "" {
					fatal(fmt.Errorf("env key is null, env: %s", env))
				}

				fatal(os.Unsetenv(_envs[0]))
				fatal(SetEnv(_envs[0], ExpandEnv(_envs[1])))

				continue
			}
			fatal(fmt.Errorf("env format not match error: %s", env))
		}
	}

	return
}

// ParseToMap 解析envs map[string]string
func ParseToMap(filterPrefix ...bool) map[string]string {
	_envs := os.Environ()

	envs := make(map[string]string, len(_envs))
	for _, env := range _envs {
		env = trim(env)

		if len(filterPrefix) > 0 && !_HasPrefix(env) {
			continue
		}

		if ev := strings.SplitAfterN(env, "=", 2); len(ev) == 2 && ev[0] != "" {
			envs[ev[0]] = ExpandEnv(ev[1])
		}
	}
	return envs
}

// Log env log
func Log(filterPrefix ...bool) {
	for k, v := range ParseToMap(filterPrefix...) {
		fmt.Println("env:", k, v)
	}
}

var _envRegexp = regexp.MustCompile(`\${(.+)}`)
var _safeEnvRegexp = regexp.MustCompile(`!{(.+)}`)

// ExpandEnv returns value of convert with environment variable.
// Return environment variable if value start with "${" and end with "}".
// Return default value if environment variable is empty or not exist.
//
// It accept value formats "${env}" , "${env||}}" , "${env||defaultValue}" , "defaultvalue".
// Examples:
//	v1 := config.ExpandValueEnv("${GOPATH}")			// return the GOPATH environment variable.
//	v2 := config.ExpandValueEnv("${GOAsta||/usr/local/go}")	// return the default value "/usr/local/go/".
//	v3 := config.ExpandValueEnv("Astaxie")				// return the value "Astaxie".
func ExpandEnv(value string) string {
	value = trim(value)

	// 匹配环境变量格式
	if _envRegexp.MatchString(value) {
		_vs := strings.Split(_envRegexp.FindStringSubmatch(value)[1], "||")
		_v := os.Getenv(upper(_vs[0]))
		if len(_vs) == 2 && _v == "" {
			_v = trim(_vs[1])
		}
		return _v
	}

	// 匹配加密数据格式
	if _safeEnvRegexp.MatchString(value) {
		_v := _safeEnvRegexp.FindStringSubmatch(value)[1]
		return string(myXorDecrypt(_v, []byte(GetEnv(DefaultSecretKey...))))
	}

	return value
}

// myXorEncrypt encrypt
func myXorEncrypt(text, key []byte) string {
	var _lk = len(key)
	for i := 0; i < len(text); i++ {
		text[i] ^= key[i*i*i%_lk]
	}
	return base32.StdEncoding.EncodeToString(text)
}

//myXorDecrypt decrypt
func myXorDecrypt(text string, key []byte) []byte {
	var _lk = len(key)
	_text, err := base32.StdEncoding.DecodeString(text)
	if err != nil {
		logger.Fatal(err)
	}

	for i := 0; i < len(_text); i++ {
		_text[i] ^= key[i*i*i%_lk]
	}
	return _text
}

// IsTrue true
func IsTrue(data string) bool {
	switch upper(data) {
	case "TRUE", "T", "1", "OK", "GOOD", "REAL", "ACTIVE", "ENABLED":
		return true
	default:
		return false
	}
}

// ExpandEnvMap Expand env map
func ExpandEnvMap(m map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		switch value := v.(type) {
		case string:
			m[k] = ExpandEnv(value)
		case map[string]interface{}:
			m[k] = ExpandEnvMap(value)
		case map[string]string:
			for k2, v2 := range value {
				value[k2] = ExpandEnv(v2)
			}
			m[k] = value
		}
	}
	return m
}
