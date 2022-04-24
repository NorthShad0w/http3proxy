package h3pclient

import (
	"fmt"

    
)

func StartClient(cdn_ip string,sni_name string,hostname string) int {
	fmt.Println(cdn_ip,sni_name,hostname)
    _, err := proxy.NewClient(config, protos[0], *server, *name, clienttypestr, proxyproto, fromaddr, toaddr)
	return 1
}
