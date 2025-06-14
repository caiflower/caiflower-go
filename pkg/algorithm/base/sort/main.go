package main

import (
	"fmt"
	"sort"
)

type item struct {
	key   string
	value int
}

func main() {
	items := make([]item, 4)
	items[0] = item{key: "test2", value: 2}
	items[1] = item{key: "test1", value: 1}
	items[2] = item{key: "test4", value: 4}
	items[3] = item{key: "test3", value: 3}

	sort.Slice(items, func(i, j int) bool {
		if items[i].value < items[j].value {
			return true
		} else {
			return false
		}
	})

	for _, v := range items {
		fmt.Printf("key = %s, value = %d\n", v.key, v.value)
	}
}
