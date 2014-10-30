package main

import (
	"github.com/opentable/logging-daemon/logger"
)

func main() {
	var buf [10240]byte

	in := make(chan string)
	out := make(chan string, 1000)
	rb := logger.NewRingBuffer(in, out)
	listener := logger.NewListener()
	forwarder := logger.NewForwarder()

	go rb.Run()

	go func() {
		for {
			v := <-out
			forwarder.Forward(v)
		}
	}()

	for {
		in <- listener.Read(buf)
	}
}
