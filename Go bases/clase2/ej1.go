package main

import "fmt"

func main(){


fmt.Println(salaryTaxes(300000))

}

func salaryTaxes(salary float32) float32{

	var tax float32
	if salary > 50000{
		tax = salary * 0.17
	}else if salary > 150000{
		tax = salary * 0.27
	}

	return tax

}
