package main

import "fmt"
import "unsafe"

func main() {
	var abelardo int16 = 5
	var berenice float64 = 3.1415
	var carlos string = "Hello, Go!"
	var diana bool = false

	fmt.Printf("Integer: %7.2d\t TypeOfData: %T\t SizeOfBytes: %d\t\t  Bits: %d\n", abelardo, abelardo, unsafe.Sizeof(abelardo), unsafe.Sizeof(abelardo)*8)
	fmt.Printf("Float: %12.5f\t TypeOfData: %T\t SizeOfBytes: %d\t\t  Bits: %d\n", berenice, berenice, unsafe.Sizeof(berenice), unsafe.Sizeof(berenice)*8)
	fmt.Printf("String: %11s\t TypeOfData: %T\t SizeOfBytes: %d\t  Bits: %d\n", carlos, carlos, unsafe.Sizeof(carlos), unsafe.Sizeof(carlos)*8)
	fmt.Printf("Boolean: %8t\t TypeOfData: %T\t SizeOfBytes: %d\t\t  Bits: %d\n", diana, diana, unsafe.Sizeof(diana), unsafe.Sizeof(diana)*8)
}
