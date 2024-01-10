package main

import "fmt"

/**
1281. 整数的各位积和之差

给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。
*/

func main() {
	fmt.Printf("%v\n", subtractProductAndSum(4421))
}

func subtractProductAndSum(n int) int {
	var n1, n2 int = 0, 1
	for n != 0 {
		n1 += n % 10
		n2 *= n % 10
		n /= 10
	}
	return n2 - n1
}
