package main

import (
	"fmt"
	"mUDP"
	"time"
)

func main() {
	cl := mUDP.Client{}
	cl.IsRead = func(b []byte) {
		fmt.Println(string(b))

	}
	go cl.Run("127.0.0.1:4545")
	time.Sleep(1 * time.Second)
	cl.Send([]byte("auth"))
	for {

	}
}
