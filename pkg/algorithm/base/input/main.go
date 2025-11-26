package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	str := s.Text()
	fields := strings.Fields(str)
	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])

	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		str = s.Text()
		fields = strings.Fields(str)
		t := make([]int, m)

		for j := 0; j < m; j++ {
			t[j], _ = strconv.Atoi(fields[j])
		}
		nums[i] = t
	}

	fmt.Println(nums)
}
