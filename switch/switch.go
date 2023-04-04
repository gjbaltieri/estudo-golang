package main

import "fmt"

func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça-feira"
	case 4:
		return "quarta-feira"
	case 5:
		return "quinta-feira"
	case 6:
		return "sexta-feira"
	case 7:
		return "sabado"
	default:
		return "numero inválido"
	}
}

func main() {

	funcDia := diaDaSemana(7)
	fmt.Println(funcDia)
}
