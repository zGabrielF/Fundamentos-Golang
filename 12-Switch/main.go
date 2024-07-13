package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sabado"
	default:
		return "numero inválido"
	}
}

func diaDaSemana(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
		//fallthrough // faz ele ir para a proxima condição, no caso pula para segunda feira
	case numero == 2:
		return "segunda"
	case numero == 3:
		return "terça"
	case numero == 4:
		return "quarta"
	case numero == 5:
		return "quinta"
	case numero == 6:
		return "sexta"
	case numero == 7:
		return "sabado"
	default:
		return "numero inválido"
	}
}
func main() {

	dia := diaSemana(1)
	fmt.Println(dia)

	dia2 := diaDaSemana(3)
	fmt.Println(dia2)

}
