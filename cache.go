package cache

type Cacher struct {
	Maps map[string]int64
}

func New() *Cacher {
	return &Cacher{
		Maps: make(map[string]int64),
	}
}

func (c *Cacher) Set(key string, number int64) {
	c.Maps[key] = number
}

func (c Cacher) Get(key string) int64 {
	return c.Maps[key]
}

func (c *Cacher) Delete(key string) {
	delete(c.Maps, key)
}
