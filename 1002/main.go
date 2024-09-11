package main

import "fmt"

func main() {
	var raio float64
	fmt.Scanln(&raio)
	const n = 3.14159
	area := n * (raio * raio)

	fmt.Printf("A=%.4f\n", area)
}
