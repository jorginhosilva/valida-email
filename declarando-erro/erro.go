package main

import (
	"fmt"
	"errors"
)

func main() {
	var erro error = errors.New("Deu ruim mano, revisa teu código!")
	fmt.Println(erro)
}