package ws

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type Server struct {
	port  string
	conns []Client
	quit  chan struct{}
}

func NewServer(port string) *Server {
	return &Server{
		port:  port,
		conns: make([]Client, 0),
		quit:  make(chan struct{}),
	}
}

func (s *Server) HandleNewConnection(c *websocket.Conn) error {
	return nil
}

type WsHandler struct {
	h echo.HandlerFunc
}
