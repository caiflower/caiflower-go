package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"strconv"
	"sync"
)

/**
优点： 池的概念
缺点： 函数变量传入不好处理，需要使用到断言操作或者方法
*/

func main() {
	// testPoolWithFunc()
	testPool()
}

func testPoolWithFunc() {
	wait := sync.WaitGroup{}

	fn := func(i interface{}) {
		fmt.Printf("i = %v \n", i)
		wait.Done()
	}

	withFunc, err := ants.NewPoolWithFunc(10, fn)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 20; i++ {
		wait.Add(1)
		err = withFunc.Invoke(i)
		if err != nil {
			continue
		}

	}

	wait.Wait()
	withFunc.Release()
}

type testPoolSt struct {
	Name string
	wait *sync.WaitGroup
}

func (t *testPoolSt) print() {
	fmt.Printf("---%s---\n", t.Name)
	t.wait.Done()
}

func testPool() {
	pool, err := ants.NewPool(10)
	if err != nil {
		panic(err)
	}

	wait := sync.WaitGroup{}
	wait.Add(10)

	for i := 1; i <= 10; i++ {
		t := testPoolSt{"test" + strconv.Itoa(i), &wait}
		err := pool.Submit(t.print)
		if err != nil {
			continue
		}
	}

	wait.Wait()
	pool.Release()
}
