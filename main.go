package main

import "fmt"

type Animal interface {
	HacerSonido() string
}

type Perro struct {
}

type Gato struct {
}

func (P Perro) HacerSonido() string {
	return "woof woof"
}

func (G Gato) HacerSonido() string {
	return "meow meow"
}

func Imprimir(A Animal) {
	fmt.Println(A.HacerSonido())
}

func main() {
	gato := Gato{}
	perro := Perro{}
	Imprimir(gato)
	Imprimir(perro)
}
