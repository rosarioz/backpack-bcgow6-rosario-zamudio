package main

import "fmt"

func main(){

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	//edad de Benjamin
	fmt.Println("la edad de Benjamin es ", employees["Benjamin"])

	//empleados mayores de 21
	olders := 0

	for key, element := range employees{
		if element > 21{
			olders ++
			fmt.Println(key, " es mayor de 21 años")
		}
	}
	fmt.Println("la cantidad de empleados mayores de 21 años es ", olders)

	//agregar a federico: 25
	employees["Federico"] = 25

	//eliminar a Pedro
	delete(employees, "Pedro")


}