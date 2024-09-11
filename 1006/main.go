/*
Leia 3 valores, no caso, variáveis A, B e C, que são as três notas de um aluno.
A seguir, calcule a média do aluno, sabendo que a nota A tem peso 2, a nota B tem peso 3 e a nota C tem peso 5
Considere que cada nota pode ir de 0 até 10.0, sempre com uma casa decimal.
*/

package main

import "fmt"

func main() {
	var A, B, C float64

	// fmt.Scanln(&A)
	// fmt.Scanln(&B)
	// fmt.Scanln(&C)

	fmt.Scan(&A, &B, &C)

	MEDIA := ((A * 2) + (B * 3) + (C * 5)) / 10

	fmt.Printf("MEDIA = %2.1f\n", MEDIA)
}
