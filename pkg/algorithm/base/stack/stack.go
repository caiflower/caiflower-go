package stack

type MyStack []int

func (s *MyStack) Push(x int) {
	*s = append(*s, x)
}

func (s *MyStack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *MyStack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *MyStack) Size() int {
	return len(*s)
}
