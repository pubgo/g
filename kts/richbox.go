package kts

import (
	"container/list"
	"sync"
	"time"
)

type richBoxEntry struct {
	key       string
	value     interface{}
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

type RichBox struct {
	mu sync.Mutex

	list  *list.List
	table map[string]*list.Element

	expiration time.Duration
	maxLen     int
}

func New(maxLen int, expiration time.Duration) *RichBox {
	return &RichBox{
		list:  list.New(),
		table: make(map[string]*list.Element, maxLen),

		expiration: expiration,
		maxLen:     maxLen,
	}
}

func (c *RichBox) Get(key string) ([]byte, bool) {
	return c.get(key)
}

func (c *RichBox) get(key string) (interface{}, bool) {
	c.mu.Lock()

	el := c.table[key]
	if el == nil {
		c.mu.Unlock()
		return nil, false
	}

	entry := el.Value.(*richBoxEntry)
	if time.Since(entry.createdAt) > c.expiration {
		c.deleteElement(el)
		c.mu.Unlock()
		return nil, false
	}

	c.list.MoveToFront(el)
	value := entry.value
	c.mu.Unlock()
	return value, true
}

func (c *RichBox) Set(key string, value []byte) {
	c.mu.Lock()
	if el := c.table[key]; el != nil {
		entry := el.Value.(*richBoxEntry)
		entry.value = value
		c.promote(el, entry)
	} else {
		c.addNew(key, value)
	}
	c.mu.Unlock()
}

func (c *RichBox) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	el := c.table[key]
	if el == nil {
		return false
	}

	c.deleteElement(el)
	return true
}

func (c *RichBox) Len() int {
	return c.list.Len()
}

func (c *RichBox) Flush() error {
	c.mu.Lock()
	c.list = list.New()
	c.table = make(map[string]*list.Element, c.maxLen)
	c.mu.Unlock()
	return nil
}

func (c *RichBox) addNew(key string, value []byte) {
	newEntry := &richBoxEntry{
		key:       key,
		value:     value,
		createdAt: time.Now(),
	}
	element := c.list.PushFront(newEntry)
	c.table[key] = element
	c.check()
}

func (c *RichBox) promote(el *list.Element, entry *richBoxEntry) {
	entry.createdAt = time.Now()
	c.list.MoveToFront(el)
}

func (c *RichBox) deleteElement(el *list.Element) {
	c.list.Remove(el)
	delete(c.table, el.Value.(*richBoxEntry).key)
}

func (c *RichBox) check() {
	for c.list.Len() > c.maxLen {
		el := c.list.Back()
		c.deleteElement(el)
	}
}
