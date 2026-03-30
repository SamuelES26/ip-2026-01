package main

import (
	"fmt"
)

func main() {
	var aluno int
	var n1, n2, n3, exercicios float64
	fmt.Println("Insira o número do aluno:")
	fmt.Scan(&aluno)
	fmt.Println("Média dos exercícios:")
	fmt.Scan(&exercicios)
	fmt.Println("Insira a 1° nota:")
	fmt.Scan(&n1)
	fmt.Println("Insira a 2° nota:")
	fmt.Scan(&n2)
	fmt.Println("Insira a 3° nota:")
	fmt.Scan(&n3)
	mediafinal := n1 + n2 + n3*3 + exercicios/7
	var conceito, situacao string
	switch {
	case mediafinal >= 90.0 && mediafinal <= 100.0:
		conceito = "A"
		situacao = "APROVADO"
	case mediafinal >= 70.5 && mediafinal < 90.0:
		conceito = "B"
		situacao = "APROVADO"
	case mediafinal >= 60.0 && mediafinal < 70.5:
		conceito = "C"
		situacao = "APROVADO"
	case mediafinal >= 40.0 && mediafinal < 60.0:
		conceito = "D"
		situacao = "REPROVADO"
	case mediafinal < 40.0:
		conceito = "E"
		situacao = "REPROVADO"
}
fmt.Println("    Boletim   ")
	fmt.Printf("Aluno: %d\n", aluno)
	fmt.Printf("Notas: %.1f, %.1f, %.1f\n", n1, n2, n3)
	fmt.Printf("Média Exercícios: %.1f\n", exercicios)
	fmt.Printf("Média de Aproveitamento: %.2f\n", mediafinal)
	fmt.Printf("Conceito: %s\n", conceito)
	fmt.Printf("Resultado: %s\n", situacao)
}

