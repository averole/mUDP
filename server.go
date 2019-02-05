package mUDP

import (
	"fmt"
	"net"
	"time"
)

const maxBufferSize = 1024

type void struct{}
type Session string
type Node struct {
	addr     *net.UDPAddr
	session  Session
	deadline int64
}

func (n *Node) String() string {
	return string(n.session)
}

type Server struct {
	hub  map[Session]*Node
	conn *net.UDPConn

	IsConnected func(*Node)
	IsRead      func(*Node, []byte)
	IsDelete    func(*Node)
}

//Send Send to client
func (s *Server) Send(node *Node, message []byte) {
	fmt.Println(node.deadline, time.Now().Unix())
	if node.deadline < time.Now().Unix() {
		s.Delete(node)
		return
	}
	_, err := s.conn.WriteToUDP(message, node.addr)
	if err != nil {
		return
	}
}

// Delete client
func (s *Server) Delete(node *Node) {
	s.IsDelete(node)
	delete(s.hub, node.session)
}

//Run server
func (s *Server) Run(port uint16) (err error) {
	s.hub = make(map[Session]*Node)
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return err
	}
	s.conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer s.conn.Close()

	buffer := make([]byte, maxBufferSize)
	for {
		var node *Node
		n, addr, err := s.conn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}
		session := Session(fmt.Sprintf("%s:%d|%s", addr.IP.String(), addr.Port, addr.Zone))
		if _, ok := s.hub[session]; ok == false {
			//new node
			node = &Node{}
			node.session = session
			node.addr = addr
			s.hub[session] = node
			s.IsConnected(node)
		} else {
			node = s.hub[session]
		}
		node.deadline = time.Now().Unix() + 5
		s.IsRead(node, buffer[0:n])
	}
}
