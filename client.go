package mUDP

import (
	"errors"
	"net"
)

type Client struct {
	*Config
	conn      *net.UDPConn
	IsConnect func()
	IsRead    func([]byte)
	IsError   func(error)
}

func (c *Client) Send(message []byte) {
	_, err := c.conn.Write(message)
	if err != nil {
		c.IsError(err)
	}
}

func (c *Client) Run(conf *Config) {
	c.Config = conf
	udpAddr, err := net.ResolveUDPAddr("udp", c.Host+":"+c.Port)
	if err != nil {
		c.IsError(err)
		return
	}
	c.conn, err = net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		c.IsError(err)
		return
	}
	defer c.conn.Close()
	c.IsConnect()
	for {
		c.handle()
	}
}

func (c *Client) handle() {
	buf := make([]byte, c.BufSize)
	n, err := c.conn.Read(buf[0:])
	if err != nil {
		c.IsError(err) //Выполняем обратную функцию
		return
	}
	if n >= c.BufSize {
		c.IsError(errors.New("Max BufSize")) //Выполняем обратную функцию
		return
	}

	//Выполняем обратную функцию
	c.IsRead(buf[0:n])
}
