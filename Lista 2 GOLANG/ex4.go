package main
import (
	"fmt"
	"math"
)
	func main(){
		var x float64
		fmt.Println("Insira o número:")
		fmt.Scan(&x)
		if x >= 0 {
			resultado := math.Sqrt(x)
		fmt.Printf("A raiz quadrada de %.2f é %.2f\n", x, resultado)
		}
		if x <= -1 {
			quad := x * x
			fmt.Printf("O quadrado de %.2f é %.2f\n", x, quad)
		}
	

	}