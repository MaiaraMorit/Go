package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 5.5

	var c int = int(b) // aqui eu converti o valor de b para int e armazenei na variável c

	fmt.Println("O valor de C é:", c + a)
	// o valor de C arredoado para baixo é 5, somado com 10 é igual a 15
}
