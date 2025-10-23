package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//
// 实现 MyQueue 类：
//
//
// void push(int x) 将元素 x 推到队列的末尾
// int pop() 从队列的开头移除并返回元素
// int peek() 返回队列开头的元素
// boolean empty() 如果队列为空，返回 true ；否则，返回 false
//
//
// 说明：
//
//
// 你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法
//的。
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//
//
//
//
// 示例 1：
//
//
//输入：
//["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
//输出：
//[null, null, null, 1, 1, false]
//
//解释：
//MyQueue myQueue = new MyQueue();
//myQueue.push(1); // queue is: [1]
//myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//myQueue.peek(); // return 1
//myQueue.pop(); // return 1, queue is [2]
//myQueue.empty(); // return false
//
//
//
//
//
//
//
// 提示：
//
//
// 1 <= x <= 9
// 最多调用 100 次 push、pop、peek 和 empty
// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）
//
//
//
//
// 进阶：
//
//
// 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
//
//
// Related Topics 栈 设计 队列 👍 1249 👎 0

func Test232(t *testing.T) {
	// 假设有如下接口
	// type MyQueue struct{...}
	// func Constructor() MyQueue
	// func (this *MyQueue) Push(x int)
	// func (this *MyQueue) Pop() int
	// func (this *MyQueue) Peek() int
	// func (this *MyQueue) Empty() bool

	// 示例1
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	assert.Equal(t, 1, queue.Peek(), "Peek() should return 1")
	assert.Equal(t, 1, queue.Pop(), "Pop() should return 1")
	assert.False(t, queue.Empty(), "Empty() should return false")

	// 边界情况：只剩一个元素
	queue.Push(3)
	assert.Equal(t, 2, queue.Peek(), "Peek() should return 2")
	assert.Equal(t, 2, queue.Pop(), "Pop() should return 2")
	assert.False(t, queue.Empty(), "Empty() should return false")
	assert.Equal(t, 3, queue.Pop(), "Pop() should return 3")
	assert.True(t, queue.Empty(), "Empty() should return true")

	// 连续 push 和 pop
	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	assert.Equal(t, 4, queue.Pop(), "Pop() should return 4")
	assert.Equal(t, 5, queue.Pop(), "Pop() should return 5")
	assert.Equal(t, 6, queue.Peek(), "Peek() should return 6")
	assert.Equal(t, 6, queue.Pop(), "Pop() should return 6")
	assert.True(t, queue.Empty(), "Empty() should return true")
}

// leetcode submit region begin(Prohibit modification and deletion)
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

func (s *MyStack) IsEmpty() bool {
	return s.Size() == 0
}

type MyQueue struct {
	s1 MyStack
	s2 MyStack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}

	return this.s2.Pop()
}

func (this *MyQueue) Peek() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
