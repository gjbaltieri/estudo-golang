package main

import "fmt"

func main() {

	var num = 12
	if num >= 10 {
		fmt.Println("o numero", num, "é maior que 10!!!")
	} else {
		fmt.Println("o numero", num, "é menor que 10")
	}

	var numero = 2
	if ifInit := numero; ifInit > 0 {
		fmt.Println("esse é o If init, a var ifInit só é valida no scopo do if")
	} else {
		fmt.Println("também pode ter um else")
	}
}
