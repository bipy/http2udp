package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

type Req struct {
	IP   string `json:"ip"`
	Msg  string `json:"msg"`
	Port int    `json:"port"`
}

func main() {
	port := flag.Int("p", 8080, "port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &Req{}
	err = json.Unmarshal(data, req)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(req.IP),
		Port: req.Port,
	})
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}

	defer socket.Close()
	_, err = socket.Write([]byte(req.Msg))
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
