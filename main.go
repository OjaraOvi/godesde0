package main

import (
	"fmt"
	"github.com/OjaraOvi/godesde0/ejercicios"
	"github.com/OjaraOvi/godesde0/variables"
	"runtime"
)

func main() {
	fmt.Println("Wenaaa!!!")
	variables.ShowEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1555)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "darwin" {
		fmt.Println("Esto no es windows", os)
	} else {
		fmt.Println("Esto es windows", os)
	}

	numero, texto := ejercicios.ConvNumerico("120")
	fmt.Println(numero)
	fmt.Println(texto)

	ejercicios.TablaDeMultiplicar()

}
