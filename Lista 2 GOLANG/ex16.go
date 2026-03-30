package main

import ("fmt"
		
	)

func main() {
	var A, B, C float64
	fmt.Println("Digite os coeficiente A:")
	fmt.Scan(&A)
	fmt.Println("Digite os coeficiente B:")
	fmt.Scan(&B)
	fmt.Println("Digite os coeficiente C:")
	fmt.Scan(&C)
	if A == 0 {
		fmt.Println("Isso não é uma equação de segundo grau")
	}
	delta := B*B - 4*A*C
	if delta < 0 {
		fmt.Printf("O resultado é %2.f e essa equação tem raiz imaginária", delta)
	}
	if delta == 0 {
		fmt.Printf("O resultado é %2.f e essa equação tem raiz unica", delta)
	}
	if delta > 0 {
		fmt.Printf("O resultado é %2.f e essa equação tem raiz distintas\n", delta)
	}

}
