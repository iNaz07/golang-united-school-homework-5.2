package cache

import "time"

type Cache struct {
	CacheMap map[string]Value
}

type Value struct {
	Val   string
	ExpAt *time.Time
}

func NewCache() Cache {
	cache := make(map[string]Value)
	return Cache{
		CacheMap: cache,
	}
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.CacheMap[key]
	if !ok {
		return "", false
	}
	if value.ExpAt.Before(time.Now()) {
		delete(c.CacheMap, key)
		return "", false
	}
	return value.Val, true
}

func (c *Cache) Put(key, value string) {
	val := Value{Val: value}
	c.CacheMap[key] = val
}

func (c *Cache) Keys() []string {
	keys := []string{}
	for k, v := range c.CacheMap {
		if !v.ExpAt.Before(time.Now()) {
			keys = append(keys, k)
		} else {
			delete(c.CacheMap, k)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	val := Value{
		Val:   value,
		ExpAt: &deadline,
	}
	c.CacheMap[key] = val
}
