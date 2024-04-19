package main

import (
	"fmt"
	"math"
)

func hypot (x, y float64) float64 {
	return math.Sqrt(x * x + y * y)
}

func soma (n1, n2 int64) int64 {
	return n1 + n2
}

func calculosMatematicos (n1, n2 int64) (int64, int64) {
	soma := n1 + n2 
	subtracao := n1 - n2 
	return soma, subtracao
}

func main() {
	somar := soma(15, 20)
	fmt.Println(somar)
	fmt.Println(hypot(3, 4))

	var f = func (txt string) string {
		fmt.Println(txt)
		return(txt)
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(8, 3)
	fmt.Println(resultadoSoma, resultadoSubtracao)

}
