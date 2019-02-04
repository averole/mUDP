package mUDP

import (
	"net"
	"strconv"
)

type Node struct {
	addr  *net.UDPAddr
	index string
}

type Server struct {
	bunch map[string]*Node

	*Config
	conn *net.UDPConn

	IsNewClient    func(*Node)
	IsDeleteClient func(*Node)
	IsRead         func(*Node, []byte)
	IsError        func(error)
}

// Send Send to client
func (s *Server) Send(c *Node, message []byte) {
	_, err := s.conn.WriteToUDP(message, c.addr)
	if err != nil {
		s.IsError(err) //Выполняем обратную функцию
		return
	}
}

//Delete client
func (s *Server) Delete(c *Node) {
	s.IsDeleteClient(c) //Выполняем обратную функцию
	//удаляем
	delete(s.bunch, c.index)
}

//Run server
func (s *Server) Run(c *Config) {
	s.bunch = make(map[string]*Node)
	s.Config = c
	udpAddr, err := net.ResolveUDPAddr("udp", c.Host+":"+c.Port)
	if err != nil {
		s.IsError(err)
		return
	}
	s.conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		s.IsError(err)
		return
	}
	defer s.conn.Close()
	for {
		s.handle()
	}
}

//Обработка входящих сообщений от клиентов
func (s *Server) handle() {
	buf := make([]byte, s.BufSize)
	n, addr, err := s.conn.ReadFromUDP(buf)

	if err != nil {
		s.IsError(err) //Выполняем обратную функцию
		return
	}

	//индексный ключ клиента
	i := addr.IP.String() + strconv.Itoa(addr.Port)

	//if new client
	if _, ok := s.bunch[i]; ok == false {
		c := &Node{addr: addr, index: i}
		s.bunch[i] = c
		s.IsNewClient(c)
	}

	//Выполняем обратную функцию
	s.IsRead(s.bunch[i], buf[0:n])
}
