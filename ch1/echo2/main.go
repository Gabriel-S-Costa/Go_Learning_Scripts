package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""  // inialização das variaveis com o tipo inferido
	for _, arg := range os.Args[1:] { // laço for para iterar sobre a lista com os params informados
		s += sep + arg
		sep = " "
	}

	fmt.Println(s) // print da nova string concatenada
}
