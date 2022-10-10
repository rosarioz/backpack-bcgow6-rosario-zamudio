package main

import "fmt"

/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/
type Student struct{
	Name string
	LastName string
	DNI string
	Date string
}

func (s Student) details() string{
	return  fmt.Sprintf("\nNombre y apellido: %s %s\nDNI: %s\nFecha: %s\n", s.Name, s.LastName, s.DNI, s.Date)
}

func main(){
	s1 := Student{Name: "Rosario", LastName: "Zamudio", DNI: "12.345.789", Date: "1/1/2022"}
	
	fmt.Println(s1.details())
	//s1.details

}



