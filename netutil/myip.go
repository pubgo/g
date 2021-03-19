package netutil

import (
	"context"
	"errors"
	"net"
	"strconv"
)

// NameServer the name server to use for this lib
var NameServer = "ns1.google.com:53"

// GetInterfaceIP get the ip of your interface, useful when you want to
// get your ip inside a private network, such as wifi network.
func GetInterfaceIP() (string, error) {
	conn, err := net.Dial("udp", NameServer)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

// GetPublicIP get the ip that is public to global.
func GetPublicIP() (string, error) {
	r := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", NameServer)
		},
	}
	txt, err := r.LookupTXT(context.Background(), "o-o.myaddr.l.google.com")
	if err != nil {
		return "", err
	}

	if len(txt) == 0 {
		return "", errors.New("[myip] can't get a ip")
	}

	return txt[0], nil
}


// Get a free port.
func Get() (port int, err error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	addr := listener.Addr().String()
	_, portString, err := net.SplitHostPort(addr)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(portString)
}
