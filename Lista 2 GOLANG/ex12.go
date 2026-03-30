package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Insira sua idade:")
	fmt.Scan(&idade)
	if idade <= 0 {
		fmt.Println("Idade Inválida")
	}
	if idade <= 2 {
		fmt.Println("Recém-Nascido")
	}
	if idade >= 3 && idade <= 11 {
		fmt.Println("Criança")
	}
	if idade >= 12 && idade <=19 {
		fmt.Println("Adolescente")
	}
	if idade >= 20 && idade <= 55 {
		fmt.Println("Adulto")
	}
	if idade > 55 {
		fmt.Println("Idoso")
	}






}