package main

import (
	"errors"
	"fmt"
)

//calucualr alimento
//hacer un case para cada animal


func main()  {

	animalDog, err := animal(dog)
	if err != nil {
		fmt.Println(err)
		return
	}

	animalCat, err := animal(cat)
	if err != nil {
		fmt.Println(err)
		return
	}

	animalHamster, err := animal(hamster)
	if err != nil {
		fmt.Println(err)
		return
	}

	animalSpider, err := animal(spider)
	if err != nil {
		fmt.Println(err)
		return
	}




	amountDog := animalDog(4.0)
	amountCat := animalCat(5.0)
	amountHamster := animalHamster(7.0)
	amountSpider := animalSpider(42.0)
	

	fmt.Println("cantidad para el perro en kg ", amountDog)
	fmt.Println("cantidad para el gato en kg ", amountCat)
	fmt.Println("cantidad para el hamster en kg ", amountHamster)
	fmt.Println("cantidad para la araña en kg ", amountSpider)


	
}

const (
	dog = "dog"
	cat = "cat"
	hamster ="hamster"
	spider = "spider"
)
 
func animal (animal string) (func (cantidad float64) float64, error){
	switch animal{
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case hamster:
		return animalHamster, nil
	case spider:
		return animalSpider, nil
	default:
		msg:= fmt.Sprintf("animal no encontrado")
		return nil, errors.New(msg)

	}
}

func animalDog(cantidad float64) float64{
	return cantidad * 10.0

}

func animalCat(cantidad float64) float64{
	return cantidad * 5.0

}

func animalHamster(cantidad float64) float64{
	return cantidad * 0.250

}

func animalSpider(cantidad float64) float64{
	return cantidad * 0.150

}
