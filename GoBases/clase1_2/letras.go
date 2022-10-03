package main

import "fmt"
import "strings"

func main() {
	palabra := "hola"

	fmt.Println("la palabra", palabra, "tiene: ", len(palabra), " letras")


	letras := strings.Split(palabra, "")


	for _, letra:= range letras{
		fmt.Println(letra)
	}





	//fmt.Println(letras)







}