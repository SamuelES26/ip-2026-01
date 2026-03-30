package main

import "fmt"

func main() {
	var val float64
	fmt.Println("Qual o valor da compra?")
	fmt.Scan(&val)
	if val < 0 {
		fmt.Println("Valor Inválido")
	}
	if val < 10.00 {
		lu70 := val * 1.30
		fmt.Printf("O valor da venda é de %2.f\n", lu70)
	}
	if 10.00 <= val && val < 30.00 {
		lu50 := val * 1.50
		fmt.Printf("O valor da venda é de %2.f\n", lu50)
	}
	if 30.00 <= val && val < 50.00 {
		lu40 := val * 1.40
		fmt.Printf("O valor da venda é de %2.f\n", lu40)
	}
	if val > 50.00 {
		lu30 := val * 1.30
		fmt.Printf("O valor da venda é de %2.f\n", lu30)
	}
}
