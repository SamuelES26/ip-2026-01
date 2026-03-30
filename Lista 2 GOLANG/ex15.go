package main

import (
	"fmt"
	"time"
)

func main() {
	
	var dia, mes, ano int
	fmt.Println("Digite a data desejada:")
	fmt.Scan(&dia,&mes, &ano)
	dt := time.Date(ano, time.Month(mes), dia, 0, 0, 0, 0, time.Local)
	mesesPt := map[time.Month]string{
		time.January: "Janeiro", 
		time.February: "Fevereiro",
		time.March: "Março",
		time.April: "Abril",
		time.May: "Maio",
		time.June: "Junho",
		time.July: "Julho",
		time.August: "Agosto",
		time.September: "Setembro",
		time.October: "Outubro",
		time.November: "Novembro",
		time.December: "Dezembro",
		
	}
	fmt.Printf("%d de %s de %d\n", dt.Day(), mesesPt[dt.Month()], dt.Year())
}