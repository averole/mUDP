package mUDP

import (
	"net"
)

type Client struct {
	conn    *net.UDPConn
	isClose bool
}

func (c *Client) Send(message []byte) error {
	_, err := c.conn.Write(message)
	return err
}

func NewClient(adres string) *Client {
	udpAddr, err := net.ResolveUDPAddr("udp", adres)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		panic(err)
	}
	return &Client{conn: conn, isClose: false}
}
func (c *Client) Close() {
	c.isClose = false
	c.conn.Close()
}

func (c *Client) Listen(read func([]byte)) {
	buffer := make([]byte, maxBufferSize)
	for {
		if c.isClose == true {
			return
		}
		n, err := c.conn.Read(buffer)
		if err != nil {
			panic(err)
		}
		read(buffer[0:n])
	}
}
