package main

import "github.com/gorilla/websocket"

type Client struct {
	Ws *websocket.Conn
}

func (c *Client) Send(msg []byte) error {
	return c.Ws.WriteMessage(websocket.TextMessage, msg)
}
