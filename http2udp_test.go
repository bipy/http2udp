package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"testing"
)

func TestServe(t *testing.T) {
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

	http.HandleFunc("/", handler)
	t.Error(http.ListenAndServe("0.0.0.0:8080", nil))
}

func TestSend(t *testing.T) {
	client := http.Client{}
	data, _ := json.Marshal(&Req{
		IP:   "127.0.0.1",
		Msg:  "Test",
		Port: 8080,
	})
	resp, err := client.Post("http://localhost:8080", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}

	data, _ = json.Marshal(&Req{
		IP:   "127.0.0.1",
		Msg:  "Test2",
		Port: 8080,
	})
	resp, err = client.Post("http://localhost:8080", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}
}
