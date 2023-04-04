package main

import (
	"fmt"
)

func multipleReturns(txt string, n1 int8, n2 int8) (string, int8) {
	var soma int8 = n1 * n2

	return txt, soma
}

func main() {
	var txt, soma = multipleReturns("texto da função", 5, 3)
	fmt.Println(txt, soma)
}
