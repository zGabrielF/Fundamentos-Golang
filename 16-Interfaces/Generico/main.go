package main

import "fmt"

//Interface
// nesse caso a interface é genercia para utilizar a interface as funções não precisam implementar obrigatoriamente, pois não tem restrição

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	// nesse caso a interface aceita qualquer coisa, e pode se tornar essa bagunça toda, é usado em casos especificos
	generica("string")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "asdasd", false, true, float64(123124))

	mapa := map[interface{}]interface{}{
		1:            "biel",
		float32(100): true,
		true:         float64(12),
		"string":     "ssdasdd",
	}
	fmt.Println(mapa)

}
