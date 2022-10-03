package main

import(
	"fmt"
	"errors"
)
/*vals min, max y promedio
func: indicar el tipo de calculo
y que devuelva otra funcion y un mensaje 
*/
func main()  {

	
	 
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}


	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}


	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("valor promedio: ", minValue)
	fmt.Println("valor minimo: ", averageValue)
	fmt.Println("valor maximo: ", maxValue)


}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation (op string)(func(values ... float64) float64, error){
	switch op{
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		msg := fmt.Sprintf("no existe esa operacion")
		return nil, errors.New(msg)
	}
	
}

func minFunc(values ... float64) float64{
	var min float64
	min = values[0]
	for _, value := range values {
		if value < min{
			min = value
		}
	}
	return min
}

func maxFunc(values ... float64) float64{
	var max float64
	max = values[0]
	for _, value := range values {
		if value > max{
			max = value
		}
	}
	return max
}

func averageFunc(values ... float64) float64{
	sumValues := 0.0
	cantValues:= 0.0

	for _, value := range values {
		sumValues += value
		cantValues ++
	}

	return sumValues/cantValues

}