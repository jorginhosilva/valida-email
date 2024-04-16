package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
    email := "ma650696@gmail.com"
    err := checkmail.ValidateFormat(email)
    if err != nil {
        fmt.Println("Erro ao validar o formato do e-mail:", err)
    } else {
        fmt.Println("Formato do e-mail é válido.")
    }
}
