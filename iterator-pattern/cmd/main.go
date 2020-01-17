package main

import (
	"fmt"
	iterator_pattern "github.com/echoloveyou/design-patterns/iterator-pattern"
)

func main() {
	var (
		iterator iterator_pattern.Menu
	)
	iterator = new(iterator_pattern.PancakeHouseMenu)
	i := iterator.CreateIterator()
	for i.HasNext() {
		fmt.Println(i.Next())
	}
	iterator = iterator_pattern.DinerMenu{}
	i = iterator.CreateIterator()
	for i.HasNext() {
		fmt.Println(i.Next())
	}
}
