package main

import (
	"fmt"
	"strconv"
)

type MyInterface interface {
	toString(a int) string
	toFloat64(b int) float64
}

type common struct {
}

func (t *common) toString(a int) string {
	return strconv.Itoa(a)
}

type test1 struct {
	common
}

func (t *test1) toFloat64(a int) float64 {
	return float64(a)
}

type test2 struct {
	common
}

func (t *test2) toFloat64(a int) float64 {
	return float64(a) + 1
}

func main() {
	t1 := test1{}
	t2 := test2{}
	val := 1

	// toString都继承了common的实现
	fmt.Printf("toString := %s\n", t1.toString(val))
	fmt.Printf("toString := %s\n", t2.toString(val))

	// 各自实现ToFloat64
	fmt.Printf("toFloat64 := %v\n", t1.toFloat64(val))
	fmt.Printf("toFloat64 := %v\n", t2.toFloat64(val))
}
