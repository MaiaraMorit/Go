package main

import "fmt"

func main() {

	var nome, email string = "Maria", "maria@hotmail.com"
	var idade int = 30
	var altura, peso float64 = 1.65, 65.0

	fmt.Print("Meu nome é ", nome, " tenho ", idade, " anos, meu email é ", email, " minha altura é ", altura, " e meu peso é ", peso, "kg.")
}

// go run variaveisMesmoTipo.go
