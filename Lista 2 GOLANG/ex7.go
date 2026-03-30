package main

import "fmt"


func main() {
	var A, B, C int
	var maior, menor, inter int
	fmt.Println("Númrero A:")
	fmt.Scan(&A)
	fmt.Println("Númrero B:")
	fmt.Scan(&B)
	fmt.Println("Númrero C:")
	fmt.Scan(&C)

	if A < B && A < C {
        menor = A
        if B < C {
            inter = B
            maior = C
        } else {
            inter = C
            maior = B
        }
    } else if B < A && B < C {
        menor = B
        if A < C {
            inter = A
            maior = C
        } else {
            inter = C
            maior = A
        }
    } else {
        menor = C
        if A < B {
            inter = A
            maior = B
        } else {
            inter = B
            maior = A
        }
    }

    fmt.Println("Valores em ordem crescente:")
    fmt.Println(menor, inter, maior)
}