package main

import (
	"log"
	"mUDP"
)

func main() {
	port := uint16(4545)
	server := mUDP.Server{}
	server.IsConnected = func(node *mUDP.Node) {
		log.Println("connect: ", node)
	}
	server.IsDelete = func(node *mUDP.Node) {
		log.Println("delete: ", node)
	}
	server.IsRead = func(node *mUDP.Node, b []byte) {
		log.Println("read: ", node, "-->", string(b))
		server.Send(node, b)
	}
	log.Println("start server in:", port)
	if err := server.Run(port); err != nil {
		panic(err)
	}
}
