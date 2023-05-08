package prototype

type prototype interface {
	clone() (prototype, error)
}

type cache struct {
	cache map[string]string
}

func (c *cache) clone() (prototype, error) {
	var cc = &cache{cache: map[string]string{}}
	for s, s2 := range c.cache {
		cc.cache[s] = s2
	}

	return cc, nil
}
