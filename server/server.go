package server

import (
	"net"
	"../error"
)

func Listen(port string) (net.UDPConn) {
	ServerAddr,err := net.ResolveUDPAddr("udp", port)
	error.CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	error.CheckError(err)

	return *ServerConn
}