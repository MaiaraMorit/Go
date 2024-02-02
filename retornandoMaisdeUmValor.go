package main

func sum(a int, ) (int, int) {
	var quadrado = a * a
	var cubo = a * a * a

	return quadrado, cubo
}

func main() {
	var a int = 10
	quadrado, cubo := sum(a)

	println("O quadrado e o cubo de", a, "são:", quadrado, cubo)
}
