package main

import "fmt"

type Man struct {
	Name string
	Age  int
	Tmp  *Tmp
}

type Tmp struct {
}

func main() {
	var mans []*Man
	mans = append(mans, &Man{Age: 18, Name: "lucc", Tmp: &Tmp{}})
	mans = append(mans, &Man{Age: 19, Name: "cost", Tmp: &Tmp{}})

	m := make(map[string]*Man)
	for _, v := range mans {
		// 浅拷贝
		v1 := *v
		fmt.Printf("%p", &v1)
		m[v.Name] = v
	}

	m["lucc"].Age = 100
	fmt.Printf("%v", m)
}
