package main

func sum(a int, ) (int, int) {
	var quadrado = a * a
	var cubo = a * a * a

	return quadrado, cubo
}

func sum2(a int, ) (quadrado int, cubo int) {
	quadrado = a * a
	cubo = a * a * a

	return
}

func main() {
	var a int = 10
	var b int = 20

	quadrado, cubo := sum(a)
	quadradoDeB, cuboDeB := sum2(b)

	println("O quadrado e o cubo de", a, "são:", quadrado, cubo)
	println("O quadrado e o cubo de", b, "são:", quadradoDeB, cuboDeB)
}