package _25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
//
// 实现 MyStack 类：
//
//
// void push(int x) 将元素 x 压入栈顶。
// int pop() 移除并返回栈顶元素。
// int top() 返回栈顶元素。
// boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
//
//
//
//
// 注意：
//
//
// 你只能使用队列的标准操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
// 你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
//
//
//
//
// 示例：
//
//
//输入：
//["MyStack", "push", "push", "top", "pop", "empty"]
//[[], [1], [2], [], [], []]
//输出：
//[null, null, null, 2, 2, false]
//
//解释：
//MyStack myStack = new MyStack();
//myStack.push(1);
//myStack.push(2);
//myStack.top(); // 返回 2
//myStack.pop(); // 返回 2
//myStack.empty(); // 返回 False
//
//
//
//
// 提示：
//
//
// 1 <= x <= 9
// 最多调用100 次 push、pop、top 和 empty
// 每次调用 pop 和 top 都保证栈不为空
//
//
//
//
// 进阶：你能否仅用一个队列来实现栈。
//
// Related Topics 栈 设计 队列 👍 973 👎 0

func Test225(t *testing.T) {
	stack := Constructor() // 假设你的构造函数是 Constructor()

	// 测试 push
	stack.Push(1)
	stack.Push(2)

	// 测试 top
	top := stack.Top()
	assert.Equal(t, 2, top, "Top should return 2")

	// 测试 pop
	pop := stack.Pop()
	assert.Equal(t, 2, pop, "Pop should return 2")

	// 测试 empty
	empty := stack.Empty()
	assert.False(t, empty, "Stack should not be empty")

	// 再 pop 一次
	pop = stack.Pop()
	assert.Equal(t, 1, pop, "Pop should return 1")

	// 现在应该空了
	empty = stack.Empty()
	assert.True(t, empty, "Stack should be empty")

	// 再压入多个元素
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	assert.Equal(t, 5, stack.Top(), "Top should return 5")
	assert.Equal(t, 5, stack.Pop(), "Pop should return 5")
	assert.Equal(t, 4, stack.Pop(), "Pop should return 4")
	assert.Equal(t, 3, stack.Pop(), "Pop should return 3")
	assert.True(t, stack.Empty(), "Stack should be empty after popping all")
}

// leetcode submit region begin(Prohibit modification and deletion)
type MyQueue []int

func (q *MyQueue) Push(x int) {
	*q = append(*q, x)
}

func (q *MyQueue) Pop() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *MyQueue) Peek() int {
	return (*q)[0]
}

func (q *MyQueue) Empty() bool {
	return len(*q) == 0
}

func (q *MyQueue) Size() int {
	return len(*q)
}

type MyStack struct {
	q1 MyQueue
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q1.Push(x)
}

func (this *MyStack) Pop() int {
	l := this.q1.Size() - 1
	for i := 0; i < l; i++ {
		this.q1.Push(this.q1.Pop())
	}
	return this.q1.Pop()
}

func (this *MyStack) Top() int {
	l := this.q1.Size()
	val := 0
	for i := 0; i < l; i++ {
		val = this.q1.Pop()
		this.q1.Push(val)
	}
	return val
}

func (this *MyStack) Empty() bool {
	return this.q1.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
