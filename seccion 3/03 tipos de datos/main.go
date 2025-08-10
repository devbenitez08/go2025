package main

import "fmt"
import "unsafe"

func main() {
	var abelardo int16 = 5
	var berenice float64 = 3.1415
	var carlos string = "Hello, Go!"
	var diana bool = false
	var messageForEveryone string = `Welcome to the Go programming language!
Enjoy learning about data types and memory usage.
This is a multi-line message.
Not use tabs for indentation this multi-line message. :)
Use backticks for multi-line strings.
Have a great day!`

	fmt.Printf("Integer: %7.2d\t TypeOfData: %T\t SizeOfBytes: %d\t\t  Bits: %d\n", abelardo, abelardo, unsafe.Sizeof(abelardo), unsafe.Sizeof(abelardo)*8)
	fmt.Printf("Float: %12.5f\t TypeOfData: %T\t SizeOfBytes: %d\t\t  Bits: %d\n", berenice, berenice, unsafe.Sizeof(berenice), unsafe.Sizeof(berenice)*8)
	fmt.Printf("String: %11s\t TypeOfData: %T\t SizeOfBytes: %d\t  Bits: %d\n", carlos, carlos, unsafe.Sizeof(carlos), unsafe.Sizeof(carlos)*8)
	fmt.Printf("Boolean: %8t\t TypeOfData: %T\t SizeOfBytes: %d\t\t  Bits: %d\n", diana, diana, unsafe.Sizeof(diana), unsafe.Sizeof(diana)*8)

	fmt.Println()
	fmt.Println(messageForEveryone)
}
