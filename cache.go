package cache

import "time"

type Cache struct {
	items map[string]Item
}

type Item struct {
	Value    string
	Deadline time.Time
	Expiring bool
}

func NewCache() Cache {

	items := make(map[string]Item)
	return Cache{items: items}
}

func (cache *Cache) Get(key string) (string, bool) {
	cache.cleanUp()
	item, ok := cache.items[key]
	if !ok {
		return "", false
	}
	return item.Value, true
}

func (cache *Cache) Put(key, value string) {
	cache.items[key] = Item{
		Value:    value,
		Expiring: false,
	}
}

func (cache *Cache) Keys() []string {
	cache.cleanUp()
	var keys []string
	for k := range cache.items {
		keys = append(keys, k)
	}
	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.items[key] = Item{
		Value:    value,
		Deadline: deadline,
		Expiring: true,
	}
}

func (cache *Cache) cleanUp() {
	for k, v := range cache.items {
		if v.Expiring && v.Deadline.Before(time.Now()) {
			delete(cache.items, k)
		}
	}
}
