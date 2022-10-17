package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

	app.POST("/", handler)

	log.Fatal(app.Start(fmt.Sprintf(":%d", *port)))
}

func handler(c echo.Context) error {
	req := &Req{}
	err := c.Bind(req)
	if err != nil || req.Msg == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(req.IP),
		Port: req.Port,
	})
	if err != nil {
		return c.NoContent(http.StatusBadGateway)
	}
	defer socket.Close()
	_, err = socket.Write([]byte(req.Msg))
	if err != nil {
		return c.NoContent(http.StatusBadGateway)
	}
	return c.NoContent(http.StatusOK)
}
