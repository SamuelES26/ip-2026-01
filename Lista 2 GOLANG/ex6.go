package main 
	import "fmt"
	func main() {
		var n1, n2 int64
		fmt.Println("Insira os números:")
		fmt.Scan(&n1, &n2)
		if n1 % n2 == 0 {
			fmt.Println("Essea números são divisíveis.")
		} else {
			fmt.Println("Esses números não são divisíveis.")
		}
	}
	