package ws

import "golang.org/x/net/websocket"

type Client struct {
	id         string
	remoteAddr string
	ws         *websocket.Conn
}

func (c *Client) Listen() error {
	return nil
}
