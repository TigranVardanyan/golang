package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//buffer - Writer
//func (b *Writer) Write(p []byte) (nn int, err error)
//func (b *Writer) WriteByte(c byte) error
//func (b *Writer) WriteRune(r rune) (size int, err error)
//func (b *Writer) WriteString(s string) (int, error)

//buffer - Reader
//func (b *Reader) Read(p []byte) (n int, err error)
//func (b *Reader) ReadByte() (byte, error)
//func (b *Reader) ReadBytes(delim byte) ([]byte, error)
//func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
//func (b *Reader) ReadRune() (r rune, size int, err error)
//func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
//func (b *Reader) ReadString(delim byte) (string, error

func main() {
	filename := "buffer.txt"
	write(filename)
	read(filename)
}

func write(filename string)  {
	rows := []string{
		"Hello go!",
		"Welcome to golang",
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}
	writer.Flush()
}
func read(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Print(line)
	}
}