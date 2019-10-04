package caching

import "math"

/*
Design and implement a data structure for Least Frequently Used (LFU) cache.
It should support the following operations: get and put.


get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity,
it should invalidate the least frequently used item before inserting a new item. For the purpose of this problem,
when there is a tie (i.e., two or more keys that have the same frequency), the least recently used key would be
evicted.

Note that the number of times an item is used is the number of calls to the get and put functions for
that item since it was inserted. This number is set to zero when the item is removed.

Follow up:
Could you do both operations in O(1) time complexity?

The simplest method to employ an LFU algorithm is to assign a counter to every block
that is loaded into the cache. Each time a reference is made to that block the counter
is increased by one. When the cache reaches capacity and has a new block waiting to
be inserted the system will search for the block with the lowest counter and remove
it from the cache.
 */

type LFUNode struct {
	key int
	value interface{}

	prev, next *LFUNode
	// Used to measure current frequency counter
	freqCount int
}

func NewLFUNode(key int, value interface{}) *LFUNode  {
	return &LFUNode{
		key:       key,
		value:     value,
		prev:      nil,
		next:      nil,
		freqCount: 1,
	}
}

type LFUList struct {
	head, tail *LFUNode
	size int
}

func NewLFUList() *LFUList  {
	node := &LFUList{
		head: NewLFUNode(math.MaxInt32, nil),
		tail: NewLFUNode(math.MinInt32, nil),
		size: 0,
	}
	// Connect head next with tail
	node.head.next = node.tail

	// Connect tail prev with head
	node.tail.prev = node.head
	return node
}

// Insert element at head. Note we keep "head" and "tail" nodes intact
func (list *LFUList)InsertHead(node *LFUNode)  {
	// Save current head to node.next
	node.next = list.head

	// Update current head next node to point to current node as prev
	list.head.next.prev = node

	// Update current head next node to point to current node
	list.head.next = node

	// Attach current node prev to current head
	node.prev = list.head
	list.size += 1
}

// Remove element at tail
func (list *LFUList)RemoveAtTail() *LFUNode  {
	// Save current tail prev node
	old := list.tail.prev
	prev := list.tail.prev

	prev.prev.next = list.tail
	list.tail.prev = prev.prev
	list.size -= 1
	return old
}

// Remove node
func (list *LFUList)RemoveNode(node *LFUNode) {
	// point previous node to next
	node.prev.next = node.next
	// point next node to prev
	node.next.prev = node.prev
	list.size -= 1
}

type LFUCache struct {
	keys map[int]*LFUNode
	freq map[int]*LFUList
	capacity int
	minFreq int
	size int
}

/*
 Creates a new LFU cache with capacity.
 */
func NewLFUCache(cap int) *LFUCache  {
	return &LFUCache{
		keys:     make(map[int]*LFUNode),
		freq:      make(map[int]*LFUList),
		capacity: cap,
		minFreq:  0,
		size:     0,
	}
}

/*
Implementing set for the LFU has a few steps. There are two cases: insert the new
item and replace an old item. When inserting a new item, a new node is created. If the
cache is not full, it can be inserted into the freq’s doubly linked list of frequency 1. If the
capacity is full, the tail item in the doubly linked list of frequency is deleted, and then the
new node is inserted.

If the element already exists and needs to be replaced, the node is brought to the
head of its corresponding frequency doubly linked list. Finally, the minimum frequency
variable, minFreq, is incremented accordingly to compute which item should be evicted
in the future.
 */
func (cache *LFUCache)Put(key int, value int) {
	node := cache.keys[key]
	// Cache miss
	if node == nil {
		node  = NewLFUNode(key, value)
		cache.keys[key] = node
		if cache.size != cache.capacity {
			// insert without deleting
			// Create a new frequency list for the new count if not exists
			if cache.freq[1] == nil {
				cache.freq[1] = NewLFUList()
			}
			// Add the node to the new freq map
			cache.freq[1].InsertHead(node)

			// update size
			cache.size += 1
		} else {
			// Capacity reached
			// delete from minFreq and insert
			oldTail := cache.freq[cache.minFreq].RemoveAtTail()
			cache.keys[oldTail.key] = nil
			// Create a new frequency list for the new count if not exists
			if cache.freq[1] == nil {
				cache.freq[1] = NewLFUList()
			}
			// Add the node to the new freq map
			cache.freq[1].InsertHead(node)
		}
		// Reset minFreq for the new element
		cache.minFreq = 1
	} else {
		// Cache hit. Increase freq counter of the current element and update minFreq
		oldFreq := node.freqCount
		node.value = value
		node.freqCount = node.freqCount + 1

		// Remove node from oldFreq list
		cache.freq[oldFreq].RemoveNode(node)

		// Create a new frequency list for the new count if not exists
		if cache.freq[node.freqCount] == nil {
			cache.freq[node.freqCount] = NewLFUList()
		}
		// Add the node to the new freq map
		cache.freq[node.freqCount].InsertHead(node)

		// If the oldFreq cache is empty update minFreq
		if oldFreq == cache.minFreq && cache.freq[oldFreq].size == 0 {
			cache.minFreq += 1
		}
	}
}

/*
To implement get, the cache needs to return existing nodes in O(1) time and
increment the counter for accessing. If the element does not exist in the cache, it is
forced to return a null element. Otherwise, the frequency for the element is increased,
the item is brought to the head of the doubly linked list, and the minimum frequency
variable, minFreq, is adjusted accordingly.
 */
func (cache *LFUCache)Get(key int) int {
	node := cache.keys[key]

	// Cache miss
	if node == nil {
		return -1
	}
	oldFreq := node.freqCount
	node.freqCount = node.freqCount + 1

	// Remove node from old frequency list
	cache.freq[oldFreq].RemoveNode(node)

	// Create a new frequency list for the new count if not exists
	if cache.freq[node.freqCount] == nil {
		cache.freq[node.freqCount] = NewLFUList()
	}
	// Add the node to the new freq map
	cache.freq[node.freqCount].InsertHead(node)

	// if the old freq is the last one then update the minFreq
	if oldFreq == cache.minFreq && cache.freq[oldFreq].size == 0 {
		cache.minFreq += 1
	}
	return node.value.(int)
}
