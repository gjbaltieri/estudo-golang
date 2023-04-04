package pacote2

import "fmt"

func Pacote2() {
	fmt.Println("pacote2")
	myMap := map[string]map[string]string{
		"nome": {
			"primeiro": "joão",
			"ultimo":   "paulo",
		},
	}
	fmt.Println(myMap)
	auxiliar()
}
