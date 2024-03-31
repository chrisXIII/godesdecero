package main

import (
	"fmt"

	"github.com/chrisXIII/godesdecero/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado,texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
