package main

import (
	"fmt"
	"io"
)

// io.Reader
//type Reader interface {
//	Read (p []byte) (n int, err error)
//}

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(ph); i++ {
		if (ph[i] >= '0' && ph[i] <= '9') {
			p[count] = ph[i]
			count++
		}
	}
	return count, io.EOF
}

type phoneWriter struct {

}

func (p phoneWriter) Write(bs []byte) (int, error)  {
	if len(bs) == 0 {
		return 0,nil
	}
	for i := 0; i < len(bs); i++ {
		if (bs[i] >= '0' && bs[i] <= '9') {
			fmt.Print(string(bs[i]))
		}
	}
	fmt.Println()
	return len(bs), nil
}


func main() {
	phone1 := phoneReader("+1(234)56789012")
	phone2 := phoneReader("+2(345)67890123")

	buffer := make([]byte, len(phone1))
	phone1.Read(buffer)
	fmt.Println(string(buffer))

	buffer = make([]byte, len(phone2))
	phone2.Read(buffer)
	fmt.Println(string(buffer))

	fmt.Println()

	bytes1 := []byte("+12456789 78(52)")
	bytes2 := []byte("+a(13456) 45 78")

	write := phoneWriter{}
	write.Write(bytes1)
	write.Write(bytes2)
}