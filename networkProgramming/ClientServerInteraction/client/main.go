package main
import (
	"fmt"
	"net"
)
func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for{
		var source string
		fmt.Print("Enter a color: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Incorrect output", err)
			continue
		}
		// send message to client
		if n, err := conn.Write([]byte(source));
			n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		// receive response
		fmt.Print("Translate:")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err !=nil{ break}
		fmt.Print(string(buff[0:n]))
		fmt.Println()
	}
}