package main

import ("fmt"
		)

func main() {
	var conta, tipo int
	var gasto float64

	fmt.Println("Digite sua conta:")
	fmt.Scan(&conta)
	fmt.Println("     Escolha seu tipo de localidade:")
	fmt.Println("1- Residencial 2- Comercial 3- Industrial")
	fmt.Scan(&tipo)
	if tipo > 3 {
		fmt.Println("Escolha não válida")
		return
	}
	fmt.Println("Digite seu gasto em M³:")
	fmt.Scan(&gasto)
	if tipo == 1 {
		resi := 5.00 + (gasto*0.05)
		fmt.Printf("Sua conta é %d, o valor a se pagar é: R$ %.2f.\n", conta, resi)
	}
	if tipo == 2 {
		if gasto <= 80 {
			fmt.Printf("Sua conta é %d, o valor se pagar é: R$ 500.00.\n", conta)
		} else {
			come := 500 + (gasto*0.25)
			fmt.Printf("Sua conta é %d, o valor a se pagar é: R$ %.2f.\n", conta, come)
		}
	}
	if tipo == 3 {
		if gasto <= 100 {
			fmt.Printf("Sua conta é %d, o valor a se pagar é: R$ 800.00.\n", conta)
		} else {
			indus := 800 + (gasto*0.04)
			fmt.Printf("Sua conta é %d, o valor a se pagar é: R$ %.2f.\n", conta, indus)
		}
	}
}