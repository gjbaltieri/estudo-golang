package main

import "fmt"

func main() {
	myMap := map[string]map[int]string{
		"Como caminhar": {
			1: "Levantar",
			2: "Caminhar",
		},
	}

	nome := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Paulo",
		},
	}

	fmt.Println(myMap)
	fmt.Println(nome)
}
