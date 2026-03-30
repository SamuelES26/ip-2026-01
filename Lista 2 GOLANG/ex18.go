package main

import (
    "fmt"
    "strings"
)

func main() {
    var preco float64
    var categoria string
    var dia int
    fmt.Print("Digite o preço normal do DVD (em R$): ")
    fmt.Scan(&preco)
    fmt.Print("Digite a categoria (comum ou lançamento): ")
    fmt.Scan(&categoria)
    fmt.Print("Digite o dia da semana (1=Domingo, 2=Segunda, ..., 7=Sábado): ")
    fmt.Scan(&dia)
    categoria = strings.ToLower(categoria)
    precofinal := preco
    if dia == 2 || dia == 3 || dia == 5 {
        precofinal = precofinal * 0.60
    }
    if categoria == "lançamento" {
        precofinal = precofinal * 1.15
    }
    fmt.Printf("Preço final da locação: R$ %.2f\n", precofinal)
}
