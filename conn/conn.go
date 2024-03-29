package conn

import (
	"errors"
	"io"
	"strings"
	"syscall"

	"git.cyru1s.com/cyru1s/http3proxy/common"
)

type Conn interface {
	io.ReadWriteCloser

	Name() string

	Info() string

	Dial(dst string) (Conn, error)

	Listen(dst string) (Conn, error)
	Accept() (Conn, error)
}

func NewConn(proto string) (Conn, error) {
	proto = strings.ToLower(proto)
	if proto == "tcp" {
		return &TcpConn{}, nil
	} else if proto == "rhttp" {
		return &RhttpConn{}, nil
	} else if proto == "rhttp3" {
		return &Rhttp3Conn{}, nil
	}
	return nil, errors.New("undefined proto " + proto)
}

func SupportReliableProtos() []string {
	ret := make([]string, 0)
	ret = append(ret, "tcp")
	ret = append(ret, "rudp")
	ret = append(ret, "ricmp")
	ret = append(ret, "kcp")
	ret = append(ret, "quic")
	ret = append(ret, "rhttp")
	ret = append(ret, "rhttp3")
	return ret
}

func SupportProtos() []string {
	ret := make([]string, 0)
	ret = append(ret, SupportReliableProtos()...)
	ret = append(ret, "udp")
	return ret
}

func HasReliableProto(proto string) bool {
	return common.HasString(SupportReliableProtos(), proto)
}

func HasProto(proto string) bool {
	return common.HasString(SupportProtos(), proto)
}

var gControlOnConnSetup func(network, address string, c syscall.RawConn) error

func RegisterDialerController(fn func(network, address string, c syscall.RawConn) error) {
	gControlOnConnSetup = fn
}
