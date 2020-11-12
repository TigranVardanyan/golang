package main
import (
	"fmt"
	"os"
	"io"
)

func main() {
	text := "Hello, Go"
	fileName := "hello2.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Unable to create file", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)

	hello2, err := os.Open("hello2.txt")
	if err != nil {
		fmt.Println("Unable to read file", err)
		os.Exit(1)
	}
	data := make([]byte, 64)

	for{
		n, err := hello2.Read(data)
		if err == io.EOF{       // End of File??
			break
		}
		fmt.Print(string(data[:n]))
	}
	fmt.Println()

	fmt.Println("Done...")
}