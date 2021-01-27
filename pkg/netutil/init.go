package netutil

import (
	"github.com/pubgo/x/xerror"
	"net"
)

func IpAddress() string {
	addrs, err := net.InterfaceAddrs()
	xerror.PanicM(err, "net.InterfaceAddrs error")

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
