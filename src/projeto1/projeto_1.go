package main

import "fmt"

const ebulicao = 373.0

func main() {

	tc := (ebulicao - 273.0)

	fmt.Println("Ponto de Ebulição da Água: ", ebulicao, "K")
	fmt.Println("A Conversão de Kelvin para Celsius é :", tc, "C")

}
