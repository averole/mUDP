package mUDP

import (
	"net"
)

type Client struct {
	conn   *net.UDPConn
	IsRead func([]byte)
}

func (c *Client) Send(message []byte) error {
	_, err := c.conn.Write(message)
	return err
}

func (c *Client) Run(adres string) {
	udpAddr, err := net.ResolveUDPAddr("udp", adres)
	if err != nil {
		panic(err)
	}
	c.conn, err = net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		panic(err)
	}
	defer c.conn.Close()
	buffer := make([]byte, maxBufferSize)
	for {
		n, err := c.conn.Read(buffer)
		if err != nil {
			panic(err)
		}
		c.IsRead(buffer[0:n])
	}
}

// func (c *Client) handle() {
// 	buf := make([]byte, maxBufferSize)
// 	n, err := c.conn.Read(buf[0:])
// 	if err != nil {
// 		c.IsError(err) //Выполняем обратную функцию
// 		return
// 	}
// 	if n >= maxBufferSize {
// 		c.IsError(errors.New("Max BufSize")) //Выполняем обратную функцию
// 		return
// 	}

// 	//Выполняем обратную функцию
// 	c.IsRead(buf[0:n])
// }
