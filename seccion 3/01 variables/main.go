package main //este paquete siempre debe estar presente

//paquetes importados
import "fmt" //las llaves o parentesis solo aplican cuando son varios paquetes

func main() {
	var myIntVar int = 27                     //enteros
	var myGreeting string = "How are things?" //string
	var amImarried bool = true                //booleanos
	var myWeightAt2018 float64 = 65.5         //float64

	fmt.Println("This is my age, when it was 2018:", myIntVar, ". How long time has passed")
	fmt.Println(myGreeting)
	fmt.Println("Are you asked if i am married? I will answer like a programmer:", amImarried)
	fmt.Println("Can you guees my weight at 2018?, If you not, I tell you, my weight was: ", myWeightAt2018)
	fmt.Println("THE ADDRESS MEMORY FOR myIntVar IS: ", &myIntVar)

	aWrongProctice := 34
	fmt.Println("This is a wrong practice, but it works:", aWrongProctice)

}
