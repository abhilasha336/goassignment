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
package cache

import (
	"container/list"
	"sync"
	"time"
)

// CacheItem represents an item in the cache. It includes a key, a value, and an expiration timer.
type CacheItem struct {
	key        string
	value      interface{}
	expiration *time.Timer
}

// LRUCache represents a Least Recently Used (LRU) cache. It includes a capacity, a map for the cache items,
// a doubly linked list to keep track of the order of items, and a mutex for concurrency control.
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	items    *list.List
	mu       sync.Mutex
}

// NewLRUCache creates a new LRUCache with the specified capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		items:    list.New(),
	}
}

// Set adds a key-value pair to the cache. If the key already exists, it updates the value and moves the key to the front of the cache.
// If the cache is at capacity, it removes the least recently used item. It also sets an expiration timer if a duration is provided.
func (c *LRUCache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.cache[key]; ok {
		c.items.Remove(element)
		delete(c.cache, key)
	}

	item := &CacheItem{
		key:   key,
		value: value,
	}
	element := c.items.PushFront(item)
	c.cache[key] = element

	if duration > 0 {
		item.expiration = time.AfterFunc(duration, func() {
			c.Delete(key)
		})
	}

	if c.items.Len() > c.capacity {
		element := c.items.Back()
		if element != nil {
			item := element.Value.(*CacheItem)
			delete(c.cache, item.key)
			c.items.Remove(element)
		}
	}
}

// Get retrieves a value from the cache using a key. If the key exists, it moves the key to the front of the cache and returns the value.
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.cache[key]; ok {
		c.items.MoveToFront(element)
		return element.Value.(*CacheItem).value, true
	}
	return nil, false
}

// Delete removes a key-value pair from the cache. If the key exists, it stops the expiration timer, removes the key from the cache map,
// and removes the item from the doubly linked list.
func (c *LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.cache[key]; ok {
		item := element.Value.(*CacheItem)
		if item.expiration != nil {
			item.expiration.Stop()
		}
		delete(c.cache, key)
		c.items.Remove(element)
	}
}
