package main

import "fmt"

/* Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. 
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio 
y un error en caso que uno de los números ingresados sea negativo*/

func main(){

	fmt.Println(avgScores(5,6,2,8,9,6,))

}

func avgScores(values ... float64) float64{
	sumScores := 0.0
	cantScores:= 0.0

	for _, value := range values {
		sumScores += value
		cantScores ++
	}

	return sumScores/cantScores

}