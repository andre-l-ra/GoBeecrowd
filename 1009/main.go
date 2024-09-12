package main

import "fmt"

func main() {
	var seller string
	var fixedSalary, salesMade float64
	const commission = 0.15
	fmt.Scan(&seller, &fixedSalary, &salesMade)

	totalSalary := fixedSalary + (salesMade * commission)

	fmt.Printf("TOTAL = R$ %2.2f\n", totalSalary)
}
