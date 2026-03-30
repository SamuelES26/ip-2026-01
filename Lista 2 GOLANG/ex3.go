package main	
	import "fmt"
	func main() {
		var n1, n2 float64
		fmt.Println("Digite os números:")
		fmt.Scan(&n1, &n2)
		soma := n1 + n2
		if soma >= 20 {
				som := soma + 8
			fmt.Println("O resultado é", som)
		} else {
				sub := soma - 5
			fmt.Println("O resultado é", sub)
		}
	}