package leetcode

import (
	"testing"
)

//
// 请你设计并实现一个满足
// LRU (最近最少使用) 缓存 约束的数据结构。
//
//
//
// 实现
// LRUCache 类：
//
//
//
//
//
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
//key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//
//
//
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
// 示例：
//
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 10⁵
// 最多调用 2 * 10⁵ 次 get 和 put
//
//
// Related Topics 设计 哈希表 链表 双向链表 👍 3647 👎 0

func Test146(t *testing.T) {
	// 创建LRU缓存，容量为2
	cache := LRUCacheNodeConstructor(2)

	// 测试基本的put和get操作
	cache.Put(1, 1) // 缓存是 {1=1}
	cache.Put(2, 2) // 缓存是 {1=1, 2=2}

	// 测试获取存在的key
	if val := cache.Get(1); val != 1 {
		t.Errorf("期望Get(1)返回1，但得到%d", val)
	}

	// 测试LRU淘汰机制
	cache.Put(3, 3) // 该操作会使得关键字2作废，缓存是 {1=1, 3=3}

	// 测试获取已淘汰的key
	if val := cache.Get(2); val != -1 {
		t.Errorf("期望Get(2)返回-1，但得到%d", val)
	}

	// 继续测试LRU淘汰机制
	cache.Put(4, 4) // 该操作会使得关键字1作废，缓存是 {4=4, 3=3}

	// 测试获取已淘汰的key
	if val := cache.Get(1); val != -1 {
		t.Errorf("期望Get(1)返回-1，但得到%d", val)
	}

	// 测试获取存在的key
	if val := cache.Get(3); val != 3 {
		t.Errorf("期望Get(3)返回3，但得到%d", val)
	}

	if val := cache.Get(4); val != 4 {
		t.Errorf("期望Get(4)返回4，但得到%d", val)
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	head     *LRUCacheNode
	tail     *LRUCacheNode
	nodes    map[int]*LRUCacheNode
	capacity int
}

type LRUCacheNode struct {
	key   int
	value int
	prev  *LRUCacheNode
	next  *LRUCacheNode
}

func LRUCacheNodeConstructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		nodes:    make(map[int]*LRUCacheNode),
	}
}

func (this *LRUCache) Get(key int) int {
	cur := this.nodes[key]
	if cur == nil {
		return -1
	}

	this.moveToHead(cur)

	return cur.value
}

func (this *LRUCache) Put(key int, value int) {
	node := this.nodes[key]
	// 没有的节点就创建
	if node == nil {
		node = &LRUCacheNode{
			key:   key,
			value: value,
		}
	} else {
		// 这里不要忘记更新值了
		node.value = value
	}

	this.nodes[key] = node
	this.moveToHead(node)

	// 容量超出了,就把尾巴给删掉
	if len(this.nodes) > this.capacity {
		this.deleteTail()
	}
}

func (this *LRUCache) moveToHead(node *LRUCacheNode) {
	// 这一步很重要
	if node == this.head {
		return
	}

	if node == this.tail {
		this.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if this.head != nil {
		this.head.prev = node
	}

	node.prev = nil
	node.next = this.head
	this.head = node

	// 即是head也是tail，说明这个节点是第一个添加进来的节点
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) deleteTail() {
	if this.tail == nil {
		return
	}

	this.tail.prev.next = nil
	// 这里注意要用delete函数不要用，this.nodes[this.tail] = nil
	delete(this.nodes, this.tail.key)
	this.tail = this.tail.prev
	if this.tail == nil {
		this.head = nil
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
