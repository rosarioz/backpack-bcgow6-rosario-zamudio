package main

import (
	"fmt"
	"strings"
)

/*Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/
func main  ()  {
	fmt.Println(salaryCalcularor(5, "a"))
}

func salaryCalcularor(hs float64, category string) float64{
	var salary float64
	cat := strings.ToUpper(category)


	switch cat {
	case "C":
		salary = 1000* hs
	case "B":
		salary =1500 * hs + 1500 * hs * 0.2	
	case "A":
		salary =3000 * hs + 3000 * hs * 0.5
	
	}

	return salary
}