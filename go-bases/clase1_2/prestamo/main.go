package main

import "fmt"

func main() {
	edad:= 20
	antiguedad:= 1
	sueldo:= 100000

	if edad >= 22 && antiguedad >=1 && sueldo >= 100000{
		fmt.Println("Usted puede recibir un prestamo")
	}else if edad < 22{
		fmt.Println("Usted no cumple con el requisito de edad")
	}else if antiguedad < 1{
		fmt.Println("Usted no cumple con el requisito de antiguedad")
	}else{
		fmt.Println("Usred no cumple con el requisito de sueldo")
	}


}