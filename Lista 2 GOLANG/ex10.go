package main

import "fmt"

func main() {
	var destino, retorno int
	var preco float64
	fmt.Println("Escolha o destino:")
	fmt.Println("1 - Região Norte")
	fmt.Println("2 - Região Nordeste")
	fmt.Println("3 - Região Centro-Oeste")
	fmt.Println("4 - Região Sul")
	fmt.Print("Digite o número do destino: ")
	fmt.Scan(&destino)

	fmt.Println("A viagem inclui retorno?")
	fmt.Println("1 - Sim (ida e volta)")
	fmt.Println("2 - Não (só ida)")
	fmt.Print("Digite sua escolha: ")
	fmt.Scan(&retorno)
	switch destino {
	case 1:
		switch retorno {
		case 1:
			preco = 900.00
		case 2:
			preco = 500.00
		default:
			fmt.Println("Opção de retorno inválida!")
			return
		}
	case 2:
		switch retorno {
		case 1:
			preco = 650.00
		case 2:
			preco = 350.00
		default:
			fmt.Println("Opção de retorno inválida!")
			return
		}
	case 3:
		switch retorno {
		case 1:
			preco = 600.00
		case 2:
			preco = 350.00
		default:
			fmt.Println("Opção de retorno inválida!")
			return
		}
	case 4:
		switch retorno {
		case 1:
			preco = 550.00
		case 2:
			preco = 300.00
		default:
			fmt.Println("Opção de retorno inválida!")
			return
		}

	}
	fmt.Printf("O preço da passagem é: R$ %.2f\n", preco)
}
