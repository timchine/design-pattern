package facade_pattern

import "fmt"

type CarService struct {

}

func (carService CarService) EmptyCar() {
	fmt.Println("清空购物车")
}