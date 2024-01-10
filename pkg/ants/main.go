package main

import (
	"fmt"
	"sync"

	"github.com/panjf2000/ants/v2"
)

/**
优点： 池的概念
缺点： 函数变量传入不好处理，需要使用到断言操作
*/

func main() {
	pool, err := ants.NewPool(10)
	defer pool.Release()
	if err != nil {
		panic(err)
	}

	wait := sync.WaitGroup{}
	for i := 0; i <= 100; i++ {
		wait.Add(1)
		err := pool.Submit(func() {
			fmt.Printf("yes \n")
			wait.Done()
		})
		if err != nil {
			return
		}
	}

	wait.Wait()

	fn := func(i int) {
		fmt.Printf("i = %v \n", i)
	}

	for i := 0; i <= 100; i++ {
		wait.Add(1)
		withFunc, err := ants.NewPoolWithFunc(10, func(fn interface{}) {
			f := fn.(func(i int))
			f(i)
			wait.Done()
		})
		if err != nil {
			panic(err)
		}
		withFunc.Invoke(fn)
	}

	wait.Wait()
}
