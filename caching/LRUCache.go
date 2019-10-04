package caching

import "math"

/*
LRU Implementation

When an item
in the cache is accessed, that item moves to the back (newest in the order) of the list.
When a page not found in the cache is accessed, the front item (or oldest in the order) is
removed, and the new item is put at the back (newest in the order) of the list.
 */

type LRUNode struct {
	key int
	value interface{}

	prev, next *LRUNode
}

func NewLRUNode(key int, value interface{}) *LRUNode  {
	return &LRUNode{
		key:       key,
		value:     value,
		prev:      nil,
		next:      nil,
	}
}

type LRUCache struct {
	keys map[int]*LRUNode
	head *LRUNode
	tail *LRUNode
	capacity int
}

/*
 Creates a new LRU cache with capacity.
*/
func NewLRUCache(cap int) *LRUCache  {
	node := &LRUCache{
		keys: make(map[int]*LRUNode),
		head: NewLRUNode(math.MaxInt32, nil),
		tail: NewLRUNode(math.MinInt32, nil),
		capacity: cap,
	}
	// Connect head next with tail
	node.head.next = node.tail

	// Connect tail prev with head
	node.tail.prev = node.head
	return node
}

func (c *LRUCache)RemoveNode(node *LRUNode)  {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (c *LRUCache)AddNode(node *LRUNode)  {
	tail := c.tail.prev
	tail.next = node

	c.tail.prev = node
	node.prev = tail
	node.next = c.tail
}

/*
Whenever get is called, the
LRU caching scheme brings that node to the head of the doubly linked list since it was
the most recently used node. This is the same as deleting and adding the node
 */
func (c *LRUCache)Get(key int) int  {
	node := c.keys[key]
	if node == nil {
		return -1
	}
	// Detach node from list
	c.RemoveNode(node)

	// Reattach at head
	c.AddNode(node)

	return node.value.(int)
}

func (c *LRUCache)Set(key int, value int)  {
	node := c.keys[key]
	// Cache hit
	if node != nil {
		// Perform update
		c.RemoveNode(node)
	}
	// Cache miss
	newNode := NewLRUNode(key, value)
	c.AddNode(newNode)
	c.keys[key] = newNode

	// evict a node
	if len(c.keys) > c.capacity {
		// grab current head and remove
		head := c.head.next
		c.RemoveNode(head)
		// clear keys table
		delete(c.keys, head.key)
	}
}

