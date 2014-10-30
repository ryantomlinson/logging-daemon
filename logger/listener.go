package logger

import (
	"fmt"
	"net"
	"os"
)

// Listener ...
type Listener struct {
	sock *net.UDPConn
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

// NewListener : initialises a new udp listener
func NewListener() *Listener {
	addr, err := net.ResolveUDPAddr("udp4", ":6380")
	checkErr(err)
	sock, err := net.ListenUDP("udp4", addr)
	checkErr(err)

	return &Listener{
		sock: sock,
	}
}

func (l *Listener) Read(buf [10240]byte) string {
	l.sock.ReadFromUDP(buf[:])
	return string(buf[:len(buf)])
}
