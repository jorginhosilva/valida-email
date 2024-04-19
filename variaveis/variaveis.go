package main

import (
	"fmt"
)

//Váriaveis de nível de pacote são inicializadas antes da função main executar.
var nivel string = "Nivel de pacote"

func main() {
	//Váriaveis locais são utilizadas à media que suas declarações são encontradas durante a execução de funções.
	local := "variavel locais"
	fmt.Println(local)
	fmt.Println(nivel)
}