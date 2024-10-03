package main

import "fmt"

type peca struct {
	Codigo        int64
	Quantidade    int64
	ValorUnitario float64
}

func main() {

	peca1 := peca{}
	peca2 := peca{}

	fmt.Scan(&peca1.Codigo, &peca1.Quantidade, &peca1.ValorUnitario, &peca2.Codigo, &peca2.Quantidade, &peca2.ValorUnitario)

	total := (peca1.ValorUnitario * float64(peca1.Quantidade)) + (peca2.ValorUnitario * float64(peca2.Quantidade))

	fmt.Printf("VALOR A PAGAR: R$ %2.2f\n", total)
}
