// package cache

// import (
// 	"container/list"
// 	"sync"
// 	"time"
// )

// // CacheItem represents an item in the cache. It includes a key, a value, and an expiration timer.
// type CacheItem struct {
// 	key        string
// 	value      interface{}
// 	expiration *time.Timer
// }

// // LRUCache represents a Least Recently Used (LRU) cache. It includes a capacity, a map for the cache items,
// // a doubly linked list to keep track of the order of items, and a mutex for concurrency control.
// type LRUCache struct {
// 	capacity int
// 	cache    map[string]*list.Element
// 	items    *list.List
// 	mu       sync.Mutex
// }

// // NewLRUCache creates a new LRUCache with the specified capacity.
// func NewLRUCache(capacity int) *LRUCache {
// 	return &LRUCache{
// 		capacity: capacity,
// 		cache:    make(map[string]*list.Element),
// 		items:    list.New(),
// 	}
// }

// // Set adds a key-value pair to the cache. If the key already exists, it updates the value and moves the key to the front of the cache.
// // If the cache is at capacity, it removes the least recently used item. It also sets an expiration timer if a duration is provided.
// func (c *LRUCache) Set(key string, value interface{}, duration time.Duration) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	if element, ok := c.cache[key]; ok {
// 		c.items.Remove(element)
// 		delete(c.cache, key)
// 	}

// 	item := &CacheItem{
// 		key:   key,
// 		value: value,
// 	}
// 	element := c.items.PushFront(item)
// 	c.cache[key] = element

// 	if duration > 0 {
// 		item.expiration = time.AfterFunc(duration, func() {
// 			c.Delete(key)
// 		})
// 	}

// 	if c.items.Len() > c.capacity {
// 		element := c.items.Back()
// 		if element != nil {
// 			item := element.Value.(*CacheItem)
// 			delete(c.cache, item.key)
// 			c.items.Remove(element)
// 		}
// 	}
// }

// // Get retrieves a value from the cache using a key. If the key exists, it moves the key to the front of the cache and returns the value.
// func (c *LRUCache) Get(key string) (interface{}, bool) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	if element, ok := c.cache[key]; ok {
// 		c.items.MoveToFront(element)
// 		return element.Value.(*CacheItem).value, true
// 	}
// 	return nil, false
// }

// // Delete removes a key-value pair from the cache. If the key exists, it stops the expiration timer, removes the key from the cache map,
// // and removes the item from the doubly linked list.
// func (c *LRUCache) Delete(key string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	if element, ok := c.cache[key]; ok {
// 		item := element.Value.(*CacheItem)
// 		if item.expiration != nil {
// 			item.expiration.Stop()
// 		}
// 		delete(c.cache, key)
// 		c.items.Remove(element)
// 	}
// }

// cache.go
// cache.go
// package cache

// import (
// 	"container/list"
// 	"sync"
// 	"time"
// )

// // CacheItem represents an item in the cache. It includes a key, a value, and an expiration timer.
// type CacheItem struct {
// 	key        string
// 	value      interface{}
// 	expiration *time.Timer
// }

// // LRUCache represents a Least Recently Used (LRU) cache. It includes a capacity, a map for the cache items,
// // a doubly linked list to keep track of the order of items, and a mutex for concurrency control.
// type LRUCache struct {
// 	capacity int
// 	cache    map[string]*list.Element
// 	items    *list.List
// 	mu       sync.Mutex
// }

// // NewLRUCache creates a new LRUCache with the specified capacity.
// func NewLRUCache(capacity int) *LRUCache {
// 	return &LRUCache{
// 		capacity: capacity,
// 		cache:    make(map[string]*list.Element),
// 		items:    list.New(),
// 	}
// }

// // Set adds a key-value pair to the cache. If the key already exists, it updates the value and moves the key to the front of the cache.
// // If the cache is at capacity, it removes the least recently used item. It also sets an expiration timer if a duration is provided.
// func (c *LRUCache) Set(key string, value interface{}, duration time.Duration) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	if element, ok := c.cache[key]; ok {
// 		c.items.Remove(element)
// 		delete(c.cache, key)
// 	}

// 	item := &CacheItem{
// 		key:   key,
// 		value: value,
// 	}
// 	element := c.items.PushFront(item)
// 	c.cache[key] = element

// 	if duration > 0 {
// 		item.expiration = time.AfterFunc(duration, func() {
// 			c.Delete(key)
// 		})
// 	}

