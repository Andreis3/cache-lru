package lru2

import "time"

type Item struct {
	key   string
	value any
	time  time.Time
}

type Cache struct {
	capacity int
	items    map[string]*Item
}

func New(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		items:    make(map[string]*Item),
	}
}

func (c *Cache) Get(key string) any {
	item, ok := c.items[key]
	if !ok {
		return -1
	}

	c.items[key].time = time.Now()
	return item.value
}

func (c *Cache) Set(key string, value any) {
	if len(c.items) == c.capacity {
		oldestTime := time.Now()
		var oldestKey string
		for i := range c.items {
			if c.items[i].time.Before(oldestTime) {
				oldestTime = c.items[i].time
				oldestKey = c.items[i].key
			}
		}
		delete(c.items, oldestKey)
	}

	c.items[key] = &Item{
		key:   key,
		value: value,
		time:  time.Time{},
	}
}
