package main

import (
	"fmt"
	
)

func main() {
	var preco float64
    var arCond, pintura, vidro, direcao int
    var precofinal float64

    fmt.Print("Digite o preço de fábrica do carro: R$ ")
    fmt.Scan(&preco)

    fmt.Print("Deseja Ar Condicionado (R$ 1750,00)? (1-Sim / 2-Não): ")
    fmt.Scan(&arCond)

    fmt.Print("Deseja Pintura Metálica (R$ 800,00)? (1-Sim / 2-Não): ")
    fmt.Scan(&pintura)

    fmt.Print("Deseja Vidro Elétrico (R$ 1200,00)? (1-Sim / 2-Não): ")
    fmt.Scan(&vidro)

    fmt.Print("Deseja Direção Hidráulica (R$ 2000,00)? (1-Sim / 2-Não): ")
    fmt.Scan(&direcao)

 
    precofinal = preco

    if arCond == 1 {
        precofinal += 1750.00
    }
    if pintura == 1 {
        precofinal += 800.00
    }
    if vidro == 1 {
        precofinal += 1200.00
    }
    if direcao == 1 {
        precofinal += 2000.00
    }

    fmt.Printf("O preço final do carro é: R$ %.2f\n", precofinal)
}