package pokecache

type Cache struct {
	cache map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string][]byte),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	result, exists := c.cache[key]
	return result, exists
}
