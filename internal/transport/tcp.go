package transport

import (
	"net"
)

func StartServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func Connect(addr string) (net.Conn, error) {
	return net.Dial("tcp", addr)
}