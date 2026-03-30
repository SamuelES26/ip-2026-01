package main

import "fmt"

func main() {
	var num int64
	fmt.Println("Insira o número de três casas:")
	fmt.Scan(&num)
	if num <99 || num >1000{
		fmt.Println("Número Inválido")
		return
	}
	div := (num/10)%10
	fmt.Println(div)
}
