package hw04lrucache

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
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		front := c.queue.Front()
		front.Value.(*ListItem).Value = value
		c.items[key] = front

		return true
	}

	if c.capacity <= 10 {
		item := c.queue.PushFront(&ListItem{Value: value})
		c.items[key] = item

		c.capacity++
	} else {
		last := c.queue.Back()
		delete(c.items, key)
		c.queue.Remove(last)

		item := c.queue.PushFront(&ListItem{Value: value})
		c.items[key] = item
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value.(*ListItem).Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {

}
