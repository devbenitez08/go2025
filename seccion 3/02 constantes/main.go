package main

import "fmt"
import "reflect" //se requiere para usar TypeOf

func main() {

	const gravityConstant float64 = 9.81
	var aStoneWeight float64 = 1
	var height int = 100

	var speed float64 = aStoneWeight * gravityConstant * float64(height) //cast

	//imprimiendo con println
	fmt.Println("An a stone has been drooped from an airplain from 100 meters, its cinetic force is [m*g*h]: ", speed)
	fmt.Println("Constant type data: ", reflect.TypeOf(gravityConstant))

	//imprimiendo con printf
	fmt.Printf("Type of data constant: (%%T) : %T\n", gravityConstant) //otra forma de obtener el tipo de dato declarado
}
