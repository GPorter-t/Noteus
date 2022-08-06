package storage

import (
	"container/list"
	"sync"
)

type lruCache struct {
	MaxEntries int

	OnEvicted func(key string, value interface{})

	ll *list.List

	mutex *sync.RWMutex

	cache map[interface{}]*list.Element
}

type entry struct {
	key   string
	value interface{}
}

func newlruCache(maxSize int) *lruCache {
	return &lruCache{
		MaxEntries: maxSize,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
		mutex:      &sync.RWMutex{},
	}
}

func (c *lruCache) Add(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

func (c *lruCache) Get(key string) (value interface{}, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

func (c *lruCache) Remove(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

func (c *lruCache) RemoveOldest() {
	if c.cache == nil {
		return
	}

	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *lruCache) GetBack() (key string, value interface{}, ok bool) {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	e := ele.Value.(*entry)
	key = e.key
	value = e.value
	ok = true
	c.removeElement(ele)
	return
}

func (c *lruCache) removeElement(ele *list.Element) {
	c.ll.Remove(ele)

	kv := ele.Value.(*entry)

	delete(c.cache, kv.key)

	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *lruCache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.ll = nil
	c.cache = nil
}

type LRUStore struct {
	MaxSize int
	Store   map[string]*lruCache
	mutex   *sync.RWMutex
}

func NewLRUStore(maxSize int) *LRUStore {
	return &LRUStore{
		MaxSize: maxSize,
		Store:   make(map[string]*lruCache),
		mutex:   &sync.RWMutex{},
	}
}

func (s *LRUStore) Get(table string, key string) (value interface{}, ok bool) {
	if c, stat := s.Store[table]; stat {
		value, ok = c.Get(key)
		return
	}
	ok = false
	return
}

func (s *LRUStore) Add(table string, key string, value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.Store[table]; !ok {
		s.Store[table] = newlruCache(s.MaxSize)
	}
	s.Store[table].Add(key, value)
	return
}

func (s *LRUStore) Remove(table string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.Store[table]; ok {
		delete(s.Store, table)
	}
}

func (s *LRUStore) RemoveItem(table string, key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c, ok := s.Store[table]; ok {
		c.Remove(key)
	}
}

func (s *LRUStore) Clear(table string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c, ok := s.Store[table]; ok {
		c.Clear()
	}
}

func (s *LRUStore) Len(table string) int {
	if c, ok := s.Store[table]; ok {
		return c.Len()
	}
	return 0
}

func (s *LRUStore) GetBack(table string) (key string, value interface{}, ok bool) {
	if c, stat := s.Store[table]; stat {
		if key, value, ok = c.GetBack(); ok {
			return
		}
	}
	return
}
