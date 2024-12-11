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

func (lru *lruCache) Set(key Key, value interface{}) bool {
	item, ok := lru.items[key]
	if !ok || item.Value == nil {
		if ok {
			delete(lru.items, key)
		}
		if lru.queue.Len() == lru.capacity {
			itemToDelete := lru.queue.Back()
			lru.queue.Remove(itemToDelete)
			itemToDelete.Value = nil
			itemToDelete.Prev = nil
			itemToDelete.Next = nil
		}
		lru.items[key] = lru.queue.PushFront(value)
		return false
	}
	item.Value = value
	lru.queue.MoveToFront(item)
	return true
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := lru.items[key]
	if !ok || item.Value == nil {
		if ok {
			delete(lru.items, key)
		}
		return nil, false
	}
	lru.queue.MoveToFront(item)
	return item.Value, true
}

func (lru *lruCache) Clear() {
	lru.queue = NewList()
	lru.items = make(map[Key]*ListItem, lru.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
