package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	initValue := "init-value"
	addValue := "add-value"

	set := mapset.NewSet[string](initValue)
	set.Add(addValue)
	fmt.Printf("contains "+initValue+" %v \n", set.Contains(initValue))
	fmt.Printf("contains "+addValue+" %v \n", set.Contains(addValue))

	slice := set.ToSlice()
	fmt.Printf("set to slice value %v \n", slice)

	s := set.String()
	fmt.Printf("set to str value %v \n", s)

	set1 := mapset.NewSet[string](initValue)
	difference := set.Difference(set1)
	fmt.Printf("set dif set1 %v \n", difference.String())
}
