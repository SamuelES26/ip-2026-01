package main

import (
	"fmt"
	"math"
)

func main() {
	var escolha int
	var raio, altura float64
	fmt.Println("1- Cone | 2- cilindro | 3- esfera")
	fmt.Scan(&escolha)
	if escolha > 3 {
		fmt.Println("Escolha Inváida")
		return
	}
	if escolha == 3 {
		fmt.Println("Digite o raio")
		fmt.Scan(&raio)
		voluesfe := (4.0/3.0) * math.Pi * math.Pow(raio, 3)
		areaesfe := 4*math.Pi*math.Pow(raio, 2)
		fmt.Printf("O volume é %.2f, e a sua área é %.2f\n", voluesfe, areaesfe)
		return
	}
	fmt.Println("Digite o raio")
	fmt.Scan(&raio)
	fmt.Println("Digite a altura")
	fmt.Scan(&altura)
	if escolha == 1 {
		volucone := math.Pi * math.Pow(raio, 2) * altura / 3
		areacone := math.Pi * raio * (math.Sqrt(raio*raio+altura*altura) + raio)
		fmt.Printf("O volume é %.2f, e a sua área é %.2f\n", volucone, areacone)
	}
	if escolha == 2 {
		volucin := math.Pi * math.Pow(raio, 2) * altura
		areacin := 2 * math.Pi* raio * (raio + altura)
		fmt.Printf("O volume é %.2f, e a sua área é %.2f\n", volucin, areacin)
	}
	
}
