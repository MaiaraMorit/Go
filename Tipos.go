package main

import "fmt"

//caso não saiba qual é o tio de uma determinada variável use o seguinte comando:
//fmtPrintf("Tipo: %T Valor: %v\n", nomeDaSuaVariável)

func intType() {
	idade := 45

	fmt.Println("Minha idade é", idade)
}

func stringType() {
	nome := "Maria"

	fmt.Println("Meu nome é", nome)
}

func floatType() {
	altura := 1.65

	fmt.Println("Minha altura é", altura)
}

func boolType() {
	verdadeiro := true

	fmt.Println("Verdadeiro é", verdadeiro)
}

func main() {
	intType()
	stringType()
	floatType()
	boolType()
}
