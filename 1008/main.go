/*
Escreva um programa que leia o número de um funcionário,
seu número de horas trabalhadas,
o valor que recebe por hora e calcula o salário desse funcionário.
A seguir, mostre o número e o salário do funcionário, com duas casas decimais.
*/

package main

import "fmt"

func main() {
	var number, hoursWorked int64
	var salaryPerHours float64

	fmt.Scan(&number, &hoursWorked, &salaryPerHours)

	salary := salaryPerHours * float64(hoursWorked)

	fmt.Printf("NUMBER = %d\n", number)
	fmt.Printf("SALARY = U$ %2.2f\n", salary)
}
