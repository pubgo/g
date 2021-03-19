package netutil

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/pubgo/x/xutil"
	"github.com/pubgo/xerror"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/axgle/mahonia"
)

// LocalIP gets the first NIC's IP address.
func LocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if nil != err {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("can't get local IP")
}

// LocalMac gets the first NIC's MAC address.
func LocalMac() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, inter := range interfaces {
		address, err := inter.Addrs()
		if err != nil {
			return "", err
		}

		for _, address := range address {
			// check the address type and if it is not a loopback the display it
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return inter.HardwareAddr.String(), nil
				}
			}
		}
	}

	return "", errors.New("can't get local mac")
}

///*
//	获取外网ip
//*/
//func GetWwwIP() (ip string) {
//	ip = ""
//	resp, err := http.Get("http://myexternalip.com/raw")
//	if err != nil {
//		mylog.Error(err)
//		return
//	}

//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return
//	}

//	ip = string(body)
//	ip = strings.Split(ip, "\n")[0]
//	return
//}

// GetWwwIP 获取公网IP地址
func GetWwwIP() (exip string) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(bytes.TrimSpace(b))
}

// GetClientIP 获取用户ip
func GetClientIP(r *http.Request) (ip string) {
	ip = r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return
}

// 爬虫 ip138 获取 ip 地理信息
// ~~~~~~ 暂时废弃，采用 IPIP
func GetIp138(ip string) string {
	result := ""
	xutil.TryCatch(func() {
		resp, err := http.Get("http://ip138.com/ips138.asp?ip=" + ip)
		xerror.Panic(err)

		defer resp.Body.Close()
		input, _ := ioutil.ReadAll(resp.Body)

		out := mahonia.NewDecoder("gbk").ConvertString(string(input))

		reg := regexp.MustCompile(`<ul class="ul1"><li>\W*`)
		arr := reg.FindAllString(string(out), -1)
		str1 := strings.Replace(arr[0], `<ul class="ul1"><li>本站数据：`, "", -1)
		str2 := strings.Replace(str1, `</`, "", -1)
		str3 := strings.Replace(str2, `  `, "", -1)
		str4 := strings.Replace(str3, " ", "", -1)
		result = strings.Replace(str4, "\n", "", -1)

		if result == "保留地址" {
			result = "本地IP"
		}
	}, func(err error) {
		log.Println("IP138", "127.0.0.1", "读取 ip138 内容异常")
	})

	return result
}

// GetLocalIP 获取内网ip
func GetLocalIP1() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return
}

func GetLocalIp() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return ""
}
