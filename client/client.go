package h3pclient

import (
	"fmt"

	proxy "git.cyru1s.com/cyru1s/http3proxy/proxy"
)

func StartClient(cdn_ip_port string, sni_name string, host_name string) int {
	var proxyprotostr []string
	proxyprotostr = append(proxyprotostr, "tcp")
	var fromaddr []string
	fromaddr = append(fromaddr, ":8080")
	_, err := proxy.NewClient(nil, "rhttp3", cdn_ip_port, "test", "REVERSE_SOCKS5", proxyprotostr, fromaddr, nil)
	if err != nil {
		fmt.Println("err")
	}
	return 1
}
