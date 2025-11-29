package leetcode

import (
	"fmt"
	"sync"
	"testing"
)

//请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
//
// 实现 LFUCache 类：
//
//
// LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
// int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
// void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量
//capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最久未使用 的键。
//
//
// 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
//
// 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
// 示例：
//
//
//输入：
//["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
//"get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
//输出：
//[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
//
//解释：
//// cnt(x) = 键 x 的使用计数
//// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
//LFUCache lfu = new LFUCache(2);
//lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
//lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
//lfu.get(1);      // 返回 1
//                 // cache=[1,2], cnt(2)=1, cnt(1)=2
//lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
//                 // cache=[3,1], cnt(3)=1, cnt(1)=2
//lfu.get(2);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3
//                 // cache=[3,1], cnt(3)=2, cnt(1)=2
//lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
//                 // cache=[4,3], cnt(4)=1, cnt(3)=2
//lfu.get(1);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3
//                 // cache=[3,4], cnt(4)=1, cnt(3)=3
//lfu.get(4);      // 返回 4
//                 // cache=[3,4], cnt(4)=2, cnt(3)=3
//
//
//
// 提示：
//
//
// 1 <= capacity <= 10⁴
// 0 <= key <= 10⁵
// 0 <= value <= 10⁹
// 最多调用 2 * 10⁵ 次 get 和 put 方法
//
//
// Related Topics 设计 哈希表 链表 双向链表 👍 894 👎 0

func Test460(t *testing.T) {
}

// leetcode submit region begin(Prohibit modification and deletion)
type LinkedHashMap struct {
	itemMap map[string]*linkedHashMapNode
	head    *linkedHashMapNode // 头结点，插入顺序第一个
	tail    *linkedHashMapNode // 尾结点，插入顺序最后一个
}

func NewLinkHashMap() *LinkedHashMap {
	return &LinkedHashMap{
		itemMap: make(map[string]*linkedHashMapNode),
	}
}

type linkedHashMapNode struct {
	key   string
	value interface{}
	prev  *linkedHashMapNode
	next  *linkedHashMapNode
}

func (m *LinkedHashMap) Put(k string, v interface{}) {
	if n, ok := m.itemMap[k]; ok {
		n.value = v
		return
	}
	n := &linkedHashMapNode{
		key:   k,
		value: v,
	}
	m.itemMap[k] = n
	if m.tail == nil {
		m.head = n
		m.tail = n
	} else {
		m.tail.next = n
		n.prev = m.tail
		m.tail = n
	}
}

func (m *LinkedHashMap) Get(k string) (interface{}, bool) {
	if n, ok := m.itemMap[k]; ok {
		return n.value, true
	}
	return nil, false
}

func (m *LinkedHashMap) Remove(k string) {
	n, ok := m.itemMap[k]
	if !ok {
		return
	}
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		m.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		m.tail = n.prev
	}
	delete(m.itemMap, k)
}

func (m *LinkedHashMap) RemoveFirst() string {
	if m.head == nil {
		return ""
	}
	key := m.head.key
	m.Remove(key)
	return key
}

func (m *LinkedHashMap) RemoveLast() string {
	if m.tail == nil {
		return ""
	}
	key := m.tail.key
	m.Remove(key)
	return key
}

func (m *LinkedHashMap) Size() int {
	return len(m.itemMap)
}

func (m *LinkedHashMap) Contains(key string) bool {
	_, ok := m.itemMap[key]
	return ok
}

func (m *LinkedHashMap) Keys() []string {
	var res []string
	for p := m.head; p != nil; p = p.next {
		res = append(res, p.key)
	}
	return res
}

func (m *LinkedHashMap) Values() []interface{} {
	var res []interface{}
	for p := m.head; p != nil; p = p.next {
		res = append(res, p.value)
	}
	return res
}

func (m *LinkedHashMap) MoveToHead(key string) {
	node, ok := m.itemMap[key]
	if !ok || node == m.head {
		return
	}
	m.moveToHead(node)
}

func (m *LinkedHashMap) moveToHead(node *linkedHashMapNode) {
	// 断链
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		m.tail = node.prev
	}
	// 插入到head
	node.prev = nil
	node.next = m.head
	m.head.prev = node
	m.head = node
	if m.tail == nil {
		m.tail = node
	}
}

type LFUCache struct {
	itemMap  map[string]interface{}
	freq     map[string]int
	freqMap  map[int]*LinkedHashMap
	capacity int
	minUsed  int
	lock     sync.Locker
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.Get1(fmt.Sprintf("%d", key)); ok {
		return v.(int)
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	this.Put1(fmt.Sprintf("%d", key), value)
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		freq:     make(map[string]int),
		freqMap:  make(map[int]*LinkedHashMap),
		itemMap:  make(map[string]interface{}),
		lock:     &sync.Mutex{},
	}
}

func (c *LFUCache) Put1(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.itemMap[key]; !ok {
		// 先删了，再放
		if len(c.itemMap) == c.capacity {
			deleteKey := c.freqMap[c.minUsed].RemoveLast()
			delete(c.freq, deleteKey)
			delete(c.itemMap, deleteKey)
		}
		c.minUsed = 1
	}

	c.itemMap[key] = value
	c.increase(key)
}

func (c *LFUCache) Get1(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if v, ok := c.itemMap[key]; ok {
		c.increase(key)
		return v, true
	} else {
		return nil, false
	}
}

func (c *LFUCache) increase(key string) {
	useCnt := c.freq[key]
	if useCnt != 0 {
		c.freqMap[useCnt].Remove(key)
		if useCnt == c.minUsed && c.freqMap[useCnt].Size() == 0 {
			delete(c.freqMap, useCnt)
			c.minUsed = useCnt + 1
		}
	}
	if c.freqMap[useCnt+1] == nil {
		c.freqMap[useCnt+1] = NewLinkHashMap()
	}
	c.freq[key] = useCnt + 1
	c.freqMap[useCnt+1].Put(key, struct{}{})
	c.freqMap[useCnt+1].MoveToHead(key)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
