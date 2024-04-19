package main

import (
	"fmt"
)

func main() {
	var numero int64 = 1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)
	
	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
}