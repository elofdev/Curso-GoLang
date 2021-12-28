package main

import "fmt"

// Criação de uma strutura chamado Carro
type Carro struct {
	Nome string
}

// metodo ander
func (c Carro) Andar() {
	fmt.Println(c.Nome, "Andou")
}

func main() {
	carro := Carro{
		Nome: "BMW",
	}
	carro.Andar()
}
