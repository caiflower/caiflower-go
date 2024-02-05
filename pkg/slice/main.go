package main

import "fmt"

type Man struct {
	Age int
}

func main() {
	test2()
}

func test1() {
	var manArray [2]Man
	manArray[0] = Man{Age: 10}
	manArray[1] = Man{Age: 10}

	nums := []int{1, 2, 3}
	mans := []Man{{Age: 10}, {Age: 20}}
	fmt.Printf("%p, %v\n", nums, nums)
	fmt.Printf("%v\n", manArray)
	fmt.Printf("%p, %v\n", mans, mans)

	changeSlice(nums)
	changeManArray(manArray)
	changeSliceAge(mans)

	fmt.Printf("%p, %v\n", nums, nums)
	fmt.Printf("%v\n", manArray)
	fmt.Printf("%p, %v\n", mans, mans)
}

func changeSlice(slices []int) {
	slices[2] = 5
}

func changeSliceAge(mans []Man) {
	mans[1].Age = 30
}

func changeManArray(mans [2]Man) {
	mans[1].Age = 30
}

func test2() {
	// 验证slice[xx:xx]是和原来的切片共用一片内存地址的
	slice := []int{1, 2, 3, 4}
	slice1 := slice[0 : len(slice)-1]
	slice1[0] = 5

	fmt.Printf("%v\n", slice)
	fmt.Printf("%v\n", slice1)
}

//func test(namespace string, arr []string) bool {
//	return slices
//}