// 	if c.items.Len() > c.capacity {
// 		element := c.items.Back()
// 		if element != nil {
// 			item := element.Value.(*CacheItem)
// 			delete(c.cache, item.key)
// 			c.items.Remove(element)
// 		}
// 	}
// }

// // Get retrieves a value from the cache using a key. If the key exists, it moves the key to the front of the cache and returns the value.
// func (c *LRUCache) Get(key string) (interface{}, bool) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	if element, ok := c.cache[key]; ok {
// 		c.items.MoveToFront(element)
// 		return element.Value.(*CacheItem).value, true
// 	}
// 	return nil, false
// }

// // Delete removes a key-value pair from the cache. If the key exists, it stops the expiration timer, removes the key from the cache map,
// // and removes the item from the doubly linked list.
// func (c *LRUCache) Delete(key string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	if element, ok := c.cache[key]; ok {
// 		item := element.Value.(*CacheItem)
// 		if item.expiration != nil {
// 			item.expiration.Stop()
// 		}
// 		delete(c.cache, key)
// 		c.items.Remove(element)
// 	}
// }

package cache

import (
	"sync"
	"time"
)

// DLLNode represents a node in the Doubly Linked List (DLL)
type DLLNode struct {
	key        string
	value      interface{}
	prev, next *DLLNode
	expiration *time.Time // New field for expiration time
}

// DoublyLinkedList represents a Doubly Linked List (DLL)
type DoublyLinkedList struct {
	head, tail *DLLNode
}

// LRUCache represents a Least Recently Used (LRU) cache with expiration.
type LRUCache struct {
	capacity int
	cache    map[string]*DLLNode
	list     *DoublyLinkedList
	mu       sync.Mutex
}

// NewLRUCache creates a new LRUCache with the specified capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*DLLNode),
		list:     &DoublyLinkedList{},
	}
}

// Set adds a key-value pair to the cache with expiration time.
func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.cache[key]; ok {
		node.value = value
		node.expiration = expirationFromNow(expiration)
		c.list.moveToFront(node)
	} else {
		if len(c.cache) >= c.capacity {
			c.evictLRU()
		}
		newNode := &DLLNode{
			key:        key,
			value:      value,
			expiration: expirationFromNow(expiration),
		}
		c.cache[key] = newNode
		c.list.addToFront(newNode)
	}
}

// Get retrieves a value from the cache using a key.
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.cache[key]; ok {
		if node.expiration != nil && node.expiration.Before(time.Now()) {
			// Expired, remove from cache
			c.deleteNode(node)
			return nil, false
		}
		c.list.moveToFront(node)
		return node.value, true
	}
	return nil, false
}

// Delete removes a key-value pair from the cache.
func (c *LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.cache[key]; ok {
		c.deleteNode(node)
	}
}

// deleteNode removes a node from the cache and the doubly linked list.
func (c *LRUCache) deleteNode(node *DLLNode) {
	delete(c.cache, node.key)
	c.list.remove(node)
}

// EvictLRU removes the least recently used item from the cache.
func (c *LRUCache) evictLRU() {
	if c.list.tail != nil {
		c.deleteNode(c.list.tail)
	}
}

// addToFront adds a node to the front of the doubly linked list.
func (dll *DoublyLinkedList) addToFront(node *DLLNode) {
	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		node.next = dll.head
		dll.head.prev = node
		dll.head = node
	}
}

// moveToFront moves a node to the front of the doubly linked list.
func (dll *DoublyLinkedList) moveToFront(node *DLLNode) {
	if dll.head == node {
		return
	}

	if node == dll.tail {
		dll.tail = node.prev
		dll.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	node.prev = nil
	node.next = dll.head
	dll.head.prev = node
	dll.head = node
}

// remove removes a node from the doubly linked list.
func (dll *DoublyLinkedList) remove(node *DLLNode) {
	if node == nil {
		return
	}

	if node == dll.head {
		dll.head = node.next
		if dll.head != nil {
			dll.head.prev = nil
		} else {
			dll.tail = nil
		}
	} else if node == dll.tail {
		dll.tail = node.prev
		if dll.tail != nil {
			dll.tail.next = nil
		} else {
			dll.head = nil
		}
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}

// expirationFromNow calculates the expiration time from the current time.
func expirationFromNow(d time.Duration) *time.Time {
	expiration := time.Now().Add(d)
	return &expiration
}
