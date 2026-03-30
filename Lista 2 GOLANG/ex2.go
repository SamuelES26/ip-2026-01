package main
	import "fmt"
func main(){
	fmt.Println("Insira o número")
	var num float64
	fmt.Scan(&num)
	if num == 0 {
		fmt.Println("Nulo")
	}
	if num >= 1 {
		fmt.Println("Positivo")
	} 
	if num <= -1 {
		fmt.Println("Negativo")
	}
}