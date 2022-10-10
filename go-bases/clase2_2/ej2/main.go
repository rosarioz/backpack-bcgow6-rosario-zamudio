package main

import (
	"fmt"

	
)

/*Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct{
	vals []float64
	high float64
	width float64
	cuadratic bool
	max float64
}

func (m Matrix) Set(vals []float64) {
	m.vals = vals
}

func (m Matrix) Print() {
	for i := 0; i < int (m.high); i++ {
		fmt.Printf("\t%.0f\n", m.vals[i*int(m.width):i*int(m.width)+int(m.width)])
	}
}




