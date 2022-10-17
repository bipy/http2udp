package main

import (
	"fmt"
	"net"
	"testing"
)

func Test(t *testing.T) {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	})

	if err != nil {
		panic(err.Error())
	}

	go func(conn *net.UDPConn) {
		for {
			buf := make([]byte, 100)
			n, err := listener.Read(buf)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(string(buf[:n]))
		}
	}(listener)

	main()

}
