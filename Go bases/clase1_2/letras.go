package main

import "fmt"
import "strings"

func main() {
	palabra := "hola"

	fmt.Println("la palabra", palabra, "tiene: ", len(palabra), " letras")


	letras := strings.Split(palabra, "")


	for i, letra:= range letras{
		fmt.Println(i, letra)
	}





	//fmt.Println(letras)







}