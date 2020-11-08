package main

import "fmt"

//type Stream interface {
//	read() string
//	write(string)
//	close()
//}
//
//func writeToStream(stream Stream, text string)  {
//	stream.write(text)
//}
//
//func closeStream(stream Stream)  {
//	stream.close()
//}
//
//type File struct {
//	text string
//}
//
//type Folder struct {}
//
//func (f *File) read() string{
//	return f.text
//}
//
//func (f *File) write(message string){
//	f.text = message
//	fmt.Println("Запись в файл строки", message)
//}
//
//func (f *File) close(){
//	fmt.Println("Файл закрыт")
//}
//
//func (f *Folder) close(){
//	fmt.Println("Папка закрыта")
//}
//
//func main() {
//
//	myFile := &File{}
//	myFolder := &Folder{}
//
//	writeToStream(myFile, "hello world")
//	closeStream(myFile)
//	//closeStream(myFolder)     // !!Folder does not implement Stream
//	myFolder.close()            // Так можно
//}

// Implementing multiple interfaces

type Reader interface{
	read()
}

type Writer interface{
	write(string)
}

func writeToStream(writer Writer, text string){
	writer.write(text)
}
func readFromStream(reader Reader){
	reader.read()
}

type File struct{
	text string
}

func (f *File) read(){
	fmt.Println(f.text)
}
func (f *File) write(message string){
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func main() {

	myFile := &File{}
	writeToStream(myFile, "hello world")
	readFromStream(myFile)
}