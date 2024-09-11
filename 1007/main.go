// Leia quatro valores inteiros A, B, C e D.
// A seguir, calcule e mostre a diferença do produto de A e B pelo produto de C e D segundo a fórmula: DIFERENCA = (A * B - C * D).

package main

import "fmt"

func main() {
	var A, B, C, D int64
	fmt.Scan(&A, &B, &C, &D)

	dif := (A * B) - (C * D)

	fmt.Printf("DIFERENCA = %d\n", dif)
}
