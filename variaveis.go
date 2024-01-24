package main

import "fmt"

//Os tipos de variáveis mais comuns em Go são:
//
//int: números inteiros
//float64: números de ponto flutuante
//string: strings
//bool: valores booleanos
//array: arrays
//slice: slices
//map: mapas
//struct: structs

func main() {
	var nome string = "Maria"
	var idade int = 30
	var email string = "maria@hotmail.com"
	var altura float64 = 1.65
	var peso float64 = 65.0

	fmt.Print("Meu nome é ", nome, " tenho ", idade, " anos, meu email é ", email, " minha altura é ", altura, " e meu peso é ", peso, "kg.")
}

//go run variaveis.go
