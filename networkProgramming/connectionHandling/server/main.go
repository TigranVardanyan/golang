package main

//func Listen(network, laddr string) (net.Listener, error)

import (
	"fmt"
"net"
)

func main() {
	message := "Hello, I am a server"   // отправляемое сообщение
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		fmt.Println("+1")
		conn.Close()
	}
}
