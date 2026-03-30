package main
	import "fmt"
func main(){
	var num int64
	fmt.Println("Insira o número:")
	fmt.Scan(&num)
	if num >= 21 && num <= 89 {
		fmt.Println("Este número está entre 20 e 90")
	} else {
		fmt.Println("Este número não está entre 20 e 90")
	}

}