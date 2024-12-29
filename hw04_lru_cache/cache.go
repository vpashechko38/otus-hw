package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem

	mutex sync.RWMutex
}

type Item struct {
	key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),

		mutex: sync.RWMutex{},
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()

	defer c.mutex.Unlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		front := c.queue.Front()
		front.Value.(*Item).Value = value
		c.items[key] = front

		return true
	}

	if c.capacity > c.queue.Len() {
		item := c.queue.PushFront(&Item{key: key, Value: value})
		c.items[key] = item
	} else {
		last := c.queue.Back()
		delete(c.items, last.Value.(*Item).key)
		c.queue.Remove(last)

		item := c.queue.PushFront(&Item{key: key, Value: value})
		c.items[key] = item
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value.(*Item).Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}
