package main	
	import "fmt"
	func main(){
		var num int64
		fmt.Println("Insira o número:")
		fmt.Scan(&num)
		if num % 5 == 0 {
			fmt.Println("Esse número e divisível por 5.")
		} else {
			fmt.Println("Esse número não é divisível por 5.")
		}
	}