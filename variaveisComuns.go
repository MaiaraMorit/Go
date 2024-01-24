package main

import "fmt"

//:= significa var

func main() {
	nome, email, idade, altura, peso := "Maria", "maria@hotmail.com", 30, 1.65, 65.0

	fmt.Print("Meu nome é ", nome, " tenho ", idade, " anos, meu email é ", email, " minha altura é ", altura, " e meu peso é ", peso, "kg.")
}

// go run variaveisComuns.go
