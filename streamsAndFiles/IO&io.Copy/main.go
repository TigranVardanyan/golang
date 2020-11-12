package main

import (
	"fmt"
	"io"
	"os"
)

//os.Stdin - in
//os.Stdout - out
//os.Stderr - err

func main() {
	file,err := os.Create("hello3.txt")
	if err != nil {
		println("Unable top create file", err)
		os.Exit(1)
	}
	file.WriteString("this is go file")

	defer file.Close()
	file2,err := os.Open("hello.txt")
	if err != nil {
		println("Unable top create file", err)
		os.Exit(1)
	}
	defer file2.Close()
	io.Copy(os.Stdout, file2)
	fmt.Println()

	//file, err := os.Open("Hello.txt")
	//if err != nil{
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//defer file.Close()
	//
	//io.Copy(os.Stdout, file)
}

//type phoneReader string
//
//func (p phoneReader) Read(bs []byte) (int, error){
//	count := 0
//	for i := 0; i < len(p); i++{
//		if(p[i] >= '0' && p[i] <= '9'){
//			bs[count] = p[i]
//			count++
//		}
//	}
//	return count, io.EOF
//}
//
//func main() {
//	phone1 := phoneReader("+1(234)567 90-10")
//	io.Copy(os.Stdout, phone1)
//	fmt.Println()
//}