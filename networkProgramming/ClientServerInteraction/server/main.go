package main
import (
	"fmt"
	"net"
)
var dict = map[string]string{
	"red": "красный",
	"green": "зеленый",
	"blue": "синий",
	"yellow": "желтый",
	"purple": "пурпурный",
	"orange":"оранжевый",
	"grey":"серый",
	"white":"белый",
	"lime":"салатовый",
}

func main() {
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
			conn.Close()
			continue
		}
		go handleConnection(conn)  // run the goroutine to process the request
	}
}
// connection handling
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// read the data received in the request
		input := make([]byte, 1024 * 4)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		source := string(input[0:n])
		// based on the data received, we get the translation from the dictionary
		target, ok := dict[source]
		if ok == false{             // if the data is not found in the dictionary
			target = "undefined"
		}
		// display diagnostic information on the server console
		fmt.Println(source, "-", target)
		// sending data to the client
		conn.Write([]byte(target))
	}
}