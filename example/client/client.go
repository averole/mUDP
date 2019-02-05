package main

import (
	"fmt"
	"mUDP"
)

func main() {
	cl := mUDP.NewClient("localhost:4545")
	go cl.Listen(func(b []byte) {
		fmt.Println(string(b))
	})
	cl.Send([]byte("auth"))
	for {
	}
}
