package main

import "fmt"

func main() {
	var preco float64
	var escolha int
	fmt.Print("Digite o preço do produto: R$ ")
	fmt.Scan(&preco)
	fmt.Println("Escolha a forma de pagamento")
	fmt.Println("1- À vista, dinheiro ou cheque, 10% de desconto")
	fmt.Println("2- À vista, cartão de crédito, 5% de desconto")
	fmt.Println("3- Em 2 vezes, preço normal de etiqueta sem juros")
	fmt.Println("4- Em 3 vezes, preço normal de etiqueta + 10% de juros")
	fmt.Scan(&escolha)
	if escolha > 5 {
		fmt.Println("Escolha Inválida")
		return
	}
	if escolha == 1 {
		vista10 := preco * 0.10
		final := preco - vista10
		fmt.Printf("Total a pagar: R$ %.2f\n", final)
	}
	if escolha == 2 {
		vista5 := preco * 0.05
		final := preco - vista5
		fmt.Printf("Total a pagar: R$ %.2f\n", final)
	}
	if escolha == 3 {
		semjuros := preco / 2
		fmt.Printf("Total a pagar: 2x de R$ %.2f\n", semjuros)
	}
	if escolha == 4 {
		comjuros := preco * 0.1*3
		total := preco + comjuros
		final := total / 3
		fmt.Printf("Total a pagar: 3x de R$ %.2f\n", final)
	}
}
