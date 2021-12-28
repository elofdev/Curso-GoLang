package main

import "fmt"

func main() {

	// Valores do range x...quantos forem informados
	resultado := SomaTudo(10, 13, 15, 100, 333)

	fmt.Println(resultado)
}

// SomaTudo é uma Função variadicas ou function variadic e qua soma todos os valores indormados no range x

func SomaTudo(x ...int) int {
	resultado := 0
	// faz um loop em todo o range x e pegara o valor esomara todos
	for _, v := range x {
		resultado += v
	}
	return resultado
}
